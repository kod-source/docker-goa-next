package main

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	interactor "github.com/kod-source/docker-goa-next/app/interactor/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/kod-source/docker-goa-next/webapi/app/test"
	"github.com/shogo82148/pointer"
)

func Test_CreateThread(t *testing.T) {
	srv := testApp.srv
	tu := &interactor.MockThreadUsecase{}
	tc := NewThreadController(srv, tu)
	wantText := "テストスレ"
	wantThreadID := model.ThreadID(1)
	wantRoomID := model.RoomID(2)
	wantUserID := model.UserID(3)
	wantImg := pointer.Ptr("test img")

	t.Run("[OK]スレッド作成", func(t *testing.T) {
		threadUser := &model.ThreadUser{
			Thread: model.Thread{
				ID:        wantThreadID,
				UserID:    wantUserID,
				RoomID:    wantRoomID,
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       wantImg,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      "test_user",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("test_avatar"),
			},
		}
		tu.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return threadUser, nil
		}
		defer func() {
			tu.CreateFunc = nil
		}()

		want := &app.ThreadUser{
			Thread: &app.Thread{
				ID:        int(wantThreadID),
				RoomID:    int(wantRoomID),
				UserID:    int(wantUserID),
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       wantImg,
			},
			User: &app.ShowUser{
				ID:        int(wantUserID),
				Name:      "test_user",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("test_avatar"),
			},
		}

		_, got := test.CreateThreadsCreated(t, ctx, srv, tc, &app.CreateThreadsPayload{
			Text:   wantText,
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
			Img:    wantImg,
		})

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]スレッド作成 - 画像なし", func(t *testing.T) {
		threadUser := &model.ThreadUser{
			Thread: model.Thread{
				ID:        wantThreadID,
				UserID:    wantUserID,
				RoomID:    wantRoomID,
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       nil,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      "test_user",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    nil,
			},
		}
		tu.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if img != nil {
				t.Errorf("img (want nil, got %v)", *img)
			}

			return threadUser, nil
		}
		defer func() {
			tu.CreateFunc = nil
		}()

		want := &app.ThreadUser{
			Thread: &app.Thread{
				ID:        int(wantThreadID),
				RoomID:    int(wantRoomID),
				UserID:    int(wantUserID),
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       nil,
			},
			User: &app.ShowUser{
				ID:        int(wantUserID),
				Name:      "test_user",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    nil,
			},
		}

		_, got := test.CreateThreadsCreated(t, ctx, srv, tc, &app.CreateThreadsPayload{
			Text:   wantText,
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
			Img:    nil,
		})

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]スレッド作成 - 不明な文字列の時", func(t *testing.T) {
		tu.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.ErrBadRequestSting
		}
		defer func() {
			tu.CreateFunc = nil
		}()

		test.CreateThreadsBadRequest(t, ctx, srv, tc, &app.CreateThreadsPayload{
			Text:   wantText,
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
			Img:    wantImg,
		})
	})

	t.Run("[NG]スレッド作成 - 不明な数字の時", func(t *testing.T) {
		tu.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.ErrBadRequestInt
		}
		defer func() {
			tu.CreateFunc = nil
		}()

		test.CreateThreadsBadRequest(t, ctx, srv, tc, &app.CreateThreadsPayload{
			Text:   wantText,
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
			Img:    wantImg,
		})
	})

	t.Run("[NG]スレッド作成 - 外部キーエラーの時", func(t *testing.T) {
		tu.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.MySQLErrorAddOrUpdateForeignKey
		}
		defer func() {
			tu.CreateFunc = nil
		}()

		test.CreateThreadsBadRequest(t, ctx, srv, tc, &app.CreateThreadsPayload{
			Text:   wantText,
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
			Img:    wantImg,
		})
	})

	t.Run("[NG]スレッド作成 - 想定外エラー", func(t *testing.T) {
		tu.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
			if diff := cmp.Diff(wantText, text); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantImg, img); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, errors.New("test error")
		}
		defer func() {
			tu.CreateFunc = nil
		}()

		test.CreateThreadsInternalServerError(t, ctx, srv, tc, &app.CreateThreadsPayload{
			Text:   wantText,
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
			Img:    wantImg,
		})
	})
}

func Test_DeleteThread(t *testing.T) {
	srv := testApp.srv
	tu := &interactor.MockThreadUsecase{}
	tc := NewThreadController(srv, tu)
	wantMyID := model.UserID(1)
	wantThreadID := model.ThreadID(2)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantMyID))

	t.Run("[OK]スレッドの削除", func(t *testing.T) {
		tu.DeleteFunc = func(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
			if diff := cmp.Diff(wantMyID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil
		}

		test.DeleteThreadsOK(t, ctx, srv, tc, int(wantThreadID))
	})

	t.Run("[NG]スレッドの削除 - IDが不明の時", func(t *testing.T) {
		tu.DeleteFunc = func(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
			if diff := cmp.Diff(wantMyID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(model.ThreadID(1000), threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return sql.ErrNoRows
		}

		test.DeleteThreadsNotFound(t, ctx, srv, tc, 1000)
	})

	t.Run("[NG]スレッドの削除 - 権限エラー", func(t *testing.T) {
		tu.DeleteFunc = func(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
			if diff := cmp.Diff(wantMyID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return myerrors.ErrBadRequestNoPermission
		}

		test.DeleteThreadsBadRequest(t, ctx, srv, tc, int(wantThreadID))
	})

	t.Run("[NG]スレッドの削除 - 想定外エラー発生", func(t *testing.T) {
		tu.DeleteFunc = func(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
			if diff := cmp.Diff(wantMyID, myID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantThreadID, threadID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return errors.New("test error")
		}

		test.DeleteThreadsInternalServerError(t, ctx, srv, tc, int(wantThreadID))
	})
}