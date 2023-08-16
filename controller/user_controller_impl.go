package controller

import (
	"go-blog-api/helper"
	"go-blog-api/model/web"
	"go-blog-api/service"
	"net/http"
	"strconv"

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

	helper.ResponToBody(writer, webResponse)
}

// FindById implements UserController.
func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.userService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}

	helper.ResponToBody(writer, webResponse)
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}
