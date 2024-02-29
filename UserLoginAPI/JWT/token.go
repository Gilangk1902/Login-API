package JWT

import (
	"UserLoginAPI/Model"
	"UserLoginAPI/Repository"
	"time"

	"github.com/golang-jwt/jwt"
)

var Jwtkey = []byte("secret_key")

type Claims struct {
	User_name string `json= "user_name"`
	Email     string `json="email`
	Role      string `json="role"`
	jwt.StandardClaims
}

func GenerateToken(user_response Model.User, user_repository Repository.IUserRepository) string {
	expires := time.Now().Add(time.Minute * 1)
	claims := Claims{
		User_name: user_response.User_name,
		Email:     user_response.Email,
		Role:      user_repository.GetRoleById(user_response.Role_id).Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err := token.SignedString(Jwtkey)
	if err != nil {
		panic(err)
	}

	return token_string
}
