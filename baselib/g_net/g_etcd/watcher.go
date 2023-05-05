package g_etcd

import (
	"fmt"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/grpc_util/service_discovery"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	gr "google.golang.org/grpc/resolver"
)

type GWatcher struct {
	*GEtcd
	RegionDC    string `json:"region_dc,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
	ServerID    string `json:"server_id,omitempty"`

	ServiceKey string
}

func (d *GEtcd) NewWatcher(discovery *service_discovery.ServiceDiscoveryClientConfig) *GWatcher {
	return &GWatcher{
		GEtcd:       d,
		RegionDC:    discovery.RegionDC,
		ServiceName: discovery.ServiceName,

		ServiceKey: fmt.Sprintf("%s/%s/%s/", KNameProjectDir, discovery.RegionDC, discovery.ServiceName),
	}
}

func (w *GWatcher) Resolver() (gr.Builder, error) {
	return resolver.NewBuilder(w.Client)
}
