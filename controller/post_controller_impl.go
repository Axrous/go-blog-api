package controller

import (
	"encoding/json"
	"go-blog-api/helper"
	"go-blog-api/model/web"
	"go-blog-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PostControllerImpl struct {
	postService service.PostService
}

// Create implements PostController.
func (controller *PostControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	postCreateRequest := web.PostCreateRequest{}
	
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&postCreateRequest)
	helper.PanicIfError(err)

	tokenString := request.Header.Get("Authorization")
	userId := helper.GetUserInfo(tokenString)
	postCreateRequest.AuthorId = userId
	postResponse := controller.postService.Save(request.Context(), postCreateRequest)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: postResponse,
	}

	helper.ResponToBody(writer, webResponse)
}

// Delete implements PostController.
func (controller *PostControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")

	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	controller.postService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.ResponToBody(writer, webResponse)
}

// FindAll implements PostController.
func (controller *PostControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	postResponses := controller.postService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: postResponses,
	}

	helper.ResponToBody(writer, webResponse)
}

// FindById implements PostController.
func (controller *PostControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := params.ByName("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	postResponse := controller.postService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: postResponse,
	}

	helper.ResponToBody(writer, webResponse)
}

// Update implements PostController.
func (controller *PostControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	postUpdateRequest  := web.PostUpdateRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&postUpdateRequest)
	helper.PanicIfError(err)

	postResponse := controller.postService.Update(request.Context(), postUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: postResponse,
	}

	helper.ResponToBody(writer, webResponse)
}

func NewPostController(postService service.PostService) PostController {
	return &PostControllerImpl{}

}
