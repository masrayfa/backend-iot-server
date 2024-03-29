package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/masrayfa/internals/controller"
	"github.com/masrayfa/internals/database"
	"github.com/masrayfa/internals/dependencies"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/middleware"
	"github.com/masrayfa/internals/repository"
	"github.com/masrayfa/internals/service"
	"github.com/rs/cors"
)

func main() {
	dbpool := database.NewDBPool()

	validate := validator.New()
	validateDependency := dependencies.NewValidator(validate)

	// Repository
	userRepository := repository.NewUserRepository()
	hardwareRepository := repository.NewHardwareRepository()
	nodeRepository := repository.NewNodeRepository()
	channelRepository := repository.NewChannelRepository()

	// Service
	userService := service.NewUserService(userRepository, dbpool, validate)
	hardwareService := service.NewHardwareService(hardwareRepository, dbpool, validate)
	nodeService := service.NewNodeService(nodeRepository, hardwareRepository, channelRepository, userRepository, dbpool, validate)
	channelService := service.NewChannelService(channelRepository, nodeRepository,dbpool, validate)

	// Controller
	userController := controller.NewUserController(userService)
	hardwareController := controller.NewHardwareController(hardwareService)
	nodeController := controller.NewNodeController(nodeService)
	channelController := controller.NewChannelController(channelService)

	authenticationMiddleware := middleware.NewAuthenticationMiddleware(&validateDependency)

	// Router
	mainRouter := NewRouter(&authenticationMiddleware)

	// users endpoint
	userRouter := NewUserRouter(userController)
	// hardwares endpoint
	hardwareRouter := NewHardwareRouter(hardwareController)
	// nodes endpoint
	nodeRouter := NewNodeRouter(nodeController)
	// channels endpoint
	channelRouter := NewChannelRouter(channelController)

	// main endpoint users
	mainRouter.appRouter.Handler("POST", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.appRouter.Handler("GET", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.appRouter.Handler("PUT", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))

	// main endpoint hardwares
	mainRouter.appRouter.Handler("POST", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	mainRouter.appRouter.Handler("GET", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	mainRouter.appRouter.Handler("PUT", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	// main endpoint nodes
	mainRouter.appRouter.Handler("POST", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", authenticationMiddleware.ValidateUser(nodeRouter)))
	mainRouter.appRouter.Handler("GET", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", authenticationMiddleware.ValidateUser(nodeRouter)))
	mainRouter.appRouter.Handler("PUT", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", authenticationMiddleware.ValidateUser(nodeRouter)))
	mainRouter.appRouter.Handler("DELETE", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", authenticationMiddleware.ValidateUser(nodeRouter)))
	// main endpoint channels
	mainRouter.appRouter.Handler("POST", "/api/v1/channel/*path", http.StripPrefix("/api/v1/channel", authenticationMiddleware.ValidateUser(channelRouter)))

	server := http.Server {
		Addr: ":8080",
		Handler: cors.AllowAll().Handler(mainRouter.appRouter),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Server running on port 8080")
}