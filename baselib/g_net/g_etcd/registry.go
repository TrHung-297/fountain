package g_etcd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/grand"
	"github.com/TrHung-297/fountain/baselib/grpc_util/service_discovery"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

type GRegistry struct {
	*GEtcd
	RegionDC    string `json:"region_dc,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	ServerID    string `json:"server_id,omitempty"`
	ServiceAddr string `json:"service_addr,omitempty"`

	LeaseID    clientv3.LeaseID `json:"lease_id,omitempty"`
	ServiceKey string
}

// func WithDCName(dcName string) DiscoveryOption {
// 	return func(do *discoveryOption) {
// 		do.DCName = dcName
// 	}
// }

// func WithServiceName(serviceName string) DiscoveryOption {
// 	return func(do *discoveryOption) {
// 		do.ServiceName = serviceName
// 	}
// }

// func WithServerID(serverID string) DiscoveryOption {
// 	return func(do *discoveryOption) {
// 		do.ServerID = serverID
// 	}
// }

func (d *GEtcd) NewRegister(discovery *service_discovery.ServiceDiscoveryServerConfig) *GRegistry {
	if discovery.ServerID == "" {
		discovery.ServerID = env.PodName
	}

	if discovery.ServerID == "" {
		discovery.ServerID = fmt.Sprintf("%s-%s", env.ServiceName, strings.ToLower(grand.RandomAlphaOrNumeric(5, true, true)))
	}

	return &GRegistry{
		GEtcd:       d,
		RegionDC:    discovery.RegionDC,
		ServiceName: discovery.ServiceName,
		ServerID:    discovery.ServerID,
		ServiceAddr: discovery.ServiceAddr,

		ServiceKey: fmt.Sprintf("%s/%s/%s", KNameProjectDir, discovery.RegionDC, discovery.ServiceName),
	}
}

func (r *GRegistry) Register() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	em, err := endpoints.NewManager(r.Client, r.ServiceKey)
	if err != nil {
		return err
	}
	return em.AddEndpoint(ctx, r.ServiceKey+"/"+r.ServerID, endpoints.Endpoint{Addr: r.ServiceAddr})
}

func (r *GRegistry) Deregister() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	em, err := endpoints.NewManager(r.Client, r.ServiceKey)
	if err != nil {
		return err
	}
	return em.DeleteEndpoint(ctx, r.ServiceKey+"/"+r.ServerID)
}
