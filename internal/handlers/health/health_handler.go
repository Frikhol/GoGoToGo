package health

import (
	"context"
	pb "link-service/proto/health_service"
)

type HealthHandler struct {
	pb.UnimplementedHealthServiceServer
}

func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

func (h *HealthHandler) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_SERVING}, nil
}
