package service

import (
	"context"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/internals/models/web"
	"github.com/masrayfa/internals/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	db *pgxpool.Pool
	validate *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *pgxpool.Pool, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		db: db,
		validate: validate,
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserReadResponse, error) {
	panic("unimplemented")
}

func (service *UserServiceImpl) Register(ctx context.Context,req web.UserCreateRequest) (web.UserReadResponse, error) {
	panic("unimplemented")
}

func (service *UserServiceImpl) Login(ctx context.Context,req web.UserLoginRequest) (web.UserReadResponse, error) {
	panic("unimplemented")
}

func (service *UserServiceImpl) Activation(ctx context.Context,token string) error {
	panic("unimplemented")
}

func (service *UserServiceImpl) ForgotPassword(ctx context.Context,req web.UserForgotPasswordRequest) error {
	panic("unimplemented")
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context,id int64, password string) error {
	panic("unimplemented")
}

func (service *UserServiceImpl) UpdateStatus(ctx context.Context,id int64, status bool) error {
	panic("unimplemented")
}

func (service *UserServiceImpl) Delete(ctx context.Context,id int64) error {
	panic("unimplemented")
}