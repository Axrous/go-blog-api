package helper

import (
	"go-blog-api/model/domain"
	"go-blog-api/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse{
	return 	web.UserResponse{
		Id: user.Id,
		Username: user.Username,
		Name: user.Name,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse  {
	var userResponses []web.UserResponse

	for _, user := range users {
		userResponses = append(userResponses, web.UserResponse{
			Id: user.Id,
			Username: user.Username,
			Name: user.Name,
		})
	}

	return userResponses
}