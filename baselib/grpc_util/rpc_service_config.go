/* !!
 * File: rpc_service_config.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:53:47 am
 
 */

package grpc_util

import (
	"github.com/TrHung-297/fountain/baselib/grpc_util/service_discovery"
)

// RPCServerConfig func
type RPCServerConfig struct {
	Addr         string
	RpcDiscovery service_discovery.ServiceDiscoveryServerConfig
}
