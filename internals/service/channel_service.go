package service

import (
	"context"

	"github.com/masrayfa/internals/models/web"
)

type ChannelService interface {
	Create(ctx context.Context, req web.ChannelCreateRequest) (web.ChannelReadResponse, error)
}