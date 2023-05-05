package env

import "github.com/spf13/viper"

var (
	LogEventAppKeyDefault    = "gtv_log_event"
	LogEventSecretKeyDefault = "this_is_gtv_event_log"
	LogEventPlatformDefault  = "g_plus2"
	LogEventLogTypeDefault   = "event_log"

	LogEventTypeInstalled  = "Installed"
	LogEventTypeRegistered = "Registered"
)

func init() {
	if s := viper.GetString("Env.LogEventAppKeyDefault"); s != "" {
		LogEventAppKeyDefault = s
	}

	if s := viper.GetString("Env.LogEventSecretKeyDefault"); s != "" {
		LogEventSecretKeyDefault = s
	}

	if s := viper.GetString("Env.LogEventPlatformDefault"); s != "" {
		LogEventPlatformDefault = s
	}

	if s := viper.GetString("Env.LogEventLogTypeDefault"); s != "" {
		LogEventLogTypeDefault = s
	}

}
