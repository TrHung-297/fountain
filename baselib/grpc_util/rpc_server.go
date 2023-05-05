

package grpc_util

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/baselib/g_net/g_etcd"
	grpc_recovery2 "github.com/TrHung-297/fountain/baselib/grpc_util/middleware/recovery2"
	"github.com/TrHung-297/fountain/baselib/grpc_util/service_discovery"

	"google.golang.org/grpc"
)

// RPCServer type
type RPCServer struct {
	addr     string
	registry *g_etcd.GRegistry
	s        *grpc.Server
}

// NewRPCServer func
func NewRPCServer(addr string, discovery *service_discovery.ServiceDiscoveryServerConfig) *RPCServer {
	s := &RPCServer{
		addr: addr,
	}

	gEtcd := g_etcd.NewEtcdDiscovery(g_etcd.WithEtcdAddrs(discovery.Addrs...), g_etcd.WithEtcdUsername(discovery.Username), g_etcd.WithEtcdPassword(discovery.Password))
	s.registry = gEtcd.NewRegister(discovery)

	s.s = grpc_recovery2.NewRecoveryServer2(BizUnaryRecoveryHandler, BizUnaryRecoveryHandler2, BizStreamRecoveryHandler)

	return s
}

// RegisterRPCServerFunc func
// type func RegisterRPCServerHandler(s *grpc.Server)
type RegisterRPCServerFunc func(s *grpc.Server)

// Serve func
func (s *RPCServer) Serve(regFunc RegisterRPCServerFunc) {
	// defer s.GracefulStop()
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		g_log.V(1).WithError(err).Errorf("RPCServer::Serve - Error failed to listen: %v", err)
		return
	}
	g_log.V(1).Infof("rpc listening on:%s", s.addr)

	if regFunc != nil {
		regFunc(s.s)
	}

	defer s.s.GracefulStop()
	go s.registry.Register()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s2 := <-ch
		g_log.V(3).Infof("exit...")
		s.registry.Deregister()
		if i, ok := s2.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	if err := s.s.Serve(listener); err != nil {
		g_log.V(1).WithError(err).Errorf("failed to serve: %s", err)
	}
}

// Stop func;
func (s *RPCServer) Stop() {
	s.registry.Deregister()
	s.s.GracefulStop()
}

// GetGRPCOriginServer func;
func (s *RPCServer) GetGRPCOriginServer() *grpc.Server {
	return s.s
}
