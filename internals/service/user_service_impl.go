package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]web.UserRead, error) {
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db

	users, err := service.userRepository.FindAll(ctx, dbpool)
	helper.PanicIfError(err)

	var userResponses []web.UserRead
	for _, user := range users {
		userResponses = append(userResponses, web.UserRead {
			IdUser: user.IdUser,
			Username: user.Username,
			Email: user.Email,
			Status: user.Status,
			IsAdmin: user.IsAdmin,
		})
	}

	return userResponses, nil
}

func (service *UserServiceImpl) FindById(ctx context.Context, id int64) (web.UserRead, error) {
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db

	user, err := service.userRepository.FindById(ctx, dbpool, id)
	helper.PanicIfError(err)

	userResponse := web.UserRead {
		IdUser: user.IdUser,
		Username: user.Username,
		Email: user.Email,
		Status: user.Status,
		IsAdmin: user.IsAdmin,
	}

	return userResponse, nil
}

func (service *UserServiceImpl) Register(ctx context.Context, req *http.Request, payload web.UserCreateRequest) (web.UserRead, error) {
    // validate request
    err := service.validate.Struct(payload)
    if err != nil {
        return web.UserRead{}, err
    }

    // establish db connection
    dbpool := service.db

    // create user object
    user := domain.User{
        Username: payload.Username,
        Email: payload.Email,
        Password: payload.Password,
    }
    log.Println("user service", user)

    _, err = service.userRepository.FindByUsername(ctx, dbpool, user.Username)
    if err == nil {
        return web.UserRead{}, fmt.Errorf("username already exists")
    }

	// Memastikan emg ga ada user sebelumnya
    if helper.IsErrorNotFound(err) {
        return web.UserRead{}, err
    }
    log.Println("username pass")

    _, err = service.userRepository.FindByEmail(ctx, dbpool, user.Email)
    if err == nil {
        return web.UserRead{}, fmt.Errorf("email already exists")
    }
    if helper.IsErrorNotFound(err) {
        return web.UserRead{}, err
    }
    log.Println("email pass")

    // save user
    res, err := service.userRepository.Save(ctx, dbpool, user)
    if err != nil {
        return web.UserRead{}, err
    }
    fmt.Println("res", res)

    // convert user to user response
    userRead := domain.UserRead{
        IdUser:   res.IdUser,
        Username: res.Username,
        Email:    res.Email,
        Status:   res.Status,
        IsAdmin:  res.IsAdmin,
    }

    sendEmail, err := strconv.ParseBool(req.URL.Query().Get("send_email"))
    if err != nil {
        return web.UserRead{}, fmt.Errorf("invalid send_email")
    }

    // send email
    if sendEmail {
        err := service.userRepository.SendEmailActivation(ctx, dbpool, userRead)
        if err != nil {
            return web.UserRead{}, fmt.Errorf("failed to send activation email")
        }
    }

    response := fmt.Sprintf("Success sign in, id: %d. Check email for activation", res.IdUser)
    config := configs.GetConfig()
    if config.Server.Env == "test" {
        jwt, err := helper.SignUserToken(userRead)
        if err != nil {
            return web.UserRead{}, fmt.Errorf("failed to generate JWT token")
        }
        response += fmt.Sprintf(". Token: %s", jwt)
    }


	userResponse := web.UserRead {
		IdUser: userRead.IdUser,
		Username: userRead.Username,
		Email: userRead.Email,
		Status: userRead.Status,
		IsAdmin: userRead.IsAdmin,
	}


    // return response
    return userResponse, nil
}

// func (service *UserServiceImpl) Register(ctx context.Context, req *http.Request, payload web.UserCreateRequest) (web.UserRead, error) {

// 	// validate request
// 	err := service.validate.Struct(req)
// 	helper.PanicIfError(err)

// 	// establish db connection
// 	dbpool := service.db

// 	// create user object
// 	user := domain.User {
// 		Username: payload.Username,
// 		Email: payload.Email,
// 		Password: payload.Password,
// 	}
// 	log.Println("user service", user)

// 	_, err = service.userRepository.FindByUsername(ctx, dbpool, user.Username)
// 	if helper.IsErrorNotFound(err) {
// 		log.Println("username tidak ada")
// 	} 
// 	if err != nil {
// 		http.Error(nil, "Username already exists", http.StatusBadRequest)
// 		panic(err)
// 	}
// 	log.Println("username pass")

// 	_, err = service.userRepository.FindByEmail(ctx, dbpool, user.Email)
// 	if helper.IsErrorNotFound(err) {
// 		log.Println("email tidak ada")
// 	}
// 	if err != nil {
// 		http.Error(nil, "Email already exists", http.StatusBadRequest)
// 		panic(err)
// 	}
// 	log.Println("email pass")

// 	// save user
// 	res, err := service.userRepository.Save(ctx, dbpool, user)
// 	helper.PanicIfError(err)
// 	fmt.Println("res", res)

// 	// convert user to user response
// 	userResponse := web.UserRead {
// 		IdUser: res.IdUser,
// 		Username: res.Username,
// 		Email: res.Email,
// 		Status: res.Status,
// 		IsAdmin: res.IsAdmin,
// 	}

// 	sendEmail, err := strconv.ParseBool(req.URL.Query().Get("send_email"))
// 	if err != nil {
// 		http.Error(nil, "Invalid send_email", http.StatusBadRequest)
// 		panic(err)
// 	}

// 	var userRead domain.UserRead
// 	userRead.IdUser = userResponse.IdUser
// 	userRead.Username = userResponse.Username
// 	userRead.Email = userResponse.Email
// 	userRead.Status = userResponse.Status
// 	userRead.IsAdmin = userResponse.IsAdmin
// 	log.Println("userRead dari register service", userRead)

// 	// send email
// 	if sendEmail {
// 		err := service.userRepository.SendEmailActivation(ctx, dbpool, userRead)
// 		if err != nil {
// 			http.Error(nil, "Failed to send activation email", http.StatusBadRequest)
// 			panic(err)
// 		}
// 	}

// 	response := fmt.Sprintf("Success sign in, id: %d. Check email for activation", user.IdUser)
//     config := configs.GetConfig()
//     if config.Server.Env == "test" {
//         jwt, err := helper.SignUserToken(userRead)
//         if err != nil {
//             http.Error(nil, "Failed to generate JWT token", http.StatusInternalServerError)
// 			panic(err)
//         }
//         response += fmt.Sprintf(". Token: %s|", jwt)
//     }

// 	// return response
// 	return userResponse, nil
// }

func (service *UserServiceImpl) Login(ctx context.Context, req web.UserLoginRequest) (web.UserRead, error) {
	// validate request
	err := service.validate.Struct(req)
	if err != nil {
		return web.UserRead{}, errors.New("invalid request")
	}

	// establish db connection
	dbpool := service.db

	// find user by username
	user, err := service.userRepository.FindByUsername(ctx, dbpool, req.Username)
	if err != nil {
		return web.UserRead{}, errors.New("invalid username")
	}
	log.Println("user diambil dari repo: ", user)

	if !user.Status {
		return web.UserRead{}, errors.New("user is not active")
	}


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
	return web.UserRead {
		IdUser: user.IdUser,
		Username: user.Username,
		Email: user.Email,
		Status: user.Status,
		Token: token,
		IsAdmin: user.IsAdmin,
	}, nil
}

func (service *UserServiceImpl) Activation(ctx context.Context, token string) error {
	log.Println("@UserService::Activation:token", token)
	err := service.validate.Struct(ctx)
	if err != nil {
		return errors.New("invalid request")
	}

	dbpool := service.db

	// validate token
	user, err := helper.ValidateToken(token)
	if err != nil {
		log.Println("token invalid")
		return err
	}
	if user.Status {
		log.Println("user sudah diaktivasi")
		return err
	}
	log.Println("user dari activation service", user)

	// update status
	log.Println("user.IdUser", user.IdUser)
	err = service.userRepository.UpdateStatus(ctx, dbpool, user.IdUser, true)
	if err != nil {
		log.Println("gagal mengaktivasi user")
		return err
	}

	log.Println("user berhasil diaktivasi")

	return nil
}

func (service *UserServiceImpl) ForgotPassword(ctx context.Context,req web.UserForgotPasswordRequest) error {
	err := service.validate.Struct(ctx)
	helper.PanicIfError(err)

	dbpool := service.db
	helper.PanicIfError(err)

	// find user by username and email
	user, err := service.userRepository.FindByUsername(ctx, dbpool, req.Username)
	helper.PanicIfError(err)

	email, err := service.userRepository.FindByEmail(ctx, dbpool, req.Email)
	helper.PanicIfError(err)

	if user.IdUser != email.IdUser {
		http.Error(nil, "Invalid username or email", http.StatusBadRequest)
		panic(err)
	} 

	err = service.validate.VarWithValue(user.Email, req.Email, "eqfield")
	helper.PanicIfError(err)

	if !user.Status {
		http.Error(nil, "User is not active", http.StatusBadRequest)
		panic(err)
	}

	newPassword := helper.GenerateRandomString(8)
	log.Println("newPassword", newPassword)

	// update password
	err = service.userRepository.UpdatePassword(ctx, dbpool, user.IdUser, newPassword)
	helper.PanicIfError(err)

	// generate token
	token, err := SignUserToken(user)
	helper.PanicIfError(err)

	// send email
	log.Println("token", token)

	return nil
}

func (service *UserServiceImpl) MatchPassword(ctx context.Context, id int64, password string) error {
	err := service.validate.Struct(ctx)
	if err != nil {
		return errors.New("invalid request")
	}

	dbpool := service.db

	user, err := service.userRepository.FindById(ctx, dbpool, id)
	if err != nil {
		return err
	}

	// compare password
	err = service.userRepository.MatchPassword(ctx, dbpool, user.IdUser, password)
	log.Println("err: ", err)
	if err != nil {
		log.Println("@UserService::MatchPassword:password-invalid")
		return err
	}

	log.Println("@UserService::MatchPassword:password-valid")
	return nil
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context, id int64, password string) error {
	err := service.validate.Struct(ctx)
	if err != nil {
		return errors.New("invalid request")
	}

	dbpool := service.db

	user, err := service.userRepository.FindById(ctx, dbpool, id)
	if err != nil {
		return err
	}

	log.Println("request password untuk update dengan pass: ", password)

	// update password
	err = service.userRepository.UpdatePassword(ctx, dbpool, user.IdUser, password)
	if err != nil {
		return err
	}

	log.Println("password berhasil diupdate")

	return nil
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
		"idUser": user.IdUser,
		"username": user.Username,
		"email": user.Email,
		"status": user.Status,
		"isAdmin": user.IsAdmin,
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}