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
			Img: sql.NullString{
				String: "test1_img",
				Valid:  true,
			},
		},
		{
			ID:        2,
			Name:      "test2_room",
			IsGroup:   false,
			CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
	}
	if err := schema.InsertRoom(ctx, db, rooms...); err != nil {
		return err
	}
	return nil
}
