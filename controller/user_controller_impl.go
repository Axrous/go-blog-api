package controller

import (
	"encoding/json"
	"go-blog-api/helper"
	"go-blog-api/model/web"
	"go-blog-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	userService service.UserService
}

// FindAll implements UserController.
func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := controller.userService.FIndAll(request.Context())

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// FindById implements UserController.
func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	panic("unimplemented")
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}