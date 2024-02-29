package main

import (
	"UserLoginAPI/Controller"
	"UserLoginAPI/DatabaseConnection"
	"UserLoginAPI/Middleware"
	"UserLoginAPI/Repository"
	"UserLoginAPI/Service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	_validator := validator.New(validator.WithRequiredStructEnabled())
	db := DatabaseConnection.OpenConnection()

	user_repository := Repository.NewUserRepository(db)
	user_service := Service.NewUserService(user_repository, _validator)
	user_controller := Controller.NewUserController(user_service)

	router := httprouter.New()

	router.POST("/user-login-api/login", user_controller.Login)
	router.POST("/user-login-api/register", user_controller.Register)
	router.GET("/user-login-api/getcookie", user_controller.GetUserFromCookie)
	router.GET("/user-login-api/logout", user_controller.Logout)
	router.POST("/user-login-api/admin-only-login", user_controller.AdminOnlyLogin)

	server := http.Server{
		Addr:    "localhost:4000",
		Handler: Middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
