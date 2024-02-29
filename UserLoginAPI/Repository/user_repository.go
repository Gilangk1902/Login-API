package Repository

import (
	"UserLoginAPI/Model"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		Db: db,
	}
}

func (user_repository *UserRepository) GetUserFromEmail(email string) Model.User {
	user := Model.User{
		Email: email,
	}
	var result Model.User
	user_repository.Db.Where("email = ?", user.Email).Find(&result)

	return result
}

func (user_repository *UserRepository) Register(user Model.User) Model.User {
	transaction := user_repository.Db.Begin()

	err := transaction.Create(user).Error
	if err != nil {
		transaction.Rollback()
	}

	transaction.Commit()
	return user
}

func (user_repository *UserRepository) GetRoleById(id int) Model.Role {
	var result Model.Role
	user_repository.Db.Where("id = ?", id).First(&result)

	return result
}
