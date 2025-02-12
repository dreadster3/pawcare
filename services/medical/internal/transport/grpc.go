package transport

import (
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

func NewGRPCServer(logger log.Logger) *grpc.Server {
	server := grpc.NewServer()

	return server
}
