package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/domain"
)

type ChannelRepository interface {
	GetNodeChannel(ctx context.Context, pool *pgxpool.Pool, nodeId int64, limit int64) ([]domain.Channel, error)
	GetNodeChannelMultiple(ctx context.Context, pool *pgxpool.Pool, nodes []domain.Node, limit int64) ([]domain.NodeWithFeed, error)
	Create(ctx context.Context, pool *pgxpool.Pool, channelPayload domain.Channel) (domain.Channel, error)
}