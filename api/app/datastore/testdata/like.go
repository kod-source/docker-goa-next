package testdata

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func LikeSeed(ctx context.Context, db *sql.DB) error {
	likes := []*schema.Like{
		{
			ID:     1,
			PostID: 1,
			UserID: 1,
		},
		{
			ID:     2,
			PostID: 2,
			UserID: 1,
		},
		{
			ID:     3,
			PostID: 3,
			UserID: 1,
		},
		{
			ID:     4,
			PostID: 1,
			UserID: 2,
		},
	}

	if err := schema.InsertLike(ctx, db, likes...); err != nil {
		return err
	}
	return nil
}
