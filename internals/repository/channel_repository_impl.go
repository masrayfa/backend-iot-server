package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
)

type ChannelRepositoryImpl struct {
}

func NewChannelRepository() ChannelRepository {
	return &ChannelRepositoryImpl{}
}

func (r *ChannelRepositoryImpl) Create(ctx context.Context, pool *pgxpool.Pool, channelPayload domain.Channel) (domain.Channel, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return domain.Channel{}, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	time := time.Now().UTC()

	channel := domain.Channel{
		Time: time, 
		Value: channelPayload.Value,
		IdNode: channelPayload.IdNode,
	}

	script := `INSERT INTO feed (time, value, id_node) VALUES ($1, $2, $3)`
	_, err = tx.Exec(ctx, script, time, channel.Value, channel.IdNode)
	if err != nil {
		return domain.Channel{}, errors.New("error when insert channel")
	}

	return channel, nil
}

func (r *ChannelRepositoryImpl) GetNodeChannel(ctx context.Context, pool *pgxpool.Pool, nodeId int64, limit int64, startDate *time.Time, endDate *time.Time) ([]domain.Channel, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)
	var args []interface{}
	args = append(args, nodeId)


	script := `SELECT time, value, id_node FROM feed WHERE id_node = $1`
	if startDate != nil {
        script += " AND time >= $2"
        args = append(args, *startDate)
    }
	if endDate != nil {
        if startDate != nil {
            script += " AND time <= $3"
            args = append(args, *endDate)
        } else {
            script+= " AND time <= $2"
            args = append(args, *endDate)
        }
    }

	if limit >= 0 {
		script += " LIMIT " + strconv.Itoa(int(limit))
	}

	rows, err := tx.Query(ctx, script, args...)
	if err != nil {
		return nil, errors.New("error when query row")
	}
	defer rows.Close()

	var channels []domain.Channel
	for rows.Next() {
		var channel domain.Channel
		err := rows.Scan(&channel.Time, &channel.Value, &channel.IdNode)
		if err != nil {
			return nil, errors.New("error when scan row")
		}

		channels = append(channels, channel)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("error when scan row")
	}

	return channels, nil
}

func (r *ChannelRepositoryImpl) GetNodeChannelMultiple(ctx context.Context, pool *pgxpool.Pool, nodes []domain.Node, limit int64) ([]domain.NodeWithFeed, error) {
	tx, err := pool.Begin(ctx)
	if err != nil {
		return nil, errors.New("error when begin transaction")
	}
	defer helper.CommitOrRollback(ctx, tx)

	idNodes := make([]int64, len(nodes))
	mapIdNodes := make(map[int64]int64)
	nodeWithFeed := make([]domain.NodeWithFeed, len(nodes))

	for idx, node := range nodes {
		idNodes[idx] = node.IdNode
		mapIdNodes[node.IdNode] = int64(idx)
		nodeWithFeed[idx] = domain.NodeWithFeed{
			Node: node,
			Feed: []domain.Channel{},
		}
	}

	script := `SELECT time, value, id_node FROM feed WHERE id_node = ANY($1)` 

	rows, err := tx.Query(ctx, script, idNodes)
	if err != nil {
		return nil, errors.New("error when query row")
	}
	defer rows.Close()

	for rows.Next() {
		var channel domain.Channel
		err := rows.Scan(&channel.Time, &channel.Value, &channel.IdNode)
		if err != nil {
			return nil, errors.New("error when scan row")
		}

		nodeIdIndex := mapIdNodes[int64(channel.IdNode)]
		if limit >= 0 && len(nodeWithFeed[nodeIdIndex].Feed) < int(limit) {
			nodeWithFeed[nodeIdIndex].Feed = append(nodeWithFeed[nodeIdIndex].Feed, channel)
		}
	}

	return nodeWithFeed, nil
}