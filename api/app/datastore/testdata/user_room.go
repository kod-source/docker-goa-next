package testdata

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func UserRoomSeed(ctx context.Context, db *sql.DB) error {
	userRooms := []*schema.UserRoom{
		{
			ID:     1,
			UserID: 1,
			RoomID: 1,
			LastReadAt: sql.NullTime{
				Time:  now,
				Valid: true,
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:     2,
			UserID: 2,
			RoomID: 1,
			LastReadAt: sql.NullTime{
				Time:  now,
				Valid: false,
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:     3,
			UserID: 1,
			RoomID: 2,
			LastReadAt: sql.NullTime{
				Time:  now,
				Valid: false,
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:     4,
			UserID: 2,
			RoomID: 2,
			LastReadAt: sql.NullTime{
				Time:  now,
				Valid: true,
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	if err := schema.InsertUserRoom(ctx, db, userRooms...); err != nil {
		return err
	}
	return nil
}
