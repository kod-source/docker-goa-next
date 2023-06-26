package interactor

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	mock_repository "github.com/kod-source/docker-goa-next/app/repository/mock"
	"github.com/shogo82148/pointer"
)

func Test_InviteRoom(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urr := mock_repository.NewMockUserRoomRepository(ctrl)
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
		urr.EXPECT().Create(ctx, wantRoomID, wantUserID).Return(want, nil)

		got, err := uru.InviteRoom(ctx, wantRoomID, wantUserID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルームに招待 - ルームIDが0の時", func(t *testing.T) {
		if _, err := uru.InviteRoom(ctx, model.RoomID(0), wantUserID); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error is want %v, got %v", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]ルームに招待 - ユーザーIDが0の時", func(t *testing.T) {
		if _, err := uru.InviteRoom(ctx, wantRoomID, model.UserID(0)); !errors.Is(err, myerrors.ErrBadRequestInt) {
			t.Errorf("error is want %v, got %v", myerrors.ErrBadRequestInt, err)
		}
	})

	t.Run("[NG]ルームに招待 - 想定外エラー発生", func(t *testing.T) {
		wantErr := errors.New("test error")
		urr.EXPECT().Create(ctx, wantRoomID, wantUserID).Return(nil, wantErr)

		if _, err := uru.InviteRoom(ctx, wantRoomID, wantUserID); !errors.Is(err, wantErr) {
			t.Errorf("error is want %v, got %v", wantErr, err)
		}
	})
}

func Test_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	urr := mock_repository.NewMockUserRoomRepository(ctrl)
	uru := NewUserRoomUsecase(urr)
	wantUserRoomID := model.UserRoomID(1)

	t.Run("[OK]UserRoomの削除", func(t *testing.T) {
		urr.EXPECT().Delete(ctx, wantUserRoomID).Return(nil)

		if err := uru.Delete(ctx, wantUserRoomID); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]UserRoomの削除 - 想定外エラー発生", func(t *testing.T) {
		wantErr := errors.New("test error")
		urr.EXPECT().Delete(ctx, wantUserRoomID).Return(wantErr)

		if err := uru.Delete(ctx, wantUserRoomID); !errors.Is(err, wantErr) {
			t.Errorf("error is want %v, got %v", wantErr, err)
		}
	})
}
