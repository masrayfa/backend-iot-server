package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/masrayfa/internals/controller"
	"github.com/masrayfa/internals/database"
	"github.com/masrayfa/internals/dependencies"
	"github.com/masrayfa/internals/helper"
	"github.com/masrayfa/internals/repository"
	"github.com/masrayfa/internals/service"
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

	// Router
	mainRouter := NewRouter()
	// router := httprouter.New()

	// users endpoint
	userRouter := NewUserRouter(userController)
	// hardwares endpoint
	hardwareRouter := NewHardwareRouter(hardwareController)
	// nodes endpoint
	nodeRouter := NewNodeRouter(nodeController)
	// channels endpoint
	channelRouter := NewChannelRouter(channelController)

	// main endpoint
	// router.Handler("POST", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	// router.Handler("GET", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	// router.Handler("PUT", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))
	// router.Handler("DELETE", "/api/v1/user/*path", http.StripPrefix("/api/v1/user", userRouter))

	// router.Handler("POST", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	// router.Handler("GET", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	// router.Handler("PUT", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))
	// router.Handler("DELETE", "/api/v1/hardware/*path", http.StripPrefix("/api/v1/hardware", hardwareRouter))

	// router.Handler("POST", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))
	// router.Handler("GET", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))
	// router.Handler("PUT", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))
	// router.Handler("DELETE", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))

	// router.Handler("POST", "/api/v1/channel/*path", http.StripPrefix("/api/v1/channel", validateDependency.GetAuthentication(channelRouter)))

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
	// main endpoint nodes
	mainRouter.Handler("POST", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))
	mainRouter.Handler("GET", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))
	mainRouter.Handler("PUT", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))
	mainRouter.Handler("DELETE", "/api/v1/node/*path", http.StripPrefix("/api/v1/node", nodeRouter))
	// main endpoint channels
	mainRouter.Handler("POST", "/api/v1/channel/*path", http.StripPrefix("/api/v1/channel", validateDependency.GetAuthentication(channelRouter)))

	server := http.Server {
		Addr: ":8080",
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Server running on port 8080")
}