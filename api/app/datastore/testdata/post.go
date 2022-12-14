package testdata

import (
	"context"
	"database/sql"

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
			CreatedAt: now,
			UpdatedAt: now,
			Img:       sql.NullString{"test1_post_img", true},
		},
		{
			ID:        3,
			UserID:    2,
			Title:     "test3_title",
			CreatedAt: now,
			UpdatedAt: now,
			Img:       sql.NullString{"", false},
		},
	}
	if err := schema.InsertPost(ctx, db, posts...); err != nil {
		return err
	}
	return nil
}
