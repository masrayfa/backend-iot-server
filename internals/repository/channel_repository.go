package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
)

type ChannelRepository interface {
	Create(ctx context.Context, tx pgx.Tx, channel domain.Channel) (domain.Channel, error)
}