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
func (controller *AuthControllerImpl) Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// userLoginRequest := web.UserLogin{}

	// decoder
}

// register implements AuthController.
func (controller *AuthControllerImpl) Register(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
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

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func NewAuthController(userService service.UserService) AuthController {
	return &AuthControllerImpl{
		UserService: userService,
	}
}
