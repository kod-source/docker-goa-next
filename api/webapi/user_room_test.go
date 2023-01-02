package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/interactor/mock"
	"github.com/kod-source/docker-goa-next/app/model"
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

	t.Run("[NG]ルームに招待 - ルームIDが0の時", func(t *testing.T) {
		test.InviteRoomUserRoomsBadRequest(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: 0,
			UserID: int(wantUserID),
		})
	})

	t.Run("[NG]ルームに招待 - ユーザーIDが0の時", func(t *testing.T) {
		test.InviteRoomUserRoomsBadRequest(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: int(wantRoomID),
			UserID: 0,
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

		test.InviteRoomUserRoomsInternalServerError(t, ctx, srv, ur, &app.InviteRoomUserRoomsPayload{
			RoomID: int(wantRoomID),
			UserID: int(wantUserID),
		})
	})
}
