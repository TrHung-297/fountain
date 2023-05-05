package crypto

import (
	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

type G_AESConfig struct {
	SecretKey  string
	InitVector string
}

var (
	gAESConfig *G_AESConfig
)

func getG_AESConfigFromEnv(envKey string) {
	gAESConfig = &G_AESConfig{}
	if err := viper.UnmarshalKey(envKey, gAESConfig); err != nil {
		g_log.V(1).WithError(err).Errorf("getG_AESConfigFromEnv - Can not parse env key %q, Error: %v", envKey, err)
		panic(err)
	}
}
