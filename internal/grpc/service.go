package grpc

import (
	"context"

	"github.com/lubie-koty/rpc-compute-service-simple/internal/core/types"
	pb "github.com/lubie-koty/rpc-compute-service-simple/protos"
)

type GRPCService struct {
	pb.UnimplementedSimpleComputeServer
	service types.MathService
}

func NewGRPCService(service types.MathService) *GRPCService {
	return &GRPCService{
		service: service,
	}
}

func (h *GRPCService) Add(ctx context.Context, request *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := h.service.Add(request.GetFirstNumber(), request.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (h *GRPCService) Sub(ctx context.Context, request *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := h.service.Sub(request.GetFirstNumber(), request.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (h *GRPCService) Mul(ctx context.Context, request *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := h.service.Mul(request.GetFirstNumber(), request.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}

func (h *GRPCService) Div(ctx context.Context, request *pb.OperationRequest) (*pb.OperationResponse, error) {
	result := h.service.Div(request.GetFirstNumber(), request.GetSecondNumber())
	return &pb.OperationResponse{Result: result}, nil
}
