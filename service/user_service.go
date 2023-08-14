package service

import (
	"context"
	"go-blog-api/model/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	// Update(ctx context.Context, )
	FindById(ctx context.Context, userId int) web.UserResponse
	FIndAll(ctx context.Context) []web.UserResponse
	FindForAuth(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
}