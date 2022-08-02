package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

var MockUser = User{
	ID:       1,
	Name:     "佐藤　太郎",
	Email:    "test@exmaple.com",
	Password: "Test-1234",
}
