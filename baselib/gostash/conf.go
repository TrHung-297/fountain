/* !!
 * File: conf.go
 * File Created: Thursday, 20th May 2021 10:33:24 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 20th May 2021 10:36:56 am
 
 */

package gostash

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/env"
)

type LogstashConfig struct {
	Hosts           []string
	SecretKey       string
	AppKey          string
	PlatformDefault string
	LogTypeDefault  string
}

var logstashConf *LogstashConfig

// default value env key is "LogStash";
// if configKeys was set, key env will be first value (not empty) of this;
func createConfigFromEnv(configKeys ...string) {
	configKey := "LogStash"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	logstashConf = &LogstashConfig{}

	if err := viper.UnmarshalKey(configKey, logstashConf); err != nil {
		err := fmt.Errorf("not found config name with env %q for logstash with error: %+v", configKey, err)
		panic(err)
	}

	if len(logstashConf.Hosts) == 0 {
		err := fmt.Errorf("not found hosts for logstash with env %q", fmt.Sprintf("%s.Hosts", configKey))
		panic(err)
	}

	if logstashConf.AppKey == "" {
		if logstashConf.AppKey = env.LogEventAppKeyDefault; logstashConf.AppKey == "" {
			err := fmt.Errorf("can not split uri from hosts env: %q", fmt.Sprintf("%s.AppKey", configKey))
			panic(err)
		}

	}

	if logstashConf.SecretKey == "" {
		if logstashConf.SecretKey = env.LogEventSecretKeyDefault; logstashConf.SecretKey == "" {
			err := fmt.Errorf("not found SecretKey for logstash with env: %q", fmt.Sprintf("%s.SecretKey", configKey))
			panic(err)
		}
	}

	if logstashConf.PlatformDefault == "" {
		logstashConf.PlatformDefault = env.LogEventPlatformDefault
	}

	if logstashConf.LogTypeDefault == "" {
		logstashConf.LogTypeDefault = env.LogEventLogTypeDefault
	}
}

type DataLogBase interface {
	FromJSON(data []byte) error
	ToJSON() string
	SetEventName(eventName string) DataLogBase
	SetLogDataJSON(dataJSON string) DataLogBase
	SetDescription(description string) DataLogBase
}

// DefaultDataLog type;
type DefaultDataLog struct {
	DCName       string      `protobuf:"bytes,1,opt,name=dc_name,json=dcName,proto3" json:"dc_name,omitempty"`
	HostName     string      `protobuf:"bytes,2,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	PodName      string      `protobuf:"bytes,3,opt,name=pod_name,json=podName,proto3" json:"pod_name,omitempty"`
	ServerID     string      `protobuf:"varint,4,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	EventName    string      `protobuf:"varint,5,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	LogDataJSON  interface{} `protobuf:"varint,6,opt,name=log_data_json,json=logDataJSON,proto3" json:"log_data_json,omitempty"`
	Description  string      `protobuf:"varint,7,opt,name=description,json=description,proto3" json:"description,omitempty"`
	TimeStarted  int64       `protobuf:"varint,8,opt,name=time_started,json=timeStarted,proto3" json:"time_started,omitempty"`
	TimeFinished int64       `protobuf:"varint,9,opt,name=time_finished,json=timeFinished,proto3" json:"time_finished,omitempty"`
	TimeExecute  int64       `protobuf:"varint,10,opt,name=time_execute,json=timeExecute,proto3" json:"time_execute,omitempty"`
}

// FromJSON func;
func (obj *DefaultDataLog) FromJSON(data []byte) error {
	return json.Unmarshal(data, obj)
}

// ToJSON func;
func (obj *DefaultDataLog) ToJSON() string {
	return base.JSONDebugDataString(obj)
}

// SetEventName func;
func (obj *DefaultDataLog) SetEventName(eventName string) DataLogBase {
	obj.EventName = eventName

	return obj
}

// SetLogDataJSON func;
func (obj *DefaultDataLog) SetLogDataJSON(dataJSON string) DataLogBase {
	obj.LogDataJSON = dataJSON

	return obj
}

// SetDescription func;
func (obj *DefaultDataLog) SetDescription(description string) DataLogBase {
	obj.Description = description

	return obj
}

func NewDefaultDataLog() *DefaultDataLog {
	timeNow := time.Now().Unix()
	return &DefaultDataLog{
		DCName:       env.DCName,
		HostName:     env.HostName,
		PodName:      env.PodName,
		ServerID:     env.PodName,
		EventName:    "DefaultLog",
		TimeStarted:  timeNow,
		TimeFinished: timeNow,
		TimeExecute:  0,
	}
}
