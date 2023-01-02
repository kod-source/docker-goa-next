package model

import (
	"time"
)

type UserRoomID uint64

type UserRoom struct {
	// ID ...
	ID UserRoomID
	// UserID ...
	UserID UserID
	// RoomID ...
	RoomID RoomID
	// LastReadAt このルームの最後に既読をつけた日
	LastReadAt *time.Time
	// CreatedAt ...
	CreatedAt time.Time
	// UpdatedAt ...
	UpdatedAt time.Time
}
