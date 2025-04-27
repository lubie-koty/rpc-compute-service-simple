package grpc

import (
	"context"

	"github.com/lubie-koty/rpc-compute-service-simple/internal/core"
	pb "github.com/lubie-koty/rpc-compute-service-simple/internal/proto"
)

type GRPCService struct {
	service core.MathService
}

func NewGRPCService(service core.MathService) *GRPCService {
	return &GRPCService{
		service: service,
	}
}

func (h *GRPCService) Add(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (h *GRPCService) Sub(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (h *GRPCService) Mul(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}

func (h *GRPCService) Div(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	return &pb.Response{}, nil
}
