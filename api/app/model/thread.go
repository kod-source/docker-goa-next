package model

import (
	"time"
)

type ThreadID uint64

type Thread struct {
	// ID ...
	ID ThreadID
	// UserID ...
	UserID UserID
	// RoomID ...
	RoomID RoomID
	// Text ...
	Text string
	// CreatedAt ...
	CreatedAt time.Time
	// UpdatedAt ...
	UpdatedAt time.Time
	// Img 画像データ
	Img *string
}

type ThreadUser struct {
	Thread Thread
	User   ShowUser
}

type IndexThread struct {
	ThreadUser   ThreadUser
	CountContent *int
}
