/* !!
 * File: conf.go
 * File Created: Thursday, 27th May 2021 10:19:32 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:32:37 am
 
 */

package g_api

import (
	"github.com/spf13/viper"
	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/grpc_util/service_discovery"
)

// Config func;
type Config struct {
	Addr         string
	Discovery    *service_discovery.ServiceDiscoveryServerConfig
	PprofEnabled bool
}

var GapiConfig *Config

func getGAPIConfigFromEnv() {
	GapiConfig = new(Config)

	if err := viper.UnmarshalKey("Server", GapiConfig); err != nil {
		g_log.V(1).WithError(err).Errorf("getGAPIConfigFromEnv - Error: %v", err)
		panic(err)
	}

	if GapiConfig.Addr == "" {
		if env.Addr != "" {
			GapiConfig.Addr = env.Addr
		} else {
			GapiConfig.Addr = ":9099"
		}
	}

	env.Addr = GapiConfig.Addr

	if GapiConfig.Discovery.RegionDC == "" {
		GapiConfig.Discovery.RegionDC = env.DCName
	}

	if GapiConfig.Discovery.ServiceName == "" {
		GapiConfig.Discovery.ServiceName = env.ServiceName
	}

	if GapiConfig.Discovery.ServerID == "" {
		GapiConfig.Discovery.ServerID = env.PodName
	}

	if GapiConfig.Discovery.ServiceAddr == "" {
		GapiConfig.Discovery.ServiceAddr = env.ServiceName
	}

	if GapiConfig.Discovery.Interval == 0 {
		GapiConfig.Discovery.Interval = 2
	}

	if GapiConfig.Discovery.TTL == 0 {
		GapiConfig.Discovery.TTL = 10
	}
}

// Models

type ReqParam struct {
	Param        string
	Name         string // Name of param when it convert to body
	DefaultValue string
}
