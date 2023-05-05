/* !!
 * File: conf.go
 * File Created: Friday, 28th May 2021 2:53:52 pm
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Friday, 28th May 2021 2:53:52 pm
 
 */

package sql_client

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/env"
)

var (
	Name        string
	Environment string
	DSN         string
	Active      int
	Idle        int
	LifeTime    int // In seconds
)

const (
	KDefaultTimeout = 30 * time.Second
)

type SQLConfig struct {
	Name        string `json:"name,omitempty"`
	Driver      string `json:"driver,omitempty"` // can be postgres but default is mysql
	Environment string `json:"environment,omitempty"`
	DSN         string `json:"dsn,omitempty"`
	Active      int    `json:"active,omitempty"`
	Idle        int    `json:"idle,omitempty"`
	Lifetime    int    `json:"lifetime,omitempty"` // Connection's lifetime in seconds
}

var configs []*SQLConfig

// default value env key is "MySQL";
// if configKeys was set, key env will be first value (not empty) of this;
func getConfigFromEnv(configKeys ...string) {
	configKey := "MySQL"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	raw := make([]*SQLConfig, 0)

	if err := viper.UnmarshalKey(configKey, &raw); err != nil {
		err := fmt.Errorf("not found config name with env %q for SQL with error: %+v", configKey, err)
		panic(err)
	}

	configs = make([]*SQLConfig, 0)
	for _, config := range raw {
		if config.DSN == "" {
			continue
		}

		if config.Name == "" {
			config.Name = "immaster"
		}

		if config.Environment == "" {
			config.Environment = env.Environment
		}

		if config.Active == 0 {
			config.Active = 50
		}

		if config.Idle == 0 {
			config.Idle = 50
		}

		if config.Lifetime == 0 {
			config.Lifetime = 5 * 60
		}

		if config.Driver == "" {
			config.Driver = "mysql"
		}

		configs = append(configs, config)
	}

	if len(configs) == 0 {
		err := fmt.Errorf("not found valid config with env %q for SQL", configKey)
		panic(err)
	}
}

func init() {}
