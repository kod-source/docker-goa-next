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
			ID:     1,
			PostID: 2,
			UserID: 1,
		},
	}

	if err := schema.InsertLike(ctx, db, likes...); err != nil {
		return err
	}
	return nil
}
