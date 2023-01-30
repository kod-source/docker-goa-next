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
		defer thr.Delete(ctx, wantUserID, got.Thread.ID)

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
		defer thr.Delete(ctx, wantUserID, got.Thread.ID)

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

func Test_DeleteThread(t *testing.T) {
	t.Helper()
	thr := NewThreadRepository(testDB, nil)
	wantUserID := model.UserID(1)
	wantRoomID := model.RoomID(2)

	t.Run("[OK]スレッドの削除", func(t *testing.T) {
		threadID := model.ThreadID(42)
		if err := schema.InsertThread(ctx, testDB, &schema.Thread{
			ID:        uint64(threadID),
			UserID:    uint64(wantUserID),
			RoomID:    uint64(wantRoomID),
			Text:      "delete_thread",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		}); err != nil {
			t.Fatal(err)
		}

		if err := thr.Delete(ctx, wantUserID, threadID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]スレッドの削除 - 他人の投稿を削除した時", func(t *testing.T) {
		threadID := model.ThreadID(43)
		if err := schema.InsertThread(ctx, testDB, &schema.Thread{
			ID:        uint64(threadID),
			UserID:    uint64(wantUserID),
			RoomID:    uint64(wantRoomID),
			Text:      "delete_thread",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		}); err != nil {
			t.Fatal(err)
		}

		if err := thr.Delete(ctx, 2, threadID); !errors.Is(err, myerrors.ErrBadRequestNoPermission) {
			t.Errorf("error mismatch (-want %v, got %v)", myerrors.ErrBadRequestNoPermission, err)
		}
		if err := thr.Delete(ctx, wantUserID, threadID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]スレッドの削除 - 存在しないスレッドを指定した時", func(t *testing.T) {
		if err := thr.Delete(ctx, wantUserID, 1000); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("error mismatch (-want %v, got %v)", sql.ErrNoRows, err)
		}
	})
}

func Test_GetThreadsByRoom(t *testing.T) {
	thr := NewThreadRepository(testDB, nil)
	wantRoomID := model.RoomID(1)

	t.Run("[OK]ルーム内のスレッドを取得", func(t *testing.T) {
		want := []*model.IndexThread{
			{
				ThreadUser: model.ThreadUser{
					Thread: model.Thread{
						ID:        3,
						UserID:    2,
						RoomID:    1,
						Text:      "thread3",
						CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
						Img:       pointer.Ptr("image3"),
					},
					User: model.ShowUser{
						ID:        2,
						Name:      "test2_name",
						CreatedAt: now,
						Avatar:    nil,
					},
				},
				CountContent: nil,
			},
			{
				ThreadUser: model.ThreadUser{
					Thread: model.Thread{
						ID:        2,
						UserID:    1,
						RoomID:    1,
						Text:      "thread2",
						CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
						Img:       nil,
					},
					User: model.ShowUser{
						ID:        1,
						Name:      "test1_name",
						CreatedAt: now,
						Avatar:    pointer.Ptr("test1_avatar"),
					},
				},
				CountContent: pointer.Ptr(1),
			},
			{
				ThreadUser: model.ThreadUser{
					Thread: model.Thread{
						ID:        1,
						UserID:    1,
						RoomID:    1,
						Text:      "thread1",
						CreatedAt: now,
						UpdatedAt: now,
						Img:       pointer.Ptr("image1"),
					},
					User: model.ShowUser{
						ID:        1,
						Name:      "test1_name",
						CreatedAt: now,
						Avatar:    pointer.Ptr("test1_avatar"),
					},
				},
				CountContent: pointer.Ptr(3),
			},
		}
		got, nextID, err := thr.GetThreadsByRoom(ctx, wantRoomID, model.ThreadID(0))
		if err != nil {
			t.Fatal(err)
		}

		if nextID != nil {
			t.Errorf("mismatch nextID (-want %v +got %v)", nil, *nextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ルーム内のスレッドの取得 - NextIDを指定するとき", func(t *testing.T) {
		want := []*model.IndexThread{
			{
				ThreadUser: model.ThreadUser{
					Thread: model.Thread{
						ID:        2,
						UserID:    1,
						RoomID:    1,
						Text:      "thread2",
						CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
						Img:       nil,
					},
					User: model.ShowUser{
						ID:        1,
						Name:      "test1_name",
						CreatedAt: now,
						Avatar:    pointer.Ptr("test1_avatar"),
					},
				},
				CountContent: pointer.Ptr(1),
			},
			{
				ThreadUser: model.ThreadUser{
					Thread: model.Thread{
						ID:        1,
						UserID:    1,
						RoomID:    1,
						Text:      "thread1",
						CreatedAt: now,
						UpdatedAt: now,
						Img:       pointer.Ptr("image1"),
					},
					User: model.ShowUser{
						ID:        1,
						Name:      "test1_name",
						CreatedAt: now,
						Avatar:    pointer.Ptr("test1_avatar"),
					},
				},
				CountContent: pointer.Ptr(3),
			},
		}

		got, nextID, err := thr.GetThreadsByRoom(ctx, wantRoomID, model.ThreadID(1))
		if err != nil {
			t.Fatal(err)
		}

		if nextID != nil {
			t.Errorf("mismatch nextID (-want %v +got %v)", nil, *nextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ルーム内のスレッドの取得 - データが存在しないとき", func(t *testing.T) {
		got, nextID, err := thr.GetThreadsByRoom(ctx, wantRoomID, model.ThreadID(1000))
		if err != nil {
			t.Fatal(err)
		}

		if nextID != nil {
			t.Errorf("mismatch nextID (-want nil +got %v)", *nextID)
		}
		if len(got) != 0 {
			t.Errorf("mismatch data length (-want 0 +got %v))", len(got))
		}
	})

	t.Run("[OK]ルーム内のスレッドの取得 - 存在しないルームIDを取得したとき", func(t *testing.T) {
		if _, _, err := thr.GetThreadsByRoom(ctx, model.RoomID(1000), model.ThreadID(0)); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("error mismatch (-want %v +got %v)", sql.ErrNoRows, err)
		}
	})
}
