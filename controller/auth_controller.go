package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthController interface {
	Register(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}