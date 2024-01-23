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
	hardwareRepository := repository.NewHardwareRepository()

	// Service
	userService := service.NewUserService(userRepository, dbpool, validate)
	hardwareService := service.NewHardwareService(hardwareRepository, dbpool, validate)

	// Controller
	userController := controller.NewUserController(userService)
	hardwareController := controller.NewHardwareController(hardwareService)

	// Router
	mainRouter := httprouter.New()

	// users endpoint
	userRouter := NewUserRouter(userController)
	// hardwares endpoint
	hardwareRouter := NewHardwareRouter(hardwareController)

	// main endpoint users
	mainRouter.Handler("POST", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.Handler("GET", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.Handler("PUT", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.Handler("DELETE", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	// main endpoint hardwares
	mainRouter.Handler("POST", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	mainRouter.Handler("GET", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	mainRouter.Handler("PUT", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	mainRouter.Handler("DELETE", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))

	server := http.Server {
		Addr: ":8080",
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Server running on port 8080")
}