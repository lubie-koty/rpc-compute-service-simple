package grpc

import (
	"log/slog"
)

type GRpcServer struct {
	Address string
	Logger  *slog.Logger
	Service *GRPCService
}

func NewGRPCServer(address string, logger *slog.Logger, service *GRPCService) *GRpcServer {
	return &GRpcServer{
		Address: address,
		Logger:  logger,
		Service: service,
	}
}

func (s *GRpcServer) Serve() error {
	return nil
}
