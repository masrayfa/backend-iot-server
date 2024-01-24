package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
)

type ChannelRepositoryImpl struct {
}

func NewChannelRepositoryImpl() ChannelRepository {
	return &ChannelRepositoryImpl{}
}

func (r *ChannelRepositoryImpl) Create(ctx context.Context, pool *pgxpool.Pool, channelPayload domain.Channel) (domain.Channel, error) {
	tx, err := pool.Begin(ctx)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx, tx)

	time := time.Now().UTC()

	channel := domain.Channel{
		Time: time, 
		Value: channelPayload.Value,
		IdNode: channelPayload.IdNode,
	}

	script := `INSERT INTO feed (time, value, id_node) VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, script, time, channel.Value, channel.IdNode)
	helper.PanicIfError(err)

	return channel, nil
}