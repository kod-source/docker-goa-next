package testdata

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func PostSeed(ctx context.Context, db *sql.DB) error {
	posts := []*schema.Post{
		{
			ID:        1,
			UserID:    1,
			Title:     "test1_title",
			CreatedAt: now,
			UpdatedAt: now,
			Img:       sql.NullString{"test1_post_img", true},
		},
		{
			ID:        2,
			UserID:    1,
			Title:     "test2_title",
			CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			Img:       sql.NullString{"test1_post_img", true},
		},
		{
			ID:        3,
			UserID:    2,
			Title:     "test3_title",
			CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			Img:       sql.NullString{"", false},
		},
	}
	if err := schema.InsertPost(ctx, db, posts...); err != nil {
		return err
	}
	return nil
}
