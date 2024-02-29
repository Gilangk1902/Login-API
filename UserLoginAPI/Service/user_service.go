package Service

import (
	"UserLoginAPI/Helper"
	"UserLoginAPI/JWT"
	"UserLoginAPI/Model"
	"UserLoginAPI/Model/Web"
	"UserLoginAPI/Repository"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
)

var Jwtkey = []byte("secret_key")

type UserService struct {
	User_repository Repository.IUserRepository
	Validate        *validator.Validate
}

func (user_service *UserService) AdminOnlyLogin(request Web.LoginRequest) (Web.LoginResponse, string) {
	err := user_service.Validate.Struct(request)
	if err != nil {
		return Web.LoginResponse{}, Helper.FAIL_VALIDATION
	}

	user_response, token := user_service.Login(request)

	ADMIN := 2
	if user_response.Role.Id != ADMIN {
		return Web.LoginResponse{}, Helper.NO_TOKEN
	}
	return user_response, token
}

func NewUserService(user_repository Repository.IUserRepository, validator *validator.Validate) IUserService {
	return &UserService{
		User_repository: user_repository,
		Validate:        validator,
	}
}

func (user_service *UserService) Register(request Web.RegisterRequest) (Web.RegisterResponse, error) {
	err := user_service.Validate.Struct(request)
	if err != nil {
		return Web.RegisterResponse{}, errors.New(Helper.FAIL_VALIDATION)
	}

	user := Model.User{
		Email:     request.Email,
		User_name: request.User_name,
		Password:  Helper.HashPassword(request.Password),
		Role_id:   request.Role_id,
	}

	result := user_service.User_repository.Register(user)

	register_response := Web.RegisterResponse{
		Email:     result.Email,
		Role_id:   result.Role_id,
		User_name: result.User_name,
	}

	return register_response, nil
}

func (user_service *UserService) Login(request Web.LoginRequest) (Web.LoginResponse, string) {
	err := user_service.Validate.Struct(request)
	if err != nil {
		return Web.LoginResponse{}, Helper.FAIL_VALIDATION
	}

	user_response := user_service.User_repository.GetUserFromEmail(request.Email)

	if Helper.CheckPasswordHash(request.Password, user_response.Password) {
		expires := time.Now().Add(time.Minute * 1)
		claims := JWT.Claims{
			User_name: user_response.User_name,
			Email:     user_response.Email,
			Role:      user_service.User_repository.GetRoleById(user_response.Role_id).Name,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expires.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token_string, err := token.SignedString(Jwtkey)
		if err != nil {
			panic(err)
		}

		return Web.LoginResponse{
			User_name: user_response.User_name,
			Email:     user_response.Email,
			Role:      user_service.User_repository.GetRoleById(user_response.Role_id),
		}, token_string
	}
	return Web.LoginResponse{}, Helper.NO_TOKEN
}
