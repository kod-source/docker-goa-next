package interactor

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/datastore/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/shogo82148/pointer"
)

func Test_InviteRoom(t *testing.T) {
	urr := &mock.MockUserRoomRepository{}
	uru := NewUserRoomUsecase(urr)
	wantRoomID := model.RoomID(1)
	wantUserID := model.UserID(2)

	t.Run("[OK]ルームに招待", func(t *testing.T) {
		want := &model.UserRoom{
			ID:         1,
			UserID:     wantUserID,
			RoomID:     wantRoomID,
			LastReadAt: pointer.Ptr(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)),
			CreatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
		}
		urr.CreateFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return want, nil
		}
		defer func() {
			urr.CreateFunc = nil
		}()

		got, err := uru.InviteRoom(ctx, wantRoomID, wantUserID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルームに招待 - ルームIDが0の時", func(t *testing.T) {
		if _, err := uru.InviteRoom(ctx, 0, wantUserID); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error is want %v, got %v", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]ルームに招待 - ユーザーIDが0の時", func(t *testing.T) {
		if _, err := uru.InviteRoom(ctx, wantRoomID, 0); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error is want %v, got %v", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]ルームに招待 - 想定外エラー発生", func(t *testing.T) {
		wantErr := errors.New("test error")
		urr.CreateFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, wantErr
		}
		defer func() {
			urr.CreateFunc = nil
		}()

		if _, err := uru.InviteRoom(ctx, wantRoomID, wantUserID); !errors.Is(err, wantErr) {
			t.Errorf("error is want %v, got %v", wantErr, err)
		}
	})
}
