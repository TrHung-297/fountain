

package social_auth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

// Provider needs to be implemented for each 3rd party authentication provider
// e.g. Facebook, Twitter, etc...
type Provider interface {
	InstallOAuth(*ProviderConfig)
	Name() string
	SetName(name string)
	BeginAuth(state string) (Session, error)
	UnmarshalSession(string) (Session, error)
	FetchUser(Session) (User, error)
	Debug(bool)
	RefreshToken(refreshToken string) (*oauth2.Token, error) //Get new access token based on the refresh token
	RefreshTokenAvailable() bool                             //Refresh token is provided by auth provider or not
}

const NoAuthUrlErrorMessage = "an AuthURL has not been set"

// Providers is list of known/available providers.
type Providers map[string]Provider

var allProviders = Providers{}

// RegisterProviders adds a list of available providers for use with SocialAuth.
// Can be called multiple times. If you pass the same provider more
// than once, the last will be used.
func RegisterProviders(providers ...Provider) {
	for _, provider := range providers {
		allProviders[provider.Name()] = provider
	}
}

// GetProviders returns a list of all the providers currently in use.
func GetProviders() Providers {
	return allProviders
}

// GetProvider returns a previously created provider. If SocialAuth has not
// been told to use the named provider it will return an error.
func GetProvider(name string) Provider {
	if provider, ok := allProviders[name]; ok {
		return provider
	}

	err := fmt.Errorf("no provider for %s exists", name)
	panic(err)
}

// ContextForClient provides a context for use with oauth2.
func ContextForClient(h *http.Client) context.Context {
	if h == nil {
		return context.Background()
	}
	return context.WithValue(context.Background(), oauth2.HTTPClient, h)
}

// HTTPClientWithFallBack to be used in all fetch operations.
func HTTPClientWithFallBack(h *http.Client) *http.Client {
	if h != nil {
		return h
	}
	return http.DefaultClient
}
