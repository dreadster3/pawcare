package transport

import (
	"google.golang.org/grpc"
)

func NewGRPCServer() *grpc.Server {
	server := grpc.NewServer()
	return server
}
