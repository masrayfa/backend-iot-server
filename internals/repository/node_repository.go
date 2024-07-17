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
	// Find hardware node by id either hardware_node or hardware_sensor
	FindHardwareNode(ctx context.Context, pool *pgxpool.Pool, userId int64, id int64) ([]domain.NodeByHardware, error)
	Create(ctx context.Context, pool *pgxpool.Pool, node domain.Node, currentUserId int64) (domain.Node, error)
	Update(ctx context.Context, pool *pgxpool.Pool, node *domain.Node, payload *web.NodeUpdateRequest) (domain.Node, error)
	Delete(ctx context.Context, pool *pgxpool.Pool, id int64) error
}
