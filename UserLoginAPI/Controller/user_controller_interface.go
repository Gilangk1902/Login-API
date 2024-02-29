package Controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IUserController interface {
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUserFromCookie(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AdminOnlyLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
