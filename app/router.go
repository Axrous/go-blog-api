package app

import (
	"go-blog-api/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(authController controller.AuthController) *httprouter.Router{

	router := httprouter.New()

	router.POST("/api/register", authController.Register)
	router.POST("/api/login", authController.Login)
	return router
}