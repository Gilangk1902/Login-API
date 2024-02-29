package Service

import (
	"UserLoginAPI/Model/Web"
)

type IUserService interface {
	Login(request Web.LoginRequest) (Web.LoginResponse, string)
	AdminOnlyLogin(request Web.LoginRequest) (Web.LoginResponse, string)
	Register(request Web.RegisterRequest) (Web.RegisterResponse, error)
}
