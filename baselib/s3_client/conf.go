

package s3_client

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/constant"
)

type Config struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
	BaseURL         string
	BucketName      string
	CDN             string
}

var config *Config

// default value env key is "AWS";
// if configKeys was set, key env will be first value (not empty) of this
func getConfigFromEnv(configKeys ...string) *Config {
	configKey := "AWS"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	config = &Config{}
	if err := viper.UnmarshalKey(configKey, config); err != nil {
		err := fmt.Errorf("not found config name with env %q for AWS with error: %+v", configKey, err)
		panic(err)
	}

	if config.Region == "" {
		config.Region = constant.KDefaultAwsRegion
	}

	if config.BucketName == "" {
		config.BucketName = constant.KDefaultS3DirData
	}

	return config
}
