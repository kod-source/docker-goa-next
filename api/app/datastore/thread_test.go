package datastore

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/pointer"
)

func Test_CreateThread(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return now
	}
	defer func() {
		tr.NowFunc = nil
	}()
	thr := NewThreadRepository(testDB, tr)
	wantText := "スレッド"
	wantRoomID := model.RoomID(1)
	wantUserID := model.UserID(2)
	wantImg := pointer.Ptr("create_thread_img")

	t.Run("[OK]スレッドの作成", func(t *testing.T) {
		got, err := thr.Create(ctx, wantText, wantRoomID, wantUserID, wantImg)
		if err != nil {
			t.Fatal(err)
		}
		thread, err := schema.SelectThread(ctx, testDB, &schema.Thread{ID: uint64(got.Thread.ID)})
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{ID: uint64(got.User.ID)})
		if err != nil {
			t.Fatal(err)
		}

		want := &model.ThreadUser{
			Thread: model.Thread{
				ID:        model.ThreadID(thread.ID),
				UserID:    wantUserID,
				RoomID:    wantRoomID,
				Text:      wantText,
				CreatedAt: now,
				UpdatedAt: now,
				Img:       wantImg,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
				Avatar:    pointer.PtrOrNil(user.Avatar.String),
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]スレッド作成 - 画像なしの時", func(t *testing.T) {
		got, err := thr.Create(ctx, wantText, wantRoomID, wantUserID, nil)
		if err != nil {
			t.Fatal(err)
		}
		thread, err := schema.SelectThread(ctx, testDB, &schema.Thread{ID: uint64(got.Thread.ID)})
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{ID: uint64(got.User.ID)})
		if err != nil {
			t.Fatal(err)
		}

		want := &model.ThreadUser{
			Thread: model.Thread{
				ID:        model.ThreadID(thread.ID),
				UserID:    wantUserID,
				RoomID:    wantRoomID,
				Text:      wantText,
				CreatedAt: now,
				UpdatedAt: now,
				Img:       nil,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      user.Name,
				CreatedAt: user.CreatedAt,
				Avatar:    pointer.PtrOrNil(user.Avatar.String),
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]スレッド作成 - 不明なルームIDの時", func(t *testing.T) {
		_, err := thr.Create(ctx, wantText, 1000, wantUserID, wantImg)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("error code (-want %d, got %d)", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})

	t.Run("[NG]スレッド作成 - 不明なユーザーIDの時", func(t *testing.T) {
		_, err := thr.Create(ctx, wantText, wantRoomID, 1000, wantImg)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("error code (-want %d, got %d)", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})
}
