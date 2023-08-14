package main

import (
	"go-blog-api/app"
	"go-blog-api/controller"
	"go-blog-api/helper"
	"go-blog-api/repository"
	"go-blog-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	authController := controller.NewAuthController(userService)

	router := app.NewRouter(authController)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}