

package social_auth

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/TrHung-297/fountain/baselib/g_net/g_api"
)

// SessionName is the key used to access the session store.
const SessionName = "_fountain_session"

// Store can/should be set by applications using fountain auth. The default is a cookie store.
var Store *session.Store

var keySet = false

type ProviderParamType string

// ProviderParamKey can be used as a key in context when passing in a provider
const ProviderParamKey ProviderParamType = "provider"

func init() {
	key := "cookie:gapi_session_key"
	keySet = len(key) != 0

	cookieStore := session.New(session.Config{
		Expiration:     30 * 24 * time.Hour,
		KeyLookup:      key,
		CookieHTTPOnly: true,
	})

	Store = cookieStore
}

/*
BeginAuthHandler is a convenience handler for starting the authentication process.
It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".

BeginAuthHandler will redirect the user to the appropriate authentication end-point
for the requested provider.

See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
func BeginAuthHandler(ctx *fiber.Ctx) error {
	url, err := GetAuthURL(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())

	}

	return ctx.Redirect(url, http.StatusTemporaryRedirect)
}

// SetState sets the state string associated with the given request.
// If no state string is associated with the request, one will be generated.
// This state is sent to the provider and can be retrieved during the
// callback.
var SetState = func(ctx *fiber.Ctx) string {
	state := ctx.Query("state")
	if len(state) > 0 {
		return state
	}

	// If a state query param is not passed in, generate a random
	// base64-encoded nonce so that the state on the auth URL
	// is unguessable, preventing CSRF attacks, as described in
	//
	// https://auth0.com/docs/protocols/oauth2/oauth-state#keep-reading
	nonceBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, nonceBytes)
	if err != nil {
		panic("fountain auth: source of randomness unavailable: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

// GetState gets the state returned by the provider during the callback.
// This is used to prevent CSRF attacks, see
// http://tools.ietf.org/html/rfc6749#section-10.12
var GetState = func(ctx *fiber.Ctx) string {
	if string(ctx.Request().Header.Method()) == http.MethodPost {
		return ctx.FormValue("state")
	}
	return ctx.Query("state")
}

/*
GetAuthURL starts the authentication process with the requested provided.
It will return a URL that should be used to send users to.

It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".

I would recommend using the BeginAuthHandler instead of doing all of these steps
yourself, but that's entirely up to you.
*/
func GetAuthURL(ctx *fiber.Ctx) (string, error) {
	if !keySet && Store != nil {
		fmt.Println("goth/fountain auth: no SESSION_SECRET environment variable is set. The default cookie store is not available and any calls will fail. Ignore this warning if you are using a different store.")
	}

	providerName, err := GetProviderName(ctx)
	if err != nil {
		return "", err
	}

	provider := GetProvider(providerName)

	sess, err := provider.BeginAuth(SetState(ctx))
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}

	err = StoreInSession(ctx, providerName, sess.Marshal())

	if err != nil {
		return "", err
	}

	return url, err
}

/*
CompleteUserAuth does what it says on the tin. It completes the authentication
process and fetches all of the basic information about the user from the provider.

It expects to be able to get the name of the provider from the query parameters
as either "provider" or ":provider".

See https://github.com/markbates/goth/examples/main.go to see this in action.
*/
var CompleteUserAuth = func(ctx *fiber.Ctx) (User, error) {
	if !keySet && Store == nil {
		fmt.Println("goth/fountain auth: no SESSION_SECRET environment variable is set. The default cookie store is not available and any calls will fail. Ignore this warning if you are using a different store.")
	}

	providerName, err := GetProviderName(ctx)
	if err != nil {
		return User{}, err
	}

	provider := GetProvider(providerName)

	value, err := GetFromSession(ctx, providerName)
	if err != nil {
		return User{}, err
	}

	defer Logout(ctx)
	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return User{}, err
	}

	err = validateState(ctx, sess)
	if err != nil {
		return User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// get new token and retry fetch
	_, err = sess.Authorize(ctx, provider)
	if err != nil {
		return User{}, err
	}

	err = StoreInSession(ctx, providerName, sess.Marshal())

	if err != nil {
		return User{}, err
	}

	gu, err := provider.FetchUser(sess)
	return gu, err
}

// validateState ensures that the state token param from the original
// AuthURL matches the one included in the current (callback) request.
func validateState(ctx *fiber.Ctx, sess Session) error {
	rawAuthURL, err := sess.GetAuthURL()
	if err != nil {
		return err
	}

	authURL, err := url.Parse(rawAuthURL)
	if err != nil {
		return err
	}

	reqState := GetState(ctx)

	originalState := authURL.Query().Get("state")
	if originalState != "" && (originalState != reqState) {
		return errors.New("state token mismatch")
	}
	return nil
}

// Logout invalidates a user session.
func Logout(ctx *fiber.Ctx) error {
	session, err := Store.Get(ctx)
	if err != nil {
		return err
	}

	log.Printf("Logout - session.ID(): %s for path: %s", session.ID(), ctx.BaseURL())

	return session.Destroy()
}

// GetProviderName is a function used to get the name of a provider
// for a given request. By default, this provider is fetched from
// the URL query string. If you provide it in a different way,
// assign your own function to this variable that returns the provider
// name for your request.
var GetProviderName = getProviderName

func getProviderName(ctx *fiber.Ctx) (string, error) {

	// try to get it from the url param "provider"
	if p := ctx.Params("provider"); p != "" {
		return p, nil
	}

	// try to get it from the url param ":provider"
	if p := ctx.Query("provider"); p != "" {
		return p, nil
	}

	// try to get it from the context's value of "provider" key
	if p := ctx.Get("provider"); p != "" {
		return p, nil
	}

	// try to get it from the go-context's value of providerContextKey key
	if p := g_api.GetContextDataString(ctx, string(ProviderParamKey)); p != "" {
		return p, nil
	}

	// As a fallback, loop over the used providers, if we already have a valid session for any provider (ie. user has already begun authentication with a provider), then return that provider name
	providers := GetProviders()
	session, _ := Store.Get(ctx)

	log.Printf("getProviderName - session.ID(): %s for path: %s", session.ID(), ctx.BaseURL())

	for _, provider := range providers {
		p := provider.Name()
		value := session.Get(p)
		if _, ok := value.(string); ok {
			return p, nil
		}
	}

	// if not found then return an empty string with the corresponding error
	return "", errors.New("you must select a provider")
}

// GetContextWithProvider returns a new request context containing the provider
func GetContextWithProvider(req *http.Request, provider string) *http.Request {
	return req.WithContext(context.WithValue(req.Context(), ProviderParamKey, provider))
}

// StoreInSession stores a specified key/value pair in the session.
func StoreInSession(ctx *fiber.Ctx, key, value string) error {
	session, _ := Store.Get(ctx)

	log.Printf("StoreInSession - session.ID(): %s for path: %s", session.ID(), ctx.BaseURL())

	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(value)); err != nil {
		return err
	}
	if err := gz.Flush(); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}

	session.Set(key, b.String())
	return session.Save()

}

// GetFromSession retrieves a previously-stored value from the session.
// If no value has previously been stored at the specified key, it will return an error.
func GetFromSession(ctx *fiber.Ctx, key string) (string, error) {
	session, err := Store.Get(ctx)
	if err != nil {
		log.Printf("GetFromSession - Store get error: %v", err)
	}

	log.Printf("GetFromSession - session.ID(): %s for path: %s", session.ID(), ctx.BaseURL())

	value := session.Get(key)
	if value == nil {
		return "", fmt.Errorf("could not find a matching session for this request")
	}

	rdata := strings.NewReader(value.(string))
	r, err := gzip.NewReader(rdata)
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(s), nil
}
