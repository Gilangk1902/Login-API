package Web

import "UserLoginAPI/Model"

type LoginResponse struct {
	User_name string     `json="user_name"`
	Email     string     `json="email"`
	Role      Model.Role `json="role`
}

type RegisterResponse struct {
	User_name string `json="user_name"`
	Role_id   int    `json="role_id"`
	Email     string `json="email"`
}
