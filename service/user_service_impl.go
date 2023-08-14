package service

import (
	"context"
	"database/sql"
	"fmt"
	"go-blog-api/helper"
	"go-blog-api/model/domain"
	"go-blog-api/model/web"
	"go-blog-api/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

// FindForAuth implements UserService.
func (service *UserServiceImpl) FindForAuth(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse {

	hmacSampleSecret := []byte("RAHASIA")

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.User{
		Username: request.Username,
		Password: request.Password,
	}

	result := service.UserRepository.FindForAuth(ctx, tx, user)

	dataUser := web.UserLoginRequest{
		Username: result.Username,
		Password: result.Password,
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(user.Password))
	fmt.Println(dataUser.Password)
	fmt.Println(user.Password)
	helper.PanicIfError(err)
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : dataUser.Username,
		"nbf": time.Now().Unix(),
	})
	tokenString, err := token.SignedString(hmacSampleSecret)
	helper.PanicIfError(err)
	
	response := web.UserLoginResponse{
		Token: tokenString,
	}

	return response
}

// Create implements UserService.
func (service *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)


	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := domain.User{
		Username: request.Username,
		Name:     request.Name,
		Password: string(hashedPassword),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)

}

// FIndAll implements UserService.
func (service *UserServiceImpl) FIndAll(ctx context.Context) []web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}

// FindById implements UserService.
func (service *UserServiceImpl) FindById(ctx context.Context, userId int) web.UserResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		helper.PanicIfError(err)
	}

	return helper.ToUserResponse(user)
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, Validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       Validate,
	}
}
