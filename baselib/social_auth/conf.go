

package social_auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

type ProviderConfig struct {
	Provider    string
	ClientKey   string
	Secret      string
	CallbackURL string
	Scopes      []string
}

var providerConfigs []*ProviderConfig

func getOAuthConfigFromEnv(configKeys ...string) {
	configKey := "SocialOAuth"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	raw := make([]*ProviderConfig, 0)
	if err := viper.UnmarshalKey(configKey, &raw); err != nil {
		err := fmt.Errorf("not found config name with env %q for social authorization with error: %+v", configKey, err)
		panic(err)
	}

	providerConfigs = make([]*ProviderConfig, 0)
	for _, config := range raw {
		if config.Provider == "" {
			g_log.V(1).Infof("getOAuthConfigFromEnv - Not found name for provider: %s", base.JSONDebugDataString(config))
			continue
		}

		if config.ClientKey == "" || config.Secret == "" {
			g_log.V(1).Infof("getOAuthConfigFromEnv - Not found client key or secret for provider name: %s", config.Provider)
			continue
		}

		if config.CallbackURL == "" {
			g_log.V(1).Infof("getOAuthConfigFromEnv - Not found callback url for provider name: %s", config.Provider)
			continue
		}

		if len(config.Scopes) == 0 {
			g_log.V(1).Infof("getOAuthConfigFromEnv - Not found scopes for provider name: %s, try to set default: email", config.Provider)
			config.Scopes = []string{"email"}
		}

		providerConfigs = append(providerConfigs, config)
	}

	if len(providerConfigs) == 0 {
		err := fmt.Errorf("not found valid config with env %q for Social OAuth", configKey)
		panic(err)
	}
}

// Please import all provider that are using before call that func;
// default value env key is "SocialOAuth";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallSocialOAuthManager(configKeys ...string) {
	getOAuthConfigFromEnv(configKeys...)
	InstallSocialOAuthManagerWithConfig(providerConfigs...)
}

// default value env key is "SocialOAuth";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallSocialOAuthManagerWithConfig(configs ...*ProviderConfig) {
	if len(allProviders) == 0 {
		panic(fmt.Errorf("InstallSocialOAuthManagerWithConfig - Need import to initialize all providers first"))
	}

	key := "cookie:secret-session-key" // Replace with your SESSION_SECRET or similar

	store := session.New(session.Config{
		KeyLookup:      key,
		Expiration:     30 * 24 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   false,
		CookiePath:     "/",
	})

	Store = store

	for _, p := range allProviders {
		for _, conf := range configs {
			if conf.Provider == p.Name() {
				p.InstallOAuth(conf)
			}
		}
	}
}
