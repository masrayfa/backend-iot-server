package main

import (
	"net/http"

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

func CheckHealthRouter() *httprouter.Router {
	router := httprouter.New()

	// health endpoint
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})

	return router
}

func NewUserRouter(userController controller.UserController) *httprouter.Router {
	router := &Router{
		appRouter: httprouter.New(),
	}

	// users endpoint
	router.appRouter.POST("/register", userController.Register) // done
	router.appRouter.POST("/login", userController.Login) // done
	router.appRouter.GET("/", userController.FindAll) // done
	router.appRouter.GET("/activate", userController.Activation) // done
	router.appRouter.GET("/detail/:user_id", userController.FindById) // done
	router.appRouter.PUT("/:id", userController.UpdatePassword) // done
	router.appRouter.DELETE("/:id", userController.Delete) // done
	router.appRouter.POST("/forgot-password", userController.ForgotPassword) // done

	return router.appRouter
}

func NewHardwareRouter(hardwareController controller.HardwareController) *httprouter.Router {
	router := httprouter.New()

	// hardwares endpoint
	router.GET("/", hardwareController.FindAll) // done
	router.GET("/by/:id", hardwareController.FindById) // done
	router.GET("/type/:id", hardwareController.FindHardwareTypeById) // done
	router.POST("/", hardwareController.Create) // done
	router.PUT("/:id", hardwareController.Update) // done
	router.DELETE("/:id", hardwareController.Delete) // done

	return router
}

func NewNodeRouter(nodeController controller.NodeController) *httprouter.Router {
	router := httprouter.New()

	// nodes endpoint
	router.GET("/", nodeController.FindAll) // done
	router.GET("/by/:id", nodeController.FindById) // done
	router.GET("/hardware/:id", nodeController.FindHardwareNode) // done
	router.POST("/", nodeController.Create) // done
	router.PUT("/:id", nodeController.Update) // done
	router.DELETE("/:id", nodeController.Delete) // done

	return router
}

func NewChannelRouter(channelController controller.ChannelController) *httprouter.Router {
	router := httprouter.New()

	// channels endpoint
	router.POST("/", channelController.Create) // done
	router.GET("/download-csv/:id", channelController.DownloadCSV)

	return router
}