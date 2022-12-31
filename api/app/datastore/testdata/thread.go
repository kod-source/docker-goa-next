package testdata

import (
	"context"
	"database/sql"
	"time"

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
			CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
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
			CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
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
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, jst),
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
			CreatedAt: time.Date(2022, 5, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 5, 1, 0, 0, 0, 0, jst),
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
