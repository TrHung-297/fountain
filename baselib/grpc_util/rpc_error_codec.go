
package grpc_util

import (
	"encoding/base64"
	"fmt"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
	"gitlab.gplay.vn/gtv-backend/fountain/proto/g_proto"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/metadata"
)

var (
	headerRPCError = "rpc_error"
)

// RPCErrorFromMD func
// Server To Client
func RPCErrorFromMD(md metadata.MD) (rpcErr *g_proto.GTVRpcError) {
	g_log.V(3).Info("rpc error from md: ", md)
	val := metautils.NiceMD(md).Get(headerRPCError)
	if val == "" {
		rpcErr = g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), fmt.Sprintf("Unknown error"))
		g_log.V(1).WithError(rpcErr).Errorf("RPCErrorFromMD - Error: %+v", rpcErr)
		return
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		rpcErr = g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), fmt.Sprintf("Base64 decode error, rpc_error: %s, error: %v", val, err))
		g_log.V(1).WithError(rpcErr).Errorf("RPCErrorFromMD - Error: %+v", rpcErr)
		return
	}

	rpcErr = &g_proto.GTVRpcError{}
	err = proto.Unmarshal(buf, rpcErr)
	if err != nil {
		rpcErr = g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_INTERNAL), fmt.Sprintf("RpcError unmarshal error, rpc_error: %s, error: %v", val, err))
		g_log.V(1).WithError(rpcErr).Errorf("RPCErrorFromMD - Error: %+v", rpcErr)
		return
	}

	return rpcErr
}

func RPCErrorToMD(md *g_proto.GTVRpcError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("RPCErrorToMD - Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRPCError, base64.StdEncoding.EncodeToString(buf)), nil
}
