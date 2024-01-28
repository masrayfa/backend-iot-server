package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/masrayfa/internals/controller"
)

func NewRouter() *httprouter.Router {
	mainRouter := httprouter.New()

	return mainRouter 
}

func NewUserRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	// users endpoint
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.GET("/", userController.FindAll)
	router.GET("/:id", userController.FindById)
	router.PUT("/:id", userController.UpdatePassword)
	router.DELETE("/:id", userController.Delete)

	return router
}

func NewHardwareRouter(hardwareController controller.HardwareController) *httprouter.Router {
	router := httprouter.New()

	// hardwares endpoint
	router.GET("/", hardwareController.FindAll)
	router.GET("/:id", hardwareController.FindById)
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

	return router
}