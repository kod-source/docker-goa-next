package model

import "time"

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Img       *string   `json:"img"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IndexPost struct {
	Post Post
	User User
}

type ShowPost struct {
	IndexPost IndexPost
	Comments  []*Comment
	Likes []*Like
}
