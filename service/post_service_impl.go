package service

import (
	"context"
	"database/sql"
	"go-blog-api/exception"
	"go-blog-api/helper"
	"go-blog-api/model/domain"
	"go-blog-api/model/web"
	"go-blog-api/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type PostServiceImpl struct {
	PostRepository repository.PostRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

// Delete implements PostService.
func (service *PostServiceImpl) Delete(ctx context.Context, postId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId) 

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PostRepository.Delete(ctx, tx, post)
}

// FindAll implements PostService.
func (service *PostServiceImpl) FindAll(ctx context.Context) []web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	posts := service.PostRepository.FindAll(ctx, tx)

	return helper.ToPostResponses(posts)
}

// FindById implements PostService.
func (service *PostServiceImpl) FindById(ctx context.Context, postId int) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	post, err := service.PostRepository.FindById(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	postResponse := web.PostResponse{
		Id: post.Id,
		Title: post.Title,
		Content: post.Content,
		AuthorId: post.AuthorId,
		CreatedAt: post.CreatedAt,
	}

	return postResponse
}

// Save implements PostService.
func (service *PostServiceImpl) Save(ctx context.Context, request web.PostCreateRequest) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	createdAt := time.Now().Unix()

	post := domain.Post{
		Title: request.Title,
		Content: request.Content,
		AuthorId: request.AuthorId,
		CreatedAt: int(createdAt),
	}

	result := service.PostRepository.Save(ctx, tx, post)

	postResponse := web.PostResponse{
		Id: result.Id,
		Title: result.Title,
		Content: result.Content,
		AuthorId: result.AuthorId,
		CreatedAt: result.CreatedAt,
	}

	return postResponse
}

// Update implements PostService.
func (service *PostServiceImpl) Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newPost := domain.Post{
		Title: request.Title,
		Content: request.Content,
	}

	result := service.PostRepository.Update(ctx, tx, newPost)

	postResponse := web.PostResponse{
		Id: result.Id,
		Title: result.Title,
		Content: result.Content,
		AuthorId: result.AuthorId,
		CreatedAt: result.CreatedAt,
	}

	return postResponse

}

func NewPostService(repository repository.PostRepository, DB *sql.DB, Validate *validator.Validate) PostService {
	return &PostServiceImpl{}
}
