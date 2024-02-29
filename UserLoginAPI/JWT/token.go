package JWT

import "github.com/golang-jwt/jwt"

type Claims struct {
	User_name string `json= "user_name"`
	Email     string `json="email`
	Role      string `json="role"`
	jwt.StandardClaims
}
