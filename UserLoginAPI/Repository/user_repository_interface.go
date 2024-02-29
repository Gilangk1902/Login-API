package Repository

import (
	"UserLoginAPI/Model"
)

type IUserRepository interface {
	GetUserFromEmail(email string) Model.User
	Register(user Model.User) Model.User
	GetRoleById(id int) Model.Role
}
