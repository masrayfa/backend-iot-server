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
	// envPath := "../.env"
	// err := godotenv.Load(envPath)
	// if err != nil {
	// 	fmt.Println("Error loading .env file")
	// }
	// databaseURL := os.Getenv("DATABASE_URL")
	// fmt.Println("ini database url",databaseURL)
	dbpool := database.NewDBPool()

	validate := validator.New()
	// dbpool, err := pgxpool.New(context.Background(), databaseURL)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer dbpool.Close()

	// var greeting string
	// err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(greeting)

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
	// // users endpoint
	// mainRouter.GET("/api/v1/users", userController.FindAll)
	// mainRouter.POST("/api/v1/users", userController.Register)

	server := http.Server {
		Addr: ":8080",
		Handler: mainRouter,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
	fmt.Println("Server running on port 8080")
}