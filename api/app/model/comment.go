package model

import "time"

type Comment struct {
	ID        *int       `json:"id"`
	PostID    *int       `json:"post_id"`
	Text      *string    `json:"text"`
	Img       *string    `json:"img"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
