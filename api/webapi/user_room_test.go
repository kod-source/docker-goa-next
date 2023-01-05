package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/interactor/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/kod-source/docker-goa-next/webapi/app/test"
	"github.com/shogo82148/pointer"
)

func Test_InviteRoom(t *testing.T) {
	srv := testApp.srv
	uru := &mock.MockUserRoomUsecase{}
	ur := NewUserRoomController(srv, uru)
	wantRoomID := model.RoomID(1)
	wantUserID := model.UserID(2)

	t.Run("[OK]ルームに招待", func(t *testing.T) {
		uru.InviteRoomFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return &model.UserRoom{
				ID:         1,
				UserID:     wantUserID,
				RoomID:     wantRoomID,
				LastReadAt: pointer.Ptr(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)),
				CreatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			}, nil
		}
		defer func() {
			uru.InviteRoomFunc = nil
		}()

		want := &app.UserRoom{
			ID:         1,
			RoomID:     int(wantRoomID),
			UserID:     int(wantUserID),
			LastReadAt: pointer.Ptr(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)),
			CreatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
		}
		_, got := test.InviteRoomUserRoomsCreated(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
		})
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ルームに招待 - last_read_atがnilの時", func(t *testing.T) {
		uru.InviteRoomFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return &model.UserRoom{
				ID:         1,
				UserID:     wantUserID,
				RoomID:     wantRoomID,
				LastReadAt: nil,
				CreatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			}, nil
		}
		defer func() {
			uru.InviteRoomFunc = nil
		}()

		want := &app.UserRoom{
			ID:         1,
			RoomID:     int(wantRoomID),
			UserID:     int(wantUserID),
			LastReadAt: nil,
			CreatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
		}
		_, got := test.InviteRoomUserRoomsCreated(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
		})
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルームに招待 - 不明なIDを指定した時", func(t *testing.T) {
		uru.InviteRoomFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, myerrors.MySQLErrorAddOrUpdateForeignKey
		}
		defer func() {
			uru.InviteRoomFunc = nil
		}()

		test.InviteRoomUserRoomsBadRequest(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
		})
	})

	t.Run("[NG]ルームに招待 - 数字が0の時", func(t *testing.T) {
		uru.InviteRoomFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			return nil, myerrors.ErrBadRequestInt
		}
		defer func() {
			uru.InviteRoomFunc = nil
		}()

		test.InviteRoomUserRoomsBadRequest(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: 0,
			UserID: 0,
		})
	})

	t.Run("[NG]ルームに招待 - 権限がないとき", func(t *testing.T) {
		uru.InviteRoomFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, myerrors.ErrBadRequestNoPermission
		}
		defer func() {
			uru.InviteRoomFunc = nil
		}()

		test.InviteRoomUserRoomsBadRequest(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
		})
	})

	t.Run("[NG]ルームに招待 - エラー発生", func(t *testing.T) {
		uru.InviteRoomFunc = func(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
			if diff := cmp.Diff(wantRoomID, roomID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, userID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, errors.New("test error")
		}
		defer func() {
			uru.InviteRoomFunc = nil
		}()

		test.InviteRoomUserRoomsInternalServerError(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
		})
	})
}

func Test_DeleteUserRoom(t *testing.T) {
	srv := testApp.srv
	uru := &mock.MockUserRoomUsecase{}
	ur := NewUserRoomController(srv, uru)
	wantUserRoomID := model.UserRoomID(1)

	t.Run("[OK]UserRoomの削除", func(t *testing.T) {
		uru.DeleteFunc = func(ctx context.Context, id model.UserRoomID) error {
			if diff := cmp.Diff(wantUserRoomID, id); diff != "" {
				t.Errorf("mismatch (-want, +got)\n%s", diff)
			}
			return nil
		}
		defer func() {
			uru.DeleteFunc = nil
		}()

		test.DeleteUserRoomsOK(t, ctx, srv, ur, int(wantUserRoomID))
	})

	t.Run("[NG]UserRoomの削除 - 想定外エラー発生", func(t *testing.T) {
		uru.DeleteFunc = func(ctx context.Context, id model.UserRoomID) error {
			if diff := cmp.Diff(wantUserRoomID, id); diff != "" {
				t.Errorf("mismatch (-want, +got)\n%s", diff)
			}
			return errors.New("test error")
		}
		defer func() {
			uru.DeleteFunc = nil
		}()

		test.DeleteUserRoomsInternalServerError(t, ctx, srv, ur, int(wantUserRoomID))
	})
}
