package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, dbpool *pgxpool.Pool) ([]domain.User, error)
	FindById(ctx context.Context, dbpool *pgxpool.Pool, id int64) (domain.User, error)
	FindByEmail(ctx context.Context, dbpool *pgxpool.Pool, email string) (domain.User, error)
	FindByUsername(ctx context.Context, dbpool *pgxpool.Pool, username string) (domain.User, error)
	FindByToken(ctx context.Context, dbpool *pgxpool.Pool, token string) (domain.User, error)
	Save(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error)
	Update(ctx context.Context, dbpool *pgxpool.Pool, user domain.User) (domain.User, error)
	Delete(ctx context.Context, dbpool *pgxpool.Pool, id int64) error
	UpdateStatus(ctx context.Context, dbpool *pgxpool.Pool, id int64, status bool) error
	UpdatePassword(ctx context.Context, dbpool *pgxpool.Pool, id int64, password string) error
	MatchPassword(ctx context.Context, dbpool *pgxpool.Pool, id int64, password string) error
}