package testdata

import (
	"context"
	"database/sql"

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
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	if err := schema.InsertRoom(ctx, db, rooms...); err != nil {
		return err
	}
	return nil
}
