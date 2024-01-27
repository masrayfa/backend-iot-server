package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/domain"
	"github.com/masrayfa/internals/models/web"
)

type NodeRepository interface {
	FindAll(ctx context.Context, pool *pgxpool.Pool, currentUser *domain.UserRead) ([]domain.Node, error)
	FindById(ctx context.Context, pool *pgxpool.Pool, id int64) (domain.Node, error)
	GetHardwareNode(ctx context.Context, pool *pgxpool.Pool, hadrdwareId int64) ([]domain.Node, error)
	Create(ctx context.Context, pool *pgxpool.Pool, node domain.Node, currentUser *domain.UserRead) (domain.Node, error)
	Update(ctx context.Context, pool *pgxpool.Pool, node domain.Node, payload *web.NodeUpdateRequest) (domain.Node, error)
	Delete(ctx context.Context, pool *pgxpool.Pool, id int64) error
}
