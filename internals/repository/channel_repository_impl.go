package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
)

type ChannelRepositoryImpl struct {
}

func NewChannelRepositoryImpl() ChannelRepository {
	return &ChannelRepositoryImpl{}
}

func (r *ChannelRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, channel domain.Channel) (domain.Channel, error) {
	var emptyChannel domain.Channel
	return emptyChannel, nil
}