package Repository

import (
	"UserLoginAPI/DatabaseConnection"
	"fmt"
	"testing"
)

func TestLogin(t *testing.T) {
	db := DatabaseConnection.OpenConnection()
	user := NewUserRepository(db).GetUserFromEmail("ThiccG@gmail.com", DatabaseConnection.OpenConnection())

	fmt.Print(user.User_name)
}
