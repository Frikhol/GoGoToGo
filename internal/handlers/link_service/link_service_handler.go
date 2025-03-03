package link_service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "link-service/proto/link_service"
)

var _ pb.LinkServiceServer = (*LinkServiceHandler)(nil)

type LinkServiceHandler struct{}

func (l *LinkServiceHandler) GetLink(ctx context.Context, request *pb.GetLinkRequest) (*pb.GetLinkResponse, error) {
	return l.GetLink(ctx, request)
}

func (l *LinkServiceHandler) AddLink(ctx context.Context, request *pb.AddLinkRequest) (*emptypb.Empty, error) {
	return l.AddLink(ctx, request)
}
