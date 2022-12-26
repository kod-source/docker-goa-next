package main

import (
	"context"
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

func Test_CreateRoom(t *testing.T) {
	srv := testApp.srv
	ru := &mock.MockRoomUsecase{}
	r := NewRoomController(srv, ru)
	wantRoomName := "room_name"
	wantUserIDs := []model.UserID{1, 2}
	wantIsGroup := true

	t.Run("[OK]ルーム作成", func(t *testing.T) {
		roomUser := &model.RoomUser{
			Room: model.Room{
				ID:        1,
				Name:      wantRoomName,
				IsGroup:   wantIsGroup,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			},
			Users: []*model.ShowUser{
				{
					ID:        wantUserIDs[0],
					Name:      "test1_user",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test1_avatar"),
				},
				{
					ID:        wantUserIDs[1],
					Name:      "test2_user",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    nil,
				},
			},
		}
		ru.CreateFunc = func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserIDs, userIDs); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantIsGroup, isGroup); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return roomUser, nil
		}

		want := &app.IndexRooUser{
			ID:        1,
			Name:      wantRoomName,
			IsGroup:   wantIsGroup,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Users: []*app.ShowUser{
				{
					ID:        1,
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Name:      "test1_user",
					Avatar:    pointer.Ptr("test1_avatar"),
				},
				{
					ID:        2,
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Name:      "test2_user",
					Avatar:    nil,
				},
			},
		}

		_, got := test.CreateRoomRoomsCreated(t, ctx, srv, r, &app.CreateRoomRoomsPayload{
			IsGroup: wantIsGroup,
			Name:    wantRoomName,
			UserIds: []int{1, 2},
		})

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルーム作成 - UserIDが空のケース", func(t *testing.T) {
		ru.CreateFunc = func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantIsGroup, isGroup); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if len(userIDs) != 0 {
				t.Errorf("want length is 0 but got length is %v", len(userIDs))
			}

			return nil, myerrors.ErrBadRequestEmptyArray
		}

		test.CreateRoomRoomsBadRequest(t, ctx, srv, r, &app.CreateRoomRoomsPayload{
			IsGroup: wantIsGroup,
			Name:    wantRoomName,
			UserIds: []int{},
		})
	})
}
