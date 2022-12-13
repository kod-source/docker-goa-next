package testdata

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/schema"
)

func UserSeed(ctx context.Context, db *sql.DB) error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	now := time.Date(2022, 1, 1, 0, 0, 0, 0, jst)

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
