package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/controller"
	"github.com/masrayfa/internals/middleware"
)

type Router struct {
	appRouter*httprouter.Router
	authMiddleware *middleware.AuthenticationMiddleware
}

func NewRouter(authMiddleware *middleware.AuthenticationMiddleware) *Router {
	return &Router{
		appRouter: httprouter.New(),
		authMiddleware: authMiddleware,
	}
}

func NewUserRouter(userController controller.UserController) *httprouter.Router {
	router := &Router{
		appRouter: httprouter.New(),
	}

	// users endpoint
	router.appRouter.POST("/register", userController.Register)
	router.appRouter.POST("/login", userController.Login)
	router.appRouter.GET("/", userController.FindAll)
	router.appRouter.GET("/:id", userController.FindById)
	router.appRouter.PUT("/:id", userController.UpdatePassword)
	router.appRouter.DELETE("/:id", userController.Delete)
	router.appRouter.POST("/activation", userController.Activation)
	router.appRouter.POST("/forgot-password", userController.ForgotPassword)

	return router.appRouter
}

func NewHardwareRouter(hardwareController controller.HardwareController) *httprouter.Router {
	router := httprouter.New()

	// hardwares endpoint
	router.GET("/", hardwareController.FindAll)
	router.GET("/:id", hardwareController.FindHardwareTypeById)
	router.POST("/", hardwareController.Create)
	router.PUT("/:id", hardwareController.Update)
	router.DELETE("/:id", hardwareController.Delete)

	return router
}

func NewNodeRouter(nodeController controller.NodeController) *httprouter.Router {
	router := httprouter.New()

	// nodes endpoint
	router.GET("/", nodeController.FindAll)
	router.GET("/:id", nodeController.FindById)
	router.POST("/", nodeController.Create)
	router.PUT("/:id", nodeController.Update)
	router.DELETE("/:id", nodeController.Delete)

	return router
}

func NewChannelRouter(channelController controller.ChannelController) *httprouter.Router {
	router := httprouter.New()

	// channels endpoint
	router.POST("/", channelController.Create)

	return router
}