package transport

import (
	ownerendpoint "github.com/dreadster3/pawcare/services/account/internal/owner/endpoint"
	ownerproto "github.com/dreadster3/pawcare/services/account/internal/owner/proto"
	ownertransport "github.com/dreadster3/pawcare/services/account/internal/owner/transport"
	petendpoint "github.com/dreadster3/pawcare/services/account/internal/pet/endpoint"
	petproto "github.com/dreadster3/pawcare/services/account/internal/pet/proto"
	pettransport "github.com/dreadster3/pawcare/services/account/internal/pet/transport"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

func NewGRPCServer(ownerEndpoints ownerendpoint.Set, petEndpoints petendpoint.Set, logger log.Logger) *grpc.Server {
	server := grpc.NewServer()

	ownerServer := ownertransport.NewGRPCServer(ownerEndpoints, logger)
	ownerproto.RegisterOwnerServiceServer(server, ownerServer)

	petServer := pettransport.NewGRPCServer(petEndpoints, logger)
	petproto.RegisterPetServiceServer(server, petServer)

	return server
}
