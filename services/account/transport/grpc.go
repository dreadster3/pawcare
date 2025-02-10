package transport

import (
	ownerendpoint "github.com/dreadster3/pawcare/services/account/owner/endpoint"
	"github.com/dreadster3/pawcare/services/account/owner/proto"
	ownertransport "github.com/dreadster3/pawcare/services/account/owner/transport"
	petendpoint "github.com/dreadster3/pawcare/services/account/pet/endpoint"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

func NewGRPCServer(ownerEndpoints ownerendpoint.Set, petEndpoints petendpoint.Set, logger log.Logger) *grpc.Server {
	server := grpc.NewServer()

	ownerServer := ownertransport.NewGRPCServer(ownerEndpoints, logger)
	proto.RegisterOwnerServiceServer(server, ownerServer)

	return server
}
