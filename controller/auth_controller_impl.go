package controller

import (
	"encoding/json"
	"go-blog-api/helper"
	"go-blog-api/model/web"
	"go-blog-api/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthControllerImpl struct {
	UserService service.UserService
}

// login implements AuthController.
func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userLoginRequest)
	helper.PanicIfError(err)

	token := controller.UserService.FindForAuth(request.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: token,
	}

	helper.ResponToBody(writer, webResponse)
}

// register implements AuthController.
func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisterRequest := web.UserCreateRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&userRegisterRequest)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Create(request.Context(), userRegisterRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func NewAuthController(userService service.UserService) AuthController {
	return &AuthControllerImpl{
		UserService: userService,
	}
}
