package Controller

import (
	"UserLoginAPI/Helper"
	"UserLoginAPI/JWT"
	"UserLoginAPI/Model/Web"
	"UserLoginAPI/Service"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	User_service Service.IUserService
}

func NewUserController(user_service Service.IUserService) IUserController {
	return &UserController{
		User_service: user_service,
	}
}

func (user_controller *UserController) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	register_request := Web.RegisterRequest{}
	Helper.ReadFromRequestBody(request, writer, &register_request)

	register_response, err := user_controller.User_service.Register(register_request)
	if err != nil {
		if err.Error() == Helper.FAIL_VALIDATION {
			Helper.WriteFailValidation(writer)
		}
		return
	}

	Helper.WriteOk(writer, register_response)
}

func (user_controller *UserController) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	login_request := Web.LoginRequest{}
	Helper.ReadFromRequestBody(request, writer, &login_request)

	user_response, token_string := user_controller.User_service.Login(login_request)

	if token_string == Helper.NO_TOKEN {
		Helper.WriteNotFound(writer)
		return
	} else if token_string == Helper.FAIL_VALIDATION {
		Helper.WriteFailValidation(writer)
		return
	}

	http.SetCookie(writer,
		&http.Cookie{
			Name:    "token",
			Value:   token_string,
			Expires: time.Now().Add(time.Minute * 1),
		},
	)

	Helper.WriteOk(writer, user_response)
}

func (user_controller *UserController) AdminOnlyLogin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	login_request := Web.LoginRequest{}
	Helper.ReadFromRequestBody(request, writer, &login_request)

	login_response, token_string := user_controller.User_service.AdminOnlyLogin(login_request)
	if token_string == Helper.NO_TOKEN {
		Helper.WriteUnAuthorized(writer)
		return
	} else if token_string == Helper.FAIL_VALIDATION {
		Helper.WriteFailValidation(writer)
		return
	}

	http.SetCookie(writer,
		&http.Cookie{
			Name:    "token",
			Value:   token_string,
			Expires: time.Now().Add(time.Minute * 1),
		},
	)

	Helper.WriteOk(writer, login_response)
}

func (user_controller *UserController) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cookie, _ := request.Cookie("token")
	if cookie == nil {
		return
	}
	http.SetCookie(writer,
		&http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: time.Now(),
		},
	)

	Helper.WriteOk(writer, "Logout Sucessful")
}

func (user_controller *UserController) GetUserFromCookie(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cookie, err := request.Cookie("token")
	if err != nil {
		//TODO
		Helper.WriteUnAuthorized(writer)
		return
	}

	claims := JWT.Claims{}
	token, err := jwt.ParseWithClaims(cookie.Value, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return JWT.Jwtkey, nil
		})
	if err != nil {
		Helper.WriteBadRequest(writer)
		return
	}
	if !token.Valid {
		Helper.WriteUnAuthorized(writer)
		return
	}

	Helper.WriteOk(writer, claims)
}
