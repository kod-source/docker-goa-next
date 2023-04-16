package model

import "time"

type UserID uint64

type User struct {
	ID        UserID    `json:"id"`
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
	ID        UserID
	Name      string
	CreatedAt time.Time
	Avatar    *string
}

type GoogleUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Locale     string `json:"locale"`
	HD         string `json:"hd"`
}
