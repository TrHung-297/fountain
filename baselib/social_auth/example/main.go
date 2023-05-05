

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
	social_auth "github.com/TrHung-297/fountain/baselib/social_auth"
	_ "github.com/TrHung-297/fountain/baselib/social_auth/provider/facebook"
	_ "github.com/TrHung-297/fountain/baselib/social_auth/provider/google"
)

type customer struct {
	FirstName, LastName string
	Gmail, Password     string
	Mobile              int
}

func main() {
	key := "cookie:secret-session-key" // Replace with your SESSION_SECRET or similar

	store := session.New(session.Config{
		KeyLookup:      key,
		Expiration:     30 * 24 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   false,
		CookiePath:     "/",
	})

	social_auth.Store = store

	providerConfig := make([]*social_auth.ProviderConfig, 0)
	providerConfig = append(providerConfig, &social_auth.ProviderConfig{
		Provider:    "facebook",
		ClientKey:   "620286092434053",
		Secret:      "73500eb7c69c82ff7379f8dcdf95b6bb",
		CallbackURL: "https://1848-118-70-109-20.ngrok.io/auth/facebook/callback",
		Scopes:      []string{"email", "first_name", "last_name", "link", "about", "id", "name", "picture", "location"},
	})
	providerConfig = append(providerConfig, &social_auth.ProviderConfig{
		Provider:    "google",
		ClientKey:   "841591509515-59d1opco3r4v1t4gqa4uitbb7tac52j4.apps.googleusercontent.com",
		Secret:      "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",
		CallbackURL: "http://localhost:8000/auth/google/callback",
		Scopes:      []string{"email", "profile"},
	})

	social_auth.InstallSocialOAuthManagerWithConfig(providerConfig...)

	engine := html.NewFileSystem(http.Dir("./static"), ".html")
	engine.Reload(true)
	engine.Debug(true)

	p := fiber.New(fiber.Config{
		Views: engine,
	})

	p.Static("/assets", "./static/assets")

	p.Get("/auth/:provider/callback", completeauth)

	p.Get("/auth/:provider", beginauth)

	p.Get("/", home)

	log.Println("listening on http://localhost:8000")

	panic(p.Listen(":8000"))
}

func home(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).Render("login", nil)
}

func beginauth(ctx *fiber.Ctx) error {
	return social_auth.BeginAuthHandler(ctx) //authentication with provider
}

func completeauth(ctx *fiber.Ctx) error {
	log.Printf("completeauth: url: %s", ctx.Context().URI())

	user, err := social_auth.CompleteUserAuth(ctx) //get autherised data's (name,id,profile)
	if err != nil {
		log.Printf("completeauth - Error: %+v", err)
		return ctx.SendString(err.Error())
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	log.Printf("userJSON: %s", userJSON)

	gmail := user.Email
	firstname := user.FirstName
	lastname := user.LastName
	userid := user.UserID
	provider := user.Provider
	name := user.Name

	fmt.Println("customer email:", gmail)

	fmt.Println("customer Firstname:", firstname)

	fmt.Println("customer Lastname:", lastname)

	fmt.Println("customer user id:", userid)

	fmt.Println("customer data provider:", provider)

	fmt.Println("customer raw data provider:", name)

	return ctx.Status(http.StatusOK).Render("success", user)
}
