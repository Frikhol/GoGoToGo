package service

import "context"

type GetLinkRequest struct {
}

type GetLinkResponse struct {
}

type AddLinkRequest struct {
}

type Service interface {
	GetLink(ctx context.Context, request GetLinkRequest) (GetLinkResponse, error)
	AddLink(ctx context.Context, request AddLinkRequest) error
}
