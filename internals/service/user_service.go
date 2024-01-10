package service

import (
	"context"

	"github.com/masrayfa/internals/models/web"
)

type UserService interface {
	FindAll(ctx context.Context ) ([]web.UserReadResponse, error)
	Register(ctx context.Context,  req web.UserCreateRequest) (web.UserReadResponse, error)
	Login(ctx context.Context,  req web.UserLoginRequest) (web.UserReadResponse, error)
	Activation(ctx context.Context,  token string) error
	ForgotPassword(ctx context.Context,  req web.UserForgotPasswordRequest) error
	UpdatePassword(ctx context.Context,  id int64, password string) error
	UpdateStatus(ctx context.Context,  id int64, status bool) error
	Delete(ctx context.Context,  id int64) error
}