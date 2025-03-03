package link_service_impl

import (
	"context"
	"link-service/internal/interface/service"
)

var _ service.Service = (*LinkServiceImpl)(nil)

type LinkServiceImpl struct{}

func (l *LinkServiceImpl) GetLink(ctx context.Context, request service.GetLinkRequest) (service.GetLinkResponse, error) {
	//TODO implement me
}

func (l *LinkServiceImpl) AddLink(ctx context.Context, request service.AddLinkRequest) error {
	//TODO implement me
}
