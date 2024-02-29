package Model

type User struct {
	User_name string
	Email     string
	Password  string
	Role_id   int
}

type Role struct {
	Id   int
	Name string
}
