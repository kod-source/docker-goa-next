package model

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

var MockUser = User{
	ID:       1,
	Name:     "佐藤　太郎",
	Email:    "test@exmaple.com",
	Password: "Test-1234",
}
