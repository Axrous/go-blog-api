package main

import (
	"go-blog-api/app"
	"go-blog-api/controller"
	"go-blog-api/helper"
	"go-blog-api/middleware"
	"go-blog-api/repository"
	"go-blog-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	//Repository
	userRepository := repository.NewUserRepository()
	postRepository := repository.NewPostRepository()

	//Service
	userService := service.NewUserService(userRepository, db, validate)
	postService := service.NewPostService(postRepository, db, validate)

	//Controller
	authController := controller.NewAuthController(userService)
	userController := controller.NewUserController(userService)
	postController := controller.NewPostController(postService)

	router := app.NewRouter(authController, userController, postController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}