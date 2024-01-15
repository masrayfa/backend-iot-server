package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/controller"
	"github.com/masrayfa/internals/database"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/repository"
	"github.com/masrayfa/internals/service"
)

func main() {
	dbpool := database.NewDBPool()

	validate := validator.New()

	// Repository
	userRepository := repository.NewUserRepository()

	// Service
	userService := service.NewUserService(userRepository, dbpool, validate)

	// Controller
	userController := controller.NewUserController(userService)

	// Router
	mainRouter := httprouter.New()

	// users endpoint
	userRouter := NewUserRouter(userController)

	mainRouter.Handler("POST", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.Handler("GET", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))

	server := http.Server {
		Addr: ":8080",
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Server running on port 8080")
}