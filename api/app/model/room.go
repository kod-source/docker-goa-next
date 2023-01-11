package model

import "time"

type RoomID uint64

type Room struct {
	ID        RoomID
	Name      string
	IsGroup   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Img       *string
}

type RoomUser struct {
	Room  Room
	Users []*ShowUser
}

type IndexRoom struct {
	Room      Room
	IsOpen    bool
	LastText  *string
	CountUser int
	ShowImg   *string
}
