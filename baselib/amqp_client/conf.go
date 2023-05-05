package amqp_client

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

const (
	KDefaultContentType string = "application/json"
)

// MessageQueue func;
type MessageQueue struct {
	MessageType string      `json:"MessageType"`
	Data        interface{} `json:"Data"`
}

type Config struct {
	Host      string `json:"host,omitempty"`
	Port      int32  `json:"port,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	QueueType string `json:"queue_type,omitempty"`
}

var (
	config *Config
)

// default value env key is "AMQP";
// if configKeys was set, key env will be first value (not empty) of this;
func getAmqpConfigFromEnv(configKeys ...string) {
	configKey := "AMQP"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	config = &Config{}

	if err := viper.UnmarshalKey(configKey, config); err != nil {
		err := fmt.Errorf("not found config name with env %q for amqp with error: %+v", configKey, err)
		panic(err)
	}

	if config.Host == "" {
		err := fmt.Errorf("not found any addr as host for amqp at %q", fmt.Sprintf("%s.Host", configKey))
		panic(err)
	}

	if config.QueueType == "" {
		config.QueueType = "classic"
	}
}
