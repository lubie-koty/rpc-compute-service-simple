package grpc

import (
	"log/slog"
)

type GRPCServer struct {
	Address string
	Logger  *slog.Logger
	Service *GRPCService
}

func NewGRPCServer(address string, logger *slog.Logger, service *GRPCService) *GRPCServer {
	return &GRPCServer{
		Address: address,
		Logger:  logger,
		Service: service,
	}
}

func (s *GRPCServer) Serve() error {
	return nil
}
