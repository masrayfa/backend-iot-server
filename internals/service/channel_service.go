package service

import (
	"context"

	"github.com/masrayfa/internals/models/web"
)

type ChannelService interface {
	GetNodeChannel(ctx context.Context, nodeId int64, channelId int64) ([]web.ChannelReadResponse, error)
	GetNodeChannelMultiple(ctx context.Context, req []web.NodeRequest, limit int64) ([]web.NodeWithFeedResponse, error)
	Create(ctx context.Context, req web.ChannelCreateRequest) (web.ChannelReadResponse, error)
}