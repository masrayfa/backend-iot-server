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

	return router
}