package main

import (
	"context"
	"database/sql"
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

		want := &app.RoomUser{
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

func Test_IndexRoom(t *testing.T) {
	srv := testApp.srv
	ru := &mock.MockRoomUsecase{}
	r := NewRoomController(srv, ru)
	wantUserID := model.UserID(1)
	wantNextID := model.RoomID(10)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantUserID))

	t.Run("[OK]ルーム表示", func(t *testing.T) {
		gotNextID := 20
		irs := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        1,
					Name:      "room_1",
					IsGroup:   true,
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				},
				IsOpen:    true,
				LastText:  "test_text1",
				CountUser: 10,
			},
			{
				Room: model.Room{
					ID:        2,
					Name:      "room_2",
					IsGroup:   false,
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				},
				IsOpen:    false,
				LastText:  "",
				CountUser: 2,
			},
		}
		ru.IndexFunc = func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantNextID, nextID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			return irs, pointer.Ptr(gotNextID), nil
		}
		defer func() {
			ru.IndexFunc = nil
		}()

		want := &app.AllRoomUser{
			NextID: pointer.Ptr(gotNextID),
			IndexRoom: app.IndexRoomCollection{
				{
					Room: &app.Room{
						ID:        1,
						Name:      "room_1",
						IsGroup:   true,
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					},
					IsOpen:    true,
					LastText:  pointer.Ptr("test_text1"),
					CountUser: 10,
				},
				{
					Room: &app.Room{
						ID:        2,
						Name:      "room_2",
						IsGroup:   false,
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					},
					IsOpen:    false,
					LastText:  nil,
					CountUser: 2,
				},
			},
		}

		_, got := test.IndexRoomsOK(t, ctx, srv, r, pointer.Ptr(int(wantNextID)))

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want got)\n%s", diff)
		}
	})

	t.Run("[OK]ルーム表示 - NextIDがnilの時", func(t *testing.T) {
		irs := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        1,
					Name:      "room_1",
					IsGroup:   true,
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				},
				IsOpen:    true,
				LastText:  "test_text1",
				CountUser: 20,
			},
			{
				Room: model.Room{
					ID:        2,
					Name:      "room_2",
					IsGroup:   false,
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				},
				IsOpen:    false,
				LastText:  "",
				CountUser: 2,
			},
		}
		ru.IndexFunc = func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if nextID != 0 {
				t.Errorf("want nextID is 0, got is %d", nextID)
			}
			return irs, nil, nil
		}
		defer func() {
			ru.IndexFunc = nil
		}()

		want := &app.AllRoomUser{
			NextID: nil,
			IndexRoom: app.IndexRoomCollection{
				{
					Room: &app.Room{
						ID:        1,
						Name:      "room_1",
						IsGroup:   true,
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					},
					IsOpen:    true,
					LastText:  pointer.Ptr("test_text1"),
					CountUser: 20,
				},
				{
					Room: &app.Room{
						ID:        2,
						Name:      "room_2",
						IsGroup:   false,
						CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
						UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					},
					IsOpen:    false,
					LastText:  nil,
					CountUser: 2,
				},
			},
		}

		_, got := test.IndexRoomsOK(t, ctx, srv, r, nil)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want got)\n%s", diff)
		}
	})

	t.Run("[NG]ルーム表示 - ルームが存在しない時", func(t *testing.T) {
		ru.IndexFunc = func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantNextID, nextID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			return nil, nil, sql.ErrNoRows
		}
		defer func() {
			ru.IndexFunc = nil
		}()

		test.IndexRoomsNotFound(t, ctx, srv, r, pointer.Ptr(int(wantNextID)))
	})

	t.Run("[NG]ルーム表示 - 想定外エラー", func(t *testing.T) {
		ru.IndexFunc = func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantNextID, nextID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			return nil, nil, errors.New("test error")
		}
		defer func() {
			ru.IndexFunc = nil
		}()

		test.IndexRoomsInternalServerError(t, ctx, srv, r, pointer.Ptr(int(wantNextID)))
	})
}

func Test_Exists(t *testing.T) {
	srv := testApp.srv
	ru := &mock.MockRoomUsecase{}
	r := NewRoomController(srv, ru)
	wantMyUserID := model.UserID(1)
	wantUserID := model.UserID(2)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantMyUserID))

	t.Run("[OK]DMの存在を確認する", func(t *testing.T) {
		ru.ExistsFunc = func(ctx context.Context, myID, id model.UserID) (*model.Room, error) {
			if diff := cmp.Diff(wantMyUserID, myID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return &model.Room{
				ID:        1,
				Name:      "DB room",
				IsGroup:   false,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			}, nil
		}

		want := &app.Room{
			ID:        1,
			Name:      "DB room",
			IsGroup:   false,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
		}

		_, got := test.ExistsRoomsOK(t, ctx, srv, r, int(wantUserID))
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want got)\n%s", diff)
		}
	})

	t.Run("[NG]DMの存在を確認する - 404エラーの時", func(t *testing.T) {
		ru.ExistsFunc = func(ctx context.Context, myID, id model.UserID) (*model.Room, error) {
			if diff := cmp.Diff(wantMyUserID, myID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return nil, sql.ErrNoRows
		}

		test.ExistsRoomsNotFound(t, ctx, srv, r, int(wantUserID))
	})

	t.Run("[NG]DMの存在を確認する - 想定外エラー発生", func(t *testing.T) {
		ru.ExistsFunc = func(ctx context.Context, myID, id model.UserID) (*model.Room, error) {
			if diff := cmp.Diff(wantMyUserID, myID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return nil, errors.New("test error")
		}

		test.ExistsRoomsInternalServerError(t, ctx, srv, r, int(wantUserID))
	})
}

func Test_Show(t *testing.T) {
	srv := testApp.srv
	ru := &mock.MockRoomUsecase{}
	r := NewRoomController(srv, ru)
	wantRoomID := model.RoomID(1)
	wantMyUserID := model.UserID(2)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantMyUserID))

	t.Run("[OK]ルームの詳細を取得", func(t *testing.T) {
		roomUser := &model.RoomUser{
			Room: model.Room{
				ID:        wantRoomID,
				Name:      "test room",
				IsGroup:   false,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			},
			Users: []*model.ShowUser{
				{
					ID:        wantMyUserID,
					Name:      "test user 1",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test avatar"),
				},
				{
					ID:        3,
					Name:      "test user 3",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test3 avatar"),
				},
			},
		}
		ru.ShowFunc = func(ctx context.Context, id model.RoomID, userID model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantMyUserID, userID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return roomUser, nil
		}
		defer func() {
			ru.ShowFunc = nil
		}()

		want := &app.RoomUser{
			ID:        int(wantRoomID),
			IsGroup:   false,
			Name:      "test room",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Users: []*app.ShowUser{
				{
					ID:        int(wantMyUserID),
					Name:      "test user 1",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test avatar"),
				},
				{
					ID:        3,
					Name:      "test user 3",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test3 avatar"),
				},
			},
		}

		_, got := test.ShowRoomsOK(t, ctx, srv, r, int(wantRoomID))

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ルームの詳細を取得 - ユーザーがたくさんいるケース", func(t *testing.T) {
		roomUser := &model.RoomUser{
			Room: model.Room{
				ID:        wantRoomID,
				Name:      "test room",
				IsGroup:   true,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			},
			Users: []*model.ShowUser{
				{
					ID:        wantMyUserID,
					Name:      "test user 1",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test avatar"),
				},
				{
					ID:        3,
					Name:      "test user 3",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test3 avatar"),
				},
				{
					ID:        4,
					Name:      "test user 4",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    nil,
				},
				{
					ID:        5,
					Name:      "test user 5",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test5 avatar"),
				},
				{
					ID:        6,
					Name:      "test user 6",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    nil,
				},
				{
					ID:        7,
					Name:      "test user 7",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test7 avatar"),
				},
			},
		}
		ru.ShowFunc = func(ctx context.Context, id model.RoomID, userID model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantMyUserID, userID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return roomUser, nil
		}
		defer func() {
			ru.ShowFunc = nil
		}()

		want := &app.RoomUser{
			ID:        int(wantRoomID),
			IsGroup:   true,
			Name:      "test room",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Users: []*app.ShowUser{
				{
					ID:        int(wantMyUserID),
					Name:      "test user 1",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test avatar"),
				},
				{
					ID:        3,
					Name:      "test user 3",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test3 avatar"),
				},
				{
					ID:        4,
					Name:      "test user 4",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    nil,
				},
				{
					ID:        5,
					Name:      "test user 5",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test5 avatar"),
				},
				{
					ID:        6,
					Name:      "test user 6",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    nil,
				},
				{
					ID:        7,
					Name:      "test user 7",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test7 avatar"),
				},
			},
		}

		_, got := test.ShowRoomsOK(t, ctx, srv, r, int(wantRoomID))

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルームの詳細を取得 - ルームが存在しない時", func(t *testing.T) {
		ru.ShowFunc = func(ctx context.Context, id model.RoomID, userID model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantMyUserID, userID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return nil, sql.ErrNoRows
		}
		defer func() {
			ru.ShowFunc = nil
		}()

		test.ShowRoomsNotFound(t, ctx, srv, r, int(wantRoomID))
	})

	t.Run("[NG]ルームの詳細を取得 - ルームの閲覧権限がない時", func(t *testing.T) {
		ru.ShowFunc = func(ctx context.Context, id model.RoomID, userID model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantMyUserID, userID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return nil, myerrors.ErrBadRequestNoPermission
		}
		defer func() {
			ru.ShowFunc = nil
		}()

		test.ShowRoomsBadRequest(t, ctx, srv, r, int(wantRoomID))
	})

	t.Run("[NG]ルームの詳細を取得 - 想定外エラー発生", func(t *testing.T) {
		ru.ShowFunc = func(ctx context.Context, id model.RoomID, userID model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantMyUserID, userID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			return nil, errors.New("test error")
		}
		defer func() {
			ru.ShowFunc = nil
		}()

		test.ShowRoomsInternalServerError(t, ctx, srv, r, int(wantRoomID))
	})
}
