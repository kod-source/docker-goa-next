package testdata

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func CommentSeed(ctx context.Context, db *sql.DB) error {
	comments := []*schema.Comment{
		{
			ID:        1,
			PostID:    1,
			UserID:    1,
			Text:      "test1_comment",
			CreatedAt: now,
			UpdatedAt: now,
			Img:       sql.NullString{"test1_comment_img", true},
		},
		{
			ID:        2,
			PostID:    1,
			UserID:    2,
			Text:      "test2_comment",
			CreatedAt: now,
			UpdatedAt: now,
			Img:       sql.NullString{"test2_comment_img", true},
		},
		{
			ID:        3,
			PostID:    2,
			UserID:    1,
			Text:      "test3_comment",
			CreatedAt: now,
			UpdatedAt: now,
			Img:       sql.NullString{"", false},
		},
	}
	if err := schema.InsertComment(ctx, db, comments...); err != nil {
		return err
	}

	return nil
}
