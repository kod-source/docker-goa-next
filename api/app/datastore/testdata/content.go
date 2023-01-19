package testdata

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func ContentSeed(ctx context.Context, db *sql.DB) error {
	cs := []*schema.Content{
		{
			ID:        1,
			UserID:    1,
			ThreadID:  1,
			Text:      "content1",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "content1_img",
				Valid:  true,
			},
		},
		{
			ID:        2,
			UserID:    2,
			ThreadID:  1,
			Text:      "content2",
			CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			Img: sql.NullString{
				String: "content2_img",
				Valid:  true,
			},
		},
		{
			ID:        3,
			UserID:    1,
			ThreadID:  1,
			Text:      "content3",
			CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
		{
			ID:        4,
			UserID:    1,
			ThreadID:  2,
			Text:      "content4",
			CreatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 4, 1, 0, 0, 0, 0, jst),
			Img: sql.NullString{
				String: "content4_img",
				Valid:  true,
			},
		},
		{
			ID:        5,
			UserID:    2,
			ThreadID:  4,
			Text:      "content5",
			CreatedAt: time.Date(2022, 5, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 5, 1, 0, 0, 0, 0, jst),
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		},
	}

	if err := schema.InsertContent(ctx, db, cs...); err != nil {
		return err
	}
	return nil
}
