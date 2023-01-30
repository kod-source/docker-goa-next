package datastore

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/pointer"
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

func Test_CreateContent(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return now
	}
	cr := NewContentRepository(testDB, tr)
	wantUserID := model.UserID(1)
	wantThreadID := model.ThreadID(2)
	wantText := "create content"
	wantImg := "content img"

	t.Run("[OK]コンテントの作成", func(t *testing.T) {
		got, err := cr.Create(ctx, wantText, wantThreadID, wantUserID, &wantImg)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := cr.Delete(ctx, wantUserID, got.Content.ID); err != nil {
				t.Fatal(err)
			}
		}()

		want := &model.ContentUser{
			Content: model.Content{
				ID:        got.Content.ID,
				UserID:    wantUserID,
				ThreadID:  wantThreadID,
				Text:      wantText,
				CreatedAt: now,
				UpdatedAt: now,
				Img:       &wantImg,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      "test1_name",
				CreatedAt: now,
				Avatar:    pointer.Ptr("test1_avatar"),
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]コンテントの作成 - 画像が空の時", func(t *testing.T) {
		got, err := cr.Create(ctx, wantText, wantThreadID, model.UserID(2), nil)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := cr.Delete(ctx, model.UserID(2), got.Content.ID); err != nil {
				t.Fatal(err)
			}
		}()

		want := &model.ContentUser{
			Content: model.Content{
				ID:        got.Content.ID,
				UserID:    2,
				ThreadID:  wantThreadID,
				Text:      wantText,
				CreatedAt: now,
				UpdatedAt: now,
				Img:       nil,
			},
			User: model.ShowUser{
				ID:        2,
				Name:      "test2_name",
				CreatedAt: now,
				Avatar:    nil,
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]コンテントの作成 - 存在しないユーザーを指定した時", func(t *testing.T) {
		_, err := cr.Create(ctx, wantText, wantThreadID, model.UserID(1000), &wantImg)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("error code (-want %d, got %d)", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})

	t.Run("[NG]コンテントの作成 - 不明なスレッドIDを指定した時", func(t *testing.T) {
		_, err := cr.Create(ctx, wantText, model.ThreadID(1000), wantUserID, &wantImg)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("error code (-want %d, got %d)", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})
}
