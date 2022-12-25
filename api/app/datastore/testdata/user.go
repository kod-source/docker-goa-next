package testdata

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func UserSeed(ctx context.Context, db *sql.DB) error {
	users := []*schema.User{
		{
			ID:        1,
			Name:      "test1_name",
			Email:     "test1@gmail.com",
			Password:  "test1_passowrd",
			CreatedAt: now,
			UpdatedAt: now,
			Avatar:    sql.NullString{"test1_avatar", true},
		},
		{
			ID:        2,
			Name:      "test2_name",
			Email:     "test2@gmail.com",
			Password:  "test2_passowrd",
			CreatedAt: now,
			UpdatedAt: now,
			Avatar:    sql.NullString{"", false},
		},
	}
	if err := schema.InsertUser(ctx, db, users...); err != nil {
		return err
	}
	return nil
}
