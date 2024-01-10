package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
)

type NodeRepository interface {
	FindAll(ctx context.Context, tx pgx.Tx) ([]domain.Node, error)
	FindById(ctx context.Context, tx pgx.Tx, id int64) (domain.Node, error)
	FindByUsername(ctx context.Context, tx pgx.Tx, username string) (domain.Node, error)
	Create(ctx context.Context, tx pgx.Tx, node domain.Node) (domain.Node, error)
	Update(ctx context.Context, tx pgx.Tx, node domain.Node) (domain.Node, error)
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
}
