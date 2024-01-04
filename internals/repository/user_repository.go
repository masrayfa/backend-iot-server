package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/masrayfa/internals/models/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, tx pgx.Tx) ([]domain.User, error)
	FindById(ctx context.Context, tx pgx.Tx, id int64) (domain.User, error)
	FindByEmail(ctx context.Context, tx pgx.Tx, email string) (domain.User, error)
	FindByUsername(ctx context.Context, tx pgx.Tx, username string) (domain.User, error)
	FindByToken(ctx context.Context, tx pgx.Tx, token string) (domain.User, error)
	Save(ctx context.Context, tx pgx.Tx, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx pgx.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
	UpdateStatus(ctx context.Context, tx pgx.Tx, id int64, status bool) error
}