package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/domain"
)

type ChannelRepository interface {
	Create(ctx context.Context, pool *pgxpool.Pool, channelPayload domain.Channel) (domain.Channel, error)
}