package middleware

import (
	"fmt"
	"go-blog-api/helper"
	"go-blog-api/model/web"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}


func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request)  {
	
	path := request.URL.Path

	if path == "/api/login" || path == "/api/register" {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	hmacSampleSecret := []byte("RAHASIA")
	tokenString := request.Header.Get("Authorization")

	if tokenString == "" {
		fmt.Println("GAADA TOKEN")
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.ResponToBody(writer, webResponse)
		return 
	}


	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("GAGAL")
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.ResponToBody(writer, webResponse)
		return
	} 
	fmt.Println("SUKSES")
	middleware.Handler.ServeHTTP(writer, request)
}