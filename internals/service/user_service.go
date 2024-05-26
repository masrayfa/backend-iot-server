package service

import (
	"context"
	"net/http"

	"github.com/masrayfa/internals/models/web"
)

type UserService interface {
	FindAll(ctx context.Context) ([]web.UserRead, error)
	FindById(ctx context.Context, id int64) (web.UserRead, error)
	Register(ctx context.Context,  req *http.Request, payload web.UserCreateRequest) (web.UserRead, error)
	Login(ctx context.Context,  req web.UserLoginRequest) (web.UserRead, error)
	Activation(ctx context.Context,  token string) error
	ForgotPassword(ctx context.Context,  req web.UserForgotPasswordRequest) (string, error)
	MatchPassword(ctx context.Context,  id int64, password string) error
	UpdatePassword(ctx context.Context,  id int64, password string) error
	Delete(ctx context.Context,  id int64) error
}