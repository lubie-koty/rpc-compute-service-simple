package grpc

import (
	"log/slog"
	"net"

	pb "github.com/lubie-koty/rpc-compute-service-simple/protos"
	"google.golang.org/grpc"
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
	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterSimpleComputeServer(server, s.Service)
	s.Logger.Info("gRPC server started", "address", s.Address)
	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
