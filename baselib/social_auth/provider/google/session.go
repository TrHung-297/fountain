

package google

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/social_auth"
)

// Session stores data during the auth process with Google.
type Session struct {
	AuthURL      string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// GetAuthURL will return the URL set by calling the `BeginAuth` function on the Google provider.
func (s Session) GetAuthURL() (string, error) {
	if s.AuthURL == "" {
		return "", errors.New(social_auth.NoAuthUrlErrorMessage)
	}
	return s.AuthURL, nil
}

// Authorize the session with Google and return the access token to be stored for future use.
func (s *Session) Authorize(ctx *fiber.Ctx, provider social_auth.Provider) (string, error) {
	p := provider.(*Provider)
	token, err := p.config.Exchange(social_auth.ContextForClient(p.Client()), ctx.Query("code"))
	if err != nil {
		return "", err
	}

	if !token.Valid() {
		return "", errors.New("invalid token received from provider")
	}

	s.AccessToken = token.AccessToken
	s.RefreshToken = token.RefreshToken
	s.ExpiresAt = token.Expiry
	return token.AccessToken, err
}

// Marshal the session into a string
func (s Session) Marshal() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s Session) String() string {
	return s.Marshal()
}

// UnmarshalSession will unmarshal a JSON string into a session.
func (p *Provider) UnmarshalSession(data string) (social_auth.Session, error) {
	sess := &Session{}
	err := json.NewDecoder(strings.NewReader(data)).Decode(sess)
	return sess, err
}
