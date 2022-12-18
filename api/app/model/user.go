package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	Avatar    *string   `json:"avatar"`
}

type UserNil struct {
	ID        *int       `json:"id"`
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	Avatar    *string    `json:"avatar"`
}

type ShowUser struct {
	ID        int
	Name      string
	CreatedAt time.Time
	Avatar    *string
}
