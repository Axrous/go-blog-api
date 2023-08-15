package app

import (
	"go-blog-api/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(authController controller.AuthController, userController controller.UserController) *httprouter.Router{

	router := httprouter.New()

	//Auth Controller
	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)

	//User Controller
	router.GET("/api/users", userController.FindAll)


	
	return router
}