package interactor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	datastore "github.com/kod-source/docker-goa-next/app/datastore/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/shogo82148/pointer"
)

func Test_CreateThread(t *testing.T) {
	tr := &datastore.MockThreadRepository{}
	tu := NewThreadUsecase(tr)
	wantText := "スレッド作成"
	wantRoomID := model.RoomID(2)
	wantUserID := model.UserID(3)
	wantImg := pointer.Ptr("test img")

	t.Run("[OK]スレッド作成", func(t *testing.T) {
		want := &model.ThreadUser{
			Thread: model.Thread{
				ID:        1,
				UserID:    wantUserID,
				RoomID:    wantRoomID,
				Text:      wantText,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       wantImg,
			},
			User: model.ShowUser{
				ID:        wantUserID,
				Name:      "test user",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("test avatar"),
			},
		}
		tr.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
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

			return want, nil
		}
		defer func() {
			tr.CreateFunc = nil
		}()

		got, err := tu.Create(ctx, wantText, wantRoomID, wantUserID, wantImg)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]スレッド作成 - テキストが空の時", func(t *testing.T) {
		if _, err := tu.Create(ctx, "", wantRoomID, wantUserID, wantImg); !errors.Is(err, myerrors.ErrBadRequestSting) {
			t.Errorf("error want %v, got %v", myerrors.ErrBadRequestSting, err)
		}
	})

	t.Run("[NG]スレッド作成 - ルームIDが0の時", func(t *testing.T) {
		if _, err := tu.Create(ctx, wantText, model.RoomID(0), wantUserID, wantImg); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error want %v, got %v", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]スレッド作成 - ユーザーIDが0の時", func(t *testing.T) {
		if _, err := tu.Create(ctx, wantText, wantRoomID, model.UserID(0), wantImg); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error want %v, got %v", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]スレッド作成 - 想定外エラー発生", func(t *testing.T) {
		wantErr := errors.New("test error")
		tr.CreateFunc = func(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
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

			return nil, wantErr
		}
		defer func() {
			tr.CreateFunc = nil
		}()

		if _, err := tu.Create(ctx, wantText, wantRoomID, wantUserID, wantImg); !errors.Is(err, wantErr) {
			t.Errorf("mismatch error want %v, got %v", wantErr, err)
		}
	})
}
