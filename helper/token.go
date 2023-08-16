package helper

import "github.com/golang-jwt/jwt/v5"

func GetToken(tokenString string) (*jwt.Token, error) {

	hmacSampleSecret := []byte("RAHASIA")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	return token, err
}