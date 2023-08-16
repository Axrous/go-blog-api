package service

import (
	"context"
	"go-blog-api/model/web"
)

type PostService interface{
	Save(ctx context.Context, request web.PostCreateRequest) web.PostResponse
	FindAll(ctx context.Context) []web.PostResponse
	FindById(ctx context.Context, postId int) web.PostResponse
	Update(ctx context.Context, request web.PostUpdateRequest) web.PostResponse 
	Delete(ctx context.Context, postId int)
}