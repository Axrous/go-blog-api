package service

import (
	"context"
	"go-blog-api/model/web"
)

type CommentService interface {
	Save(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse
	FindByPostId(ctx context.Context, postId int) []web.CommentResponse
	FindById(ctx context.Context, commentId int) web.CommentResponse
	Update(ctx context.Context, request web.CommentUpdateRequest) web.CommentResponse
	Delete(ctx context.Context, commentId int)
}