package Web

type LoginRequest struct {
	Email    string `validate:"required,email" json = "email"`
	Password string `validate:"required" json = "password"`
}

type RegisterRequest struct {
	User_name string `validate:"required" json="user_name"`
	Password  string `validate:"required" json="password"`
	Role_id   int    `validate:"required" json="role_id"`
	Email     string `validate:"required,email" json="email"`
}
