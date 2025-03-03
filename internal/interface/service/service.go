package service

import (
	"context"
	"time"
)

type GetLinkRequest struct {
	URI string
}

type GetLinkResponse struct {
	Link string
}

type AddLinkRequest struct {
	Link      string
	FakeLink  string
	EraseTime time.Time
}

type Service interface {
	GetLink(ctx context.Context, request GetLinkRequest) (GetLinkResponse, error)
	AddLink(ctx context.Context, request AddLinkRequest) error
}
