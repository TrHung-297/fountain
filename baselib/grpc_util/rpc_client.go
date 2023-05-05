

package grpc_util

import (
	"fmt"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_net/g_etcd"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/grpc_util/service_discovery"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/grpc_util/middleware/otgrpc"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/balancer/weightedroundrobin"
)

const (
	rls                   = "rls"
	grpcLB                = "grpclb"
	pickFirstBalancerName = "pick_first"
	roundRobin            = "round_robin"
	weightedRoundRobin    = "weighted_round_robin"
	// consistentHash        = "consistent_hash"
)

// NewRPCClientByServiceDiscovery func
func NewRPCClientByServiceDiscovery(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *grpc.ClientConn, err error) {
	tracer := opentracing.GlobalTracer()

	dirSpace := fmt.Sprintf("%s/%s", g_etcd.KNameProjectDir, discovery.RegionDC)
	g_log.V(3).Infof("NewRPCClientByServiceDiscovery - Find config %s from: %s", discovery.ServiceName, dirSpace)

	gEtcd := g_etcd.NewEtcdDiscovery(g_etcd.WithEtcdAddrs(discovery.Addrs...), g_etcd.WithEtcdUsername(discovery.Username), g_etcd.WithEtcdPassword(discovery.Password))
	r := gEtcd.NewWatcher(discovery)

	b, e := r.Resolver()
	if e != nil {
		g_log.V(1).WithError(err).Errorf("NewRPCClientByServiceDiscovery - Error: %+v", e)
		panic(e)
	}

	balancer := roundrobin.Name

	switch discovery.Balancer {
	case rls:
		balancer = rls
	case grpcLB:
		balancer = grpcLB
	case pickFirstBalancerName:
		balancer = grpc.PickFirstBalancerName
	case weightedRoundRobin:
		balancer = weightedroundrobin.Name
		// case consistentHash:
		// b = g_loadbalancer.NewKBalancer(r, g_loadbalancer.NewKetamaSelector(g_loadbalancer.DefaultKetamaKey))
	}

	c, err = grpc.Dial("etcd:///"+r.ServiceKey, grpc.WithResolvers(b), grpc.WithInsecure(), grpc.WithBalancerName(balancer), grpc.WithTimeout(time.Second*5), grpc.WithMaxMsgSize(5*1024*1024), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())), grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads())))
	if err != nil {
		g_log.V(1).WithError(err).Errorf("NewRPCClientByServiceDiscovery - Error: %+v", err)
		panic(err)
	}
	return
}

// NewRPCClientExtraByServiceDiscovery func
func NewRPCClientExtraByServiceDiscovery(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *grpc.ClientConn, err error) {
	tracer := opentracing.GlobalTracer()

	dirSpace := fmt.Sprintf("%s/%s", g_etcd.KNameProjectDir, discovery.RegionDC)
	g_log.V(3).Infof("NewRPCClientExtraByServiceDiscovery - Find config of %s from: %s", discovery.ServiceName, dirSpace)

	gEtcd := g_etcd.NewEtcdDiscovery(g_etcd.WithEtcdAddrs(discovery.Addrs...), g_etcd.WithEtcdUsername(discovery.Username), g_etcd.WithEtcdPassword(discovery.Password))
	g := gEtcd.NewWatcher(discovery)

	b, e := g.Resolver()
	if e != nil {
		g_log.V(1).WithError(e).Errorf("NewRPCClientByServiceDiscovery - Error: %+v", e)
		panic(e)
	}

	balancer := roundrobin.Name

	switch discovery.Balancer {
	case rls:
		balancer = rls
	case grpcLB:
		balancer = grpcLB
	case pickFirstBalancerName:
		balancer = grpc.PickFirstBalancerName
	case weightedRoundRobin:
		balancer = weightedroundrobin.Name
		// case consistentHash:
		// b = g_loadbalancer.NewKBalancer(r, g_loadbalancer.NewKetamaSelector(g_loadbalancer.DefaultKetamaKey))
	}

	c, err = grpc.Dial("etcd:///"+g.ServiceKey, grpc.WithResolvers(b), grpc.WithInsecure(), grpc.WithBalancerName(balancer), grpc.WithTimeout(time.Second*5), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())), grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads())))

	if err != nil {
		g_log.V(1).WithError(err).Errorf("NewRPCClientExtraByServiceDiscovery - Error: %+v", err)
		panic(err)
	}

	return
}

// RPCClient type
type RPCClient struct {
	conn *grpc.ClientConn
}

// NewRPCClient func
func NewRPCClient(discovery *service_discovery.ServiceDiscoveryClientConfig) (c *RPCClient, err error) {
	conn, err := NewRPCClientExtraByServiceDiscovery(discovery)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("NewRPCClient - Error: %+v", err)
		panic(err)
	}
	c = &RPCClient{
		conn: conn,
	}

	// go c.RunLoopGetListAddresses()
	return
}

// GetClientConn func
func (c *RPCClient) GetClientConn() *grpc.ClientConn {
	return c.conn
}
