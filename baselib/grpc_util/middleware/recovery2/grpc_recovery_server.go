

package grpc_recovery2

import (
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/grpc_util/middleware/otgrpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// NewRecoveryServer func
// Initialization shows an initialization sequence with a custom recovery handler func.
// Not use
func NewRecoveryServer(unaryCustomFunc UnaryRecoveryHandlerFunc, streamCustomFunc StreamRecoveryHandlerFunc) *grpc.Server {
	tracer := opentracing.GlobalTracer()

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []Option{
		// grpc_recovery2.WithUnaryRecoveryHandler(unaryCustomFunc),
		WithUnaryRecoveryHandler(unaryCustomFunc),
		WithStreamRecoveryHandler(streamCustomFunc),
	}
	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			UnaryServerInterceptor(opts...),
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
		grpc_middleware.WithStreamServerChain(
			StreamServerInterceptor(opts...),
			otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
	)
	return server
}

// NewRecoveryServer2 func
// Initialization shows an initialization sequence with a custom recovery handler func.
func NewRecoveryServer2(unaryCustomFunc UnaryRecoveryHandlerFunc, unaryCustomFunc2 UnaryRecoveryHandlerFunc, streamCustomFunc StreamRecoveryHandlerFunc) *grpc.Server {
	tracer := opentracing.GlobalTracer()

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []Option{
		// grpc_recovery2.WithUnaryRecoveryHandler(unaryCustomFunc),
		WithUnaryRecoveryHandler(unaryCustomFunc),
		WithUnaryRecoveryHandler2(unaryCustomFunc2),
		WithStreamRecoveryHandler(streamCustomFunc),
	}
	_ = opts
	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g. logging) can operate on the recovered state instead of being directly affected by any panic
	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			UnaryServerInterceptor(opts...),
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
		grpc_middleware.WithStreamServerChain(
			StreamServerInterceptor(opts...),
			otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads()),
		),
	)
	return server
}
