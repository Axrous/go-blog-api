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

type CommentControllerImpl struct {
	CommentService service.CommentService
}

// Create implements CommentController.
func (controller *CommentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentCreateRequest := web.CommentCreateRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&commentCreateRequest)
	helper.PanicIfError(err)

	tokenString := request.Header.Get("Authorization")
	userId := helper.GetUserInfo(tokenString)
	commentCreateRequest.UserId = userId
	commentResponse := controller.CommentService.Save(request.Context(), commentCreateRequest)

	webResonse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: commentResponse,
	}
	helper.ResponToBody(writer, webResonse)
}

// Delete implements CommentController.
func (controller *CommentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	CommentId := params.ByName("commentId")
	id, err := strconv.Atoi(CommentId)
	helper.PanicIfError(err)

	controller.CommentService.Delete(request.Context(), id)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.ResponToBody(writer, webResponse)
}

// FindById implements CommentController.
func (controller *CommentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	CommentId := params.ByName("commentId")
	id, err := strconv.Atoi(CommentId)
	helper.PanicIfError(err)

	commentResponse := controller.CommentService.FindById(request.Context(), id)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: commentResponse,
	}

	helper.ResponToBody(writer, webResponse)
}

// FindByPostId implements CommentController.
func (controller *CommentControllerImpl) FindByPostId(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	postId := request.URL.Query().Get("postId")
	id, err := strconv.Atoi(postId)
	helper.PanicIfError(err)

	commentResponse := controller.CommentService.FindByPostId(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: commentResponse,
	}

	helper.ResponToBody(writer, webResponse)
}

// Update implements CommentController.
func (controller *CommentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentUpdateRequest := web.CommentUpdateRequest{}

	commentId := params.ByName("commentId")
	id, err := strconv.Atoi(commentId)
	helper.PanicIfError(err)

	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&commentUpdateRequest)
	helper.PanicIfError(err)
	commentUpdateRequest.Id = id

	commentResponse := controller.CommentService.Update(request.Context(), commentUpdateRequest)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: commentResponse,
	}

	helper.ResponToBody(writer, webResponse)
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}
