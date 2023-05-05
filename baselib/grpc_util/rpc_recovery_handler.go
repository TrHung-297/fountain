

package grpc_util

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/TrHung-297/fountain/baselib/env"
	"github.com/TrHung-297/fountain/baselib/g_log"
	"github.com/TrHung-297/fountain/proto/g_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BizUnaryRecoveryHandler func
func BizUnaryRecoveryHandler(ctx context.Context, p interface{}) (err error) {
	switch code := p.(type) {
	case *g_proto.GTVRpcError:
		md, _ := RPCErrorToMD(code)
		grpc.SetTrailer(ctx, md)
		err = status.Errorf(codes.Unknown, "panic triggered rpc_error: {%v}", p)
	default:
		err = status.Errorf(codes.Unknown, "panic unknown triggered: %v", p)
		errDesc := fmt.Sprintf("💣💣💣 At %s.\nPanic unknown triggered: %v, trace: %s", env.Environment, err.Error(), debug.Stack())
		g_log.V(1).WithError(err).Errorf("BizUnaryRecoveryHandler - Error: %+v", errDesc)

		// Send log to notify
	}
	return
}

// BizUnaryRecoveryHandler2 func
func BizUnaryRecoveryHandler2(ctx context.Context, p interface{}) (err error) {
	switch code := p.(type) {
	case *g_proto.GTVRpcError:
		md, _ := RPCErrorToMD(code)
		grpc.SetTrailer(ctx, md)
		err = code
	default:
		err = status.Errorf(codes.Unknown, "panic unknown triggered: %v", p)
		errDesc := fmt.Sprintf("💣💣💣 At %s.\nPanic unknown triggered: %+v", env.Environment, err.Error())
		g_log.V(1).WithError(err).Errorf("BizUnaryRecoveryHandler2 - Error: %v", errDesc)

		// Send log to notify
	}

	return
}

// BizStreamRecoveryHandler func
func BizStreamRecoveryHandler(stream grpc.ServerStream, p interface{}) (err error) {
	return
}
