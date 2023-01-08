package testdata

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func RoomSeed(ctx context.Context, db *sql.DB) error {
	rooms := []*schema.Room{
		{
			ID:        1,
			Name:      "test1_room",
			IsGroup:   true,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        2,
			Name:      "test2_room",
			IsGroup:   false,
			CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
		},
	}
	if err := schema.InsertRoom(ctx, db, rooms...); err != nil {
		return err
	}
	return nil
}
