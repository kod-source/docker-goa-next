package datastore

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/schema"
)

func Test_DeleteContent(t *testing.T) {
	cr := NewContentRepository(testDB, nil)
	wantUserID := model.UserID(1)
	wantThreadID := model.ThreadID(2)

	t.Run("[OK]コンテントの削除", func(t *testing.T) {
		contentID := model.ContentID(6)
		if err := schema.InsertContent(ctx, testDB, &schema.Content{
			ID:        uint64(contentID),
			UserID:    uint64(wantUserID),
			ThreadID:  uint64(wantThreadID),
			Text:      "delete_text",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "delete_img",
				Valid:  true,
			},
		}); err != nil {
			t.Fatal(err)
		}

		if err := cr.Delete(ctx, wantUserID, contentID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]コンテントの削除 - 他人の投稿を削除した時", func(t *testing.T) {
		contentID := model.ContentID(7)
		if err := schema.InsertContent(ctx, testDB, &schema.Content{
			ID:        uint64(contentID),
			UserID:    uint64(wantUserID),
			ThreadID:  uint64(wantThreadID),
			Text:      "delete_text",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "delete_img",
				Valid:  true,
			},
		}); err != nil {
			t.Fatal(err)
		}

		if err := cr.Delete(ctx, 2, contentID); !errors.Is(err, myerrors.ErrBadRequestNoPermission) {
			t.Errorf("error mismatch (-want %v, +got %v)", myerrors.ErrBadRequestNoPermission, err)
		}
		if err := cr.Delete(ctx, wantUserID, contentID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]コンテントの削除 - 存在しないIDを指定した時", func(t *testing.T) {
		if err := cr.Delete(ctx, wantUserID, 1000); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("error mismatch (-want %v, +got %v)", sql.ErrNoRows, err)
		}
	})
}
