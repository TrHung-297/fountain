

package main

import (
	"log"
	"time"

	"github.com/TrHung-297/fountain/baselib/g_net/g_etcd"
	"github.com/TrHung-297/fountain/baselib/grpc_util/service_discovery"
	"github.com/TrHung-297/fountain/baselib/grpc_util/service_discovery/examples/proto"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

func main() {
	gEtcd := g_etcd.NewEtcdDiscovery(g_etcd.WithEtcdAddrs("http://10.3.80.74:2379"))
	g := gEtcd.NewWatcher(&service_discovery.ServiceDiscoveryClientConfig{
		RegionDC:    "test",
		ServiceName: "test-sv",
	})
	b, e := g.Resolver()
	if e != nil {
		panic(e)
	}

	var DefaultBackoffConfig = grpc.BackoffConfig{
		MaxDelay: 5 * time.Second,
	}
	c, err := grpc.Dial("etcd:///"+g.ServiceKey, grpc.WithResolvers(b), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name), grpc.WithTimeout(time.Second*5), grpc.WithBackoffConfig(DefaultBackoffConfig))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()

	client := proto.NewEchoServiceClient(c)

	for i := 0; i < 1000; i++ {
		resp, err := client.Echo(context.Background(), &proto.EchoReq{EchoData: "round robin"})
		if err != nil {
			log.Println("aa:", err)
			time.Sleep(time.Second)
			continue
		}
		log.Printf(resp.EchoData)
		time.Sleep(time.Second)
	}
}
