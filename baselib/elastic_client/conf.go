/* !!
 * File: conf.go
 * File Created: Thursday, 20th May 2021 10:33:24 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 20th May 2021 10:35:05 am
 
 */

package elastic_client

import (
	"fmt"
	"strings"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

var ()

const (
	KDefaultTimeout = 30 * time.Second
)

// ElasticConfig type;
type ElasticConfig struct {
	AppKey          string
	SecretKey       string
	PlatformDefault string
	LogTypeDefault  string

	Name        string
	Environment string
	Hosts       []string
	Username    string
	Password    string
}

var esConf *ElasticConfig

// NewElasticClient type;
func NewElasticClient() (client *elastic.Client) {
	cfg := elastic.Config{
		Addresses: esConf.Hosts,
		Username:  esConf.Username,
		Password:  esConf.Password,
	}
	var err error

	client, err = elastic.NewClient(cfg)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("NewElasticClient - Can't connect to Elastic Search server...!")
	}

	g_log.V(1).Info("NewElasticClient - ElasticManager initialized successfully!")

	return client
}

// default value env key is "Elastic";
// if configKeys was set, key env will be first value (not empty) of this;
func getConfigFromEnv(configKeys ...string) {
	configKey := "Elastic"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	esConf = &ElasticConfig{}

	if err := viper.UnmarshalKey(configKey, esConf); err != nil {
		err := fmt.Errorf("not found config with env %q for Elastic with error: %+v", configKey, err)
		panic(err)
	}

	if esConf.Name == "" {
		err := fmt.Errorf("not found config name with env %q for Elastic", fmt.Sprintf("%s.Name", configKey))
		panic(err)
	}

	if esConf.Environment == "" {
		err := fmt.Errorf("not found config environment with env %q for Elastic", fmt.Sprintf("%s.Environment", configKey))
		panic(err)
	}

	if len(esConf.Hosts) == 0 {
		err := fmt.Errorf("not found config hosts with env %q for Elastic", fmt.Sprintf("%s.Hosts", configKey))
		panic(err)
	}

	if esConf.AppKey == "" {
		if esConf.AppKey = env.LogEventAppKeyDefault; esConf.AppKey == "" {
			err := fmt.Errorf("not found config AppKey with env %q for Elastic", fmt.Sprintf("%s.AppKey", configKey))
			panic(err)
		}
	}

	if esConf.SecretKey == "" {
		if esConf.SecretKey = env.LogEventSecretKeyDefault; esConf.SecretKey == "" {
			err := fmt.Errorf("not found config SecretKey from hosts env: %q", fmt.Sprintf("%s.SecretKey", configKey))
			panic(err)
		}
	}

	if esConf.PlatformDefault == "" {
		if esConf.PlatformDefault = env.LogEventPlatformDefault; esConf.PlatformDefault == "" {
			err := fmt.Errorf("not found config PlatformDefault from hosts env: %q", fmt.Sprintf("%s.PlatformDefault", configKey))
			panic(err)
		}
	}

	if esConf.LogTypeDefault == "" {
		if esConf.LogTypeDefault = env.LogEventLogTypeDefault; esConf.LogTypeDefault == "" {
			err := fmt.Errorf("not found config LogTypeDefault from hosts env: %q", fmt.Sprintf("%s.LogTypeDefault", configKey))
			panic(err)
		}
	}
}

func SetESConfig(appKey, platformDefault, logTypeDefault string) *ElasticConfig {
	if esConf == nil {
		esConf = &ElasticConfig{}
	}

	esConf.AppKey = appKey
	esConf.PlatformDefault = platformDefault
	esConf.LogTypeDefault = logTypeDefault

	return esConf
}

func GetESConfig() *ElasticConfig {
	if esConf == nil {
		err := fmt.Errorf("need config for elastic client first")
		panic(err)
	}

	return esConf
}
