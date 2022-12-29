package testdata

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func ThreadSeed(ctx context.Context, db *sql.DB) error {
	threads := []*schema.Thread{
		{
			ID:        1,
			UserID:    1,
			RoomID:    1,
			Text:      "thread1",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "image1",
				Valid:  true,
			},
		},
		{
			ID:        2,
			UserID:    1,
			RoomID:    1,
			Text:      "thread2",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			ID:        3,
			UserID:    2,
			RoomID:    1,
			Text:      "thread3",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "image3",
				Valid:  true,
			},
		},
		{
			ID:        4,
			UserID:    1,
			RoomID:    2,
			Text:      "thread4",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "image4",
				Valid:  true,
			},
		},
		{
			ID:        5,
			UserID:    2,
			RoomID:    2,
			Text:      "thread5",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
	}
	if err := schema.InsertThread(ctx, db, threads...); err != nil {
		return err
	}
	return nil
}
