

package grpc_util

import (
	"context"
	"encoding/base64"
	"fmt"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"
	"gitlab.gplay.vn/gtv-backend/fountain/proto/g_proto"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/metadata"
)

var (
	headerRPCMetadata = "rpc_metadata"
)

// RPCMetadataFromMD func
func RPCMetadataFromMD(md metadata.MD) (*RpcMetadata, error) {
	val := metautils.NiceMD(md).Get(headerRPCMetadata)
	if val == "" {
		return nil, nil
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, fmt.Errorf("base64 decode error, rpc_metadata: %s, error: %v", val, err)
	}

	rpcMetadata := &RpcMetadata{}
	err = proto.Unmarshal(buf, rpcMetadata)
	if err != nil {
		return nil, fmt.Errorf("RpcMetadata unmarshal error, rpc_metadata: %s, error: %v", val, err)
	}

	return rpcMetadata, nil
}

// RPCMetadataFromIncoming func
func RPCMetadataFromIncoming(ctx context.Context) *RpcMetadata {
	md2, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	md, err := RPCMetadataFromMD(md2)
	if err != nil {
		panic(g_proto.NewRpcError(int32(g_proto.GTVRpcErrorCodes_ERROR_OTHER), fmt.Sprintf("%s", err)))
	}

	return md
}

// RPCMetadataToOutgoing func
func RPCMetadataToOutgoing(ctx context.Context, md *RpcMetadata) (context.Context, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("RPCMetadataToOutgoing - Marshal rpc_metadata error: %+v", err)
		return nil, err
	}

	return metadata.NewOutgoingContext(ctx, metadata.Pairs(headerRPCMetadata, base64.StdEncoding.EncodeToString(buf))), nil
}

// RPCMetadataToOutgoingForInternal func
// For send internal server
func RPCMetadataToOutgoingForInternal(ctx context.Context, md *RpcMetadata) (context.Context, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("RPCMetadataToOutgoingForInternal - Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.NewIncomingContext(ctx, metadata.Pairs(headerRPCMetadata, base64.StdEncoding.EncodeToString(buf))), nil
}
