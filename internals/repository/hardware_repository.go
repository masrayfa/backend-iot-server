package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
)

type HardwareRepository interface {
	FindAllItem(ctx context.Context, tx pgx.Tx, statement string) ([]domain.Hardware, error)
	FindAllHardware(ctx context.Context, tx pgx.Tx) ([]domain.Hardware, error)
	FindAllNode(ctx context.Context, tx pgx.Tx) ([]domain.Hardware, error)
	FindAllSensor(ctx context.Context, tx pgx.Tx) ([]domain.Hardware, error)
	FindById(ctx context.Context, tx pgx.Tx, id int64) (domain.Hardware, error)
	Create(ctx context.Context, tx pgx.Tx, hardware domain.Hardware) (domain.Hardware, error)
	Update(ctx context.Context, tx pgx.Tx, hardware domain.Hardware) (domain.Hardware, error)
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
}
