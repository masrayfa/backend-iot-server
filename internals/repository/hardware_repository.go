package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/domain"
)

type HardwareRepository interface {
	FindAllItem(ctx context.Context, pool *pgxpool.Pool, statement string) ([]domain.Hardware, error)
	FindAllHardware(ctx context.Context, pool *pgxpool.Pool) ([]domain.Hardware, error)
	FindAllNode(ctx context.Context, pool *pgxpool.Pool) ([]domain.Hardware, error)
	FindAllSensor(ctx context.Context, pool *pgxpool.Pool) ([]domain.Hardware, error)
	FindById(ctx context.Context, pool *pgxpool.Pool, id int64) (domain.Hardware, error)
	FindHardwareTypeById(ctx context.Context, pool *pgxpool.Pool, id int64) (string, error)
	Create(ctx context.Context, pool *pgxpool.Pool , hardware domain.Hardware) (domain.Hardware, error)
	Update(ctx context.Context, pool *pgxpool.Pool , hardware domain.Hardware) error
	Delete(ctx context.Context, pool *pgxpool.Pool, id int64) error
}
