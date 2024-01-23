package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/masrayfa/configs"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/models/domain"
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
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db

	users, err := service.userRepository.FindAll(ctx, dbpool)
	helper.PanicIfError(err)

	var userResponses []web.UserReadResponse
	for _, user := range users {
		userResponses = append(userResponses, web.UserReadResponse{
			IdUser: user.IdUser,
			Username: user.Username,
			Email: user.Email,
			Status: user.Status,
			Token: user.Token,
			IsAdmin: user.IsAdmin,
		})
	}

	return userResponses, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int64) (web.UserReadResponse, error) {
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db

	user, err := service.userRepository.FindById(ctx, dbpool, id)
	helper.PanicIfError(err)

	userResponse := web.UserReadResponse {
		IdUser: user.IdUser,
		Username: user.Username,
		Email: user.Email,
		Status: user.Status,
		Token: user.Token,
		IsAdmin: user.IsAdmin,
	}

	return userResponse, nil
}

func (service *UserServiceImpl) Register(ctx context.Context, req web.UserCreateRequest) (web.UserReadResponse, error) {
	// validate request
	err := service.validate.Struct(req)
	helper.PanicIfError(err)

	// establish db connection
	dbpool := service.db

	// create user object
	user := domain.User {
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	}

	// save user
	res, err := service.userRepository.Save(ctx, dbpool, user)
	helper.PanicIfError(err)
	fmt.Println("res", res)

	// convert user to user response
	userResponse := web.UserReadResponse {
		IdUser: res.IdUser,
		Username: res.Username,
		Email: res.Email,
		Status: res.Status,
		Token: res.Token,
		IsAdmin: res.IsAdmin,
	}

	// return response
	return userResponse, nil
}

func (service *UserServiceImpl) Login(ctx context.Context, req web.UserLoginRequest) (web.UserReadResponse, error) {
	// validate request
	err := service.validate.Struct(req)
	helper.PanicIfError(err)

	// establish db connection
	dbpool := service.db

	// find user by username
	user, err := service.userRepository.FindByUsername(ctx, dbpool, req.Username)
	helper.PanicIfError(err)
	log.Println("user diambil dari repo: ", user)

	// compare password
	err = service.userRepository.MatchPassword(ctx, dbpool, user.IdUser, req.Password)
	if err != nil {
		http.Error(nil, "Invalid password", http.StatusBadRequest)
		panic(err)
	}

	// generate token
	token, err := SignUserToken(user)
	helper.PanicIfError(err)

	// return response
	return web.UserReadResponse{
		IdUser: user.IdUser,
		Username: user.Username,
		Email: user.Email,
		Status: user.Status,
		Token: token,
		IsAdmin: user.IsAdmin,
	}, nil
}

func (service *UserServiceImpl) Activation(ctx context.Context, token string) error {
	panic("unimplemented")
}

func (service *UserServiceImpl) ForgotPassword(ctx context.Context,req web.UserForgotPasswordRequest) error {
	panic("unimplemented")
}

func (service *UserServiceImpl) MatchPassword(ctx context.Context,id int64, password string) error {
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db
	helper.PanicIfError(err)

	user, err := service.userRepository.FindById(ctx, dbpool, id)
	helper.PanicIfError(err)

	// // compare password
	err = service.userRepository.MatchPassword(ctx, dbpool, user.IdUser, password)
	if err != nil {
		http.Error(nil, "Invalid password", http.StatusBadRequest)
		panic(err)
	}

	return nil
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context,id int64, password string) error {
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db
	helper.PanicIfError(err)

	// update password
	err = service.userRepository.UpdatePassword(ctx, dbpool, id, password)
	helper.PanicIfError(err)

	log.Println("password berhasil diupdate")

	return nil
}

func (service *UserServiceImpl) UpdateStatus(ctx context.Context,id int64, status bool) error {
	panic("unimplemented")
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int64) error {
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db
	helper.PanicIfError(err)

	// delete user
	err = service.userRepository.Delete(ctx, dbpool, id)
	helper.PanicIfError(err)

	log.Println("user berhasil dihapus")

	return nil
}

func SignUserToken(user domain.User) (string, error) {
	config := configs.GetConfig()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id_user": user.IdUser,
		"username": user.Username,
		"email": user.Email,
		"status": user.Status,
		"isadmin": user.IsAdmin,
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}