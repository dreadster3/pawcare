package transport

import (
	"github.com/dreadster3/pawcare/services/account/endpoint"
	"github.com/dreadster3/pawcare/services/account/proto"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
)

type grpcServer struct {
	proto.UnimplementedOwnerServiceServer
}

func NewGRPCServer(endpoints endpoint.Set, logger log.Logger) proto.OwnerServiceServer {
	_ = []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	return &grpcServer{}
}
