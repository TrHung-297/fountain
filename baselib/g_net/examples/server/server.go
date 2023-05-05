package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"sync"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_net/g_etcd"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/grpc_util/service_discovery"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/grpc_util/service_discovery/examples/proto"

	"context"

	"google.golang.org/grpc"
	// "time"
)

var nodeID = flag.String("node", "node1", "node ID")
var port = flag.Int("port", 8080, "listening port")

type RpcServer struct {
	addr string
	s    *grpc.Server
}

func NewRPCServer(addr string) *RpcServer {
	s := grpc.NewServer()
	rs := &RpcServer{
		addr: addr,
		s:    s,
	}
	return rs
}

func (s *RpcServer) Run() {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return
	}
	log.Printf("rpc listening on:%s", s.addr)

	proto.RegisterEchoServiceServer(s.s, s)
	s.s.Serve(listener)
}

func (s *RpcServer) Stop() {
	s.s.GracefulStop()
}

func (s *RpcServer) Echo(ctx context.Context, req *proto.EchoReq) (*proto.EchoRsp, error) {
	text := "Hello " + req.EchoData + ", I am " + *nodeID
	log.Println(text)

	return &proto.EchoRsp{EchoData: text}, nil
}

func StartService() {
	gEtcd := g_etcd.NewEtcdDiscovery(g_etcd.WithEtcdAddrs("http://10.3.80.74:2379"))
	registry := gEtcd.NewRegister(&service_discovery.ServiceDiscoveryServerConfig{
		RegionDC:    "test",
		ServiceName: "test-sv",
		ServerID:    "1",
		ServiceAddr: fmt.Sprintf("127.0.0.1:%d", *port),
	})

	server := NewRPCServer(fmt.Sprintf("0.0.0.0:%d", *port))
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		server.Run()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		registry.Register()
		wg.Done()
	}()

	//stop the server after one minute
	//go func() {
	//	time.Sleep(time.Minute)
	//	server.Stop()
	//	registry.Deregister()
	//}()

	wg.Wait()
}

//go run server.go -node node1 -port 28544
//go run server.go -node node2 -port 18562
//go run server.go -node node3 -port 27772
func main() {
	flag.Parse()
	StartService()
}
