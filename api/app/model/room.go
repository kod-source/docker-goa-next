package model

import "time"

type RoomID uint64

type Room struct {
	ID        RoomID
	Name      string
	IsGroup   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoomUser struct {
	room  Room
	users []*ShowUser
}
