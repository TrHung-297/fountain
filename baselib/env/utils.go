package env

import (
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/base"
	"github.com/TrHung-297/fountain/baselib/g_log"
)

func SetupConfigEnv() {
	EnvConfigData = &EnvConfig{}
	if err := viper.UnmarshalKey("Env", EnvConfigData); err != nil {
		g_log.V(1).WithError(err).Errorf("SetupConfigEnv - init Env config - Error: %+v", err)
		panic(err)
	}
	if hostName := os.Getenv("K8S_NODE_IP"); hostName != "" {
		EnvConfigData.HostName = hostName
	}

	if podName := os.Getenv("K8S_POD_NAME"); podName != "" {
		EnvConfigData.PodName = podName
	}

	if podID := os.Getenv("K8S_POD_IP"); podID != "" {
		EnvConfigData.PodID = podID
	}

	if EnvConfigData.PodName == "" {
		EnvConfigData.PodName = uuid.New().String()
		if EnvConfigData.PodID == "" {
			EnvConfigData.PodID = EnvConfigData.PodName
		}
	}

	if EnvConfigData.PodID == "" {
		EnvConfigData.PodID = uuid.New().String()
	}

	LogConfigData = &LogConfig{}
	if err := viper.UnmarshalKey("Log", LogConfigData); err != nil {
		g_log.V(1).WithError(err).Errorf("SetupConfigEnv - init Env config - Error: %v", err)
		panic(err)
	}

	if LogConfigData.Path == "" {
		LogConfigData.Path = "/var/log/GtvPlus"
	}

	if LogConfigData.Prefix == "" {
		if EnvConfigData.ServiceName != "" {
			LogConfigData.Prefix = EnvConfigData.ServiceName
		} else {
			LogConfigData.Prefix = "backend-game"
		}
	}

	if LogConfigData.LogFileLevel == 0 {
		LogConfigData.LogFileLevel = 5
	}

	if viper.GetBool(`Debug`) || LogConfigData.LogPrintLevel == 0 {
		LogConfigData.LogPrintLevel = 5
	}

	log.Printf("SetupConfigEnv - EnvConfigData: %s", base.JSONDebugDataString(EnvConfigData))

	//
	Environment = EnvConfigData.Environment
	Addr = EnvConfigData.Addr
	DCName = EnvConfigData.DCName
	HostName = EnvConfigData.HostName
	PodName = EnvConfigData.PodName
	PodID = EnvConfigData.PodID
	ServiceName = EnvConfigData.ServiceName
	if EndpointPrefix = EnvConfigData.EndpointPrefix; EndpointPrefix == "" {
		prefix := strings.TrimPrefix(ServiceName, "backend-")
		prefix = strings.Replace(prefix, "-", "/", 1)
		EndpointPrefix = prefix
	}

	//
	LogPath = LogConfigData.Path
	LogPrefix = LogConfigData.Prefix
	LogFileLevel = LogConfigData.LogFileLevel
	LogPrintLevel = LogConfigData.LogPrintLevel
	//
}
