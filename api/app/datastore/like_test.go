package datastore

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/schema"
)

func Test_CreateLike(t *testing.T) {
	tx, err := testDB.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	ld := NewLikeDatastore()

	t.Run("[OK]いいね登録", func(t *testing.T) {
		got, err := ld.Create(ctx, tx, 2, 2)
		if err != nil {
			t.Fatal(err)
		}

		like, err := schema.SelectLike(ctx, testDB, &schema.Like{ID: uint64(got.ID)})
		if err != nil {
			t.Fatal(err)
		}
		want := &model.Like{
			ID:     int(like.ID),
			PostID: int(like.PostID),
			UserID: int(like.UserID),
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := ld.Delete(ctx, tx, got.UserID, got.PostID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]いいね登録 - ユニークエラー", func(t *testing.T) {
		_, err := ld.Create(ctx, tx, 1, 1)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorDuplicate.Number {
			t.Errorf("want error code %v, but got error code is %v", myerrors.MySQLErrorDuplicate.Number, code)
		}
	})

	t.Run("[NG]いいね登録 - 存在しない投稿をいいね", func(t *testing.T) {
		_, err := ld.Create(ctx, tx, 1, 1000)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error code %v, but got error code is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})

	t.Run("[NG]いいね登録 - 存在しないユーザーでいいね", func(t *testing.T) {
		_, err := ld.Create(ctx, tx, 1000, 1)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error code %v, but got error code is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})
}

func Test_DeleteLike(t *testing.T) {
	tx, err := testDB.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	ld := NewLikeDatastore()

	t.Run("[OK]いいね削除", func(t *testing.T) {
		likeID := 4
		like := &schema.Like{
			ID:     uint64(likeID),
			UserID: 2,
			PostID: 2,
		}
		if err := schema.InsertLike(ctx, testDB, like); err != nil {
			t.Fatal(err)
		}

		if err := ld.Delete(ctx, tx, int(like.UserID), int(like.PostID)); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]いいね削除 - 存在しないものを選択", func(t *testing.T) {
		if err := ld.Delete(ctx, tx, 1000, 1000); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want err %v, but got error is %v", sql.ErrNoRows, err)
		}
	})
}

func Test_GetPostIDs(t *testing.T) {
	tx, err := testDB.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()
	ld := NewLikeDatastore()

	t.Run("[OK]いいねした投稿IDを取得", func(t *testing.T) {
		want := []int{1, 2, 3}
		got, err := ld.GetPostIDs(ctx, tx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]いいねした投稿IDを取得 - いいねした投稿が存在しない時", func(t *testing.T) {
		got, err := ld.GetPostIDs(ctx, tx, 1000)
		if err != nil {
			t.Fatal(err)
		}

		if len(got) != 0 {
			t.Errorf("want length is 0, but got length is %v", len(got))
		}
	})
}
