package service

import (
	"context"
	"database/sql"
	"go-blog-api/exception"
	"go-blog-api/helper"
	"go-blog-api/model/domain"
	"go-blog-api/model/web"
	"go-blog-api/repository"
)

type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	DB                *sql.DB
}

// Delete implements CommentService.
func (service *CommentServiceImpl) Delete(ctx context.Context, commentId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CommentRepository.Delete(ctx, tx, comment.Id)
}

// FindById implements CommentService.
func (service *CommentServiceImpl) FindById(ctx context.Context, commentId int) web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, commentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	commentResponse := web.CommentResponse{
		Id: comment.Id,
		Content: comment.Content,
		PostId: comment.PostId,
		User: web.CommentUserResponse{
			Id: comment.Id,
			Name: comment.UserName,
		},
	}

	return commentResponse
}

// FindByPostId implements CommentService.
func (service *CommentServiceImpl) FindByPostId(ctx context.Context, postId int) []web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comments, err := service.CommentRepository.FindByPostId(ctx, tx, postId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var commentResponses []web.CommentResponse
	for _, comment := range comments {
		commentResponses = append(commentResponses, web.CommentResponse{
			Id: comment.Id,
			Content: comment.Content,
			PostId: comment.PostId,
			User: web.CommentUserResponse{
				Id: comment.Id,
				Name: comment.UserName,
			},
		})
	}

	return commentResponses
}

// Save implements CommentService.
func (service *CommentServiceImpl) Save(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment := domain.Comment{
		Content: request.Content,
		PostId: request.PostId,
		UserId: request.UserId,
	}
	result := service.CommentRepository.Save(ctx, tx, comment)

	commentResponse := web.CommentResponse{
		Id: result.Id,
		Content: result.Content,
		PostId: result.PostId,
		User: web.CommentUserResponse{
			Id: comment.Id,
		},
	}

	return commentResponse
}

// Update implements CommentService.
func (service *CommentServiceImpl) Update(ctx context.Context, request web.CommentUpdateRequest) web.CommentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	comment, err := service.CommentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	comment.Content = request.Content

	comment = service.CommentRepository.Update(ctx, tx, comment)

	commentResponse := web.CommentResponse{
		Id: comment.Id,
		Content: comment.Content,
		PostId: comment.PostId,
		User: web.CommentUserResponse{
			Id: comment.Id,
		},
	}

	return commentResponse
}

func NewCommentService(commentRepository repository.CommentRepository, DB *sql.DB) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		DB: DB,
	}
}
