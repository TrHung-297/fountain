/* !!
 * File: service_discovery_config.go
 * File Created: Thursday, 27th May 2021 10:19:57 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Thursday, 27th May 2021 10:54:33 am
 
 */

package service_discovery

// ServiceDiscoveryServerConfig func
type ServiceDiscoveryServerConfig struct {
	RegionDC    string
	ServiceName string
	ServerID    string
	ServiceAddr string
	Addrs       []string
	Username    string
	Password    string
	Interval    int // in seconds
	TTL         int // in seconds
}

// ServiceDiscoveryClientConfig func
type ServiceDiscoveryClientConfig struct {
	RegionDC    string
	ServiceName string
	Addrs       []string
	Username    string
	Password    string
	Balancer    string
}
