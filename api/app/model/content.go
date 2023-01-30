package model

import (
	"time"
)

type ContentID uint64

// Content スレッドの返信
type Content struct {
	// ID ...
	ID ContentID
	// UserID ...
	UserID UserID
	// ThreadID ...
	ThreadID ThreadID
	// Text ...
	Text string
	// CreatedAt ...
	CreatedAt time.Time
	// UpdatedAt ...
	UpdatedAt time.Time
	// Img 画像データ
	Img *string
}

type ContentUser struct {
	Content Content
	User    ShowUser
}
