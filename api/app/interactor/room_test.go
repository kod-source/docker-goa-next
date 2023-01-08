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

func Test_CreateRoom(t *testing.T) {
	rr := &mock.MockRoomRepository{}
	ri := NewRoomInterractor(rr)
	wantRoomName := "test_room"
	wantIsGroup := true
	wantUserIDs := []model.UserID{1, 2}

	t.Run("[OK]ルーム作成", func(t *testing.T) {
		want := &model.RoomUser{
			Room: model.Room{
				ID:        1,
				Name:      wantRoomName,
				IsGroup:   wantIsGroup,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       pointer.Ptr("test img"),
			},
			Users: []*model.ShowUser{
				{
					ID:        1,
					Name:      "user_1",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    pointer.Ptr("test1_avatar"),
				},
				{
					ID:        2,
					Name:      "user_2",
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Avatar:    nil,
				},
			},
		}
		rr.CreateFunc = func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantIsGroup, isGroup); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserIDs, userIDs); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return want, nil
		}

		got, err := ri.Create(ctx, wantRoomName, wantIsGroup, wantUserIDs)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルーム作成 - DM作成時にUserIDの数が違う時", func(t *testing.T) {
		rr.CreateFunc = func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(false, isGroup); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, nil
		}

		if _, err := ri.Create(ctx, wantRoomName, false, []model.UserID{1, 2, 3, 4}); !errors.Is(err, myerrors.ErrBadRequestSting) {
			t.Errorf("want error is %v, but got error is %v", myerrors.ErrBadRequestSting, err)
		}
	})

	t.Run("[NG]ルーム作成 - UserIDがない時", func(t *testing.T) {
		rr.CreateFunc = func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantIsGroup, isGroup); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, nil
		}

		if _, err := ri.Create(ctx, wantRoomName, wantIsGroup, []model.UserID{}); !errors.Is(err, myerrors.ErrBadRequestEmptyArray) {
			t.Errorf("want error is %v, but got error is %v", myerrors.ErrBadRequestEmptyArray, err)
		}
	})

	t.Run("[NG]ルーム作成 - Datastoreでエラー発生", func(t *testing.T) {
		rr.CreateFunc = func(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantIsGroup, isGroup); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserIDs, userIDs); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, myerrors.MySQLErrorDuplicate
		}

		if _, err := ri.Create(ctx, wantRoomName, wantIsGroup, wantUserIDs); !errors.Is(err, myerrors.MySQLErrorDuplicate) {
			t.Errorf("want error is %v, but got error is %v", myerrors.MySQLErrorDuplicate, err)
		}
	})
}

func Test_IndexRoom(t *testing.T) {
	rr := &mock.MockRoomRepository{}
	ri := NewRoomInterractor(rr)
	wantUserID := model.UserID(1)
	wantNextID := model.RoomID(10)

	t.Run("[OK]全てのルーム取得", func(t *testing.T) {
		gotNextID := 20
		want := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        1,
					Name:      "room_1",
					IsGroup:   true,
					CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
					Img:       pointer.Ptr("test img"),
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
					Img:       nil,
				},
				IsOpen:    false,
				LastText:  "test_text2",
				CountUser: 2,
			},
		}
		rr.IndexFunc = func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantNextID, nextID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return want, &gotNextID, nil
		}
		defer func() {
			rr.IndexFunc = nil
		}()

		got, nextID, err := ri.Index(ctx, wantUserID, wantNextID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(&gotNextID, nextID); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]全てのルーム取得 - 想定外エラー", func(t *testing.T) {
		gotErr := errors.New("test error")
		rr.IndexFunc = func(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantNextID, nextID); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, nil, gotErr
		}
		defer func() {
			rr.IndexFunc = nil
		}()

		if _, _, err := ri.Index(ctx, wantUserID, wantNextID); !errors.Is(err, gotErr) {
			t.Errorf("want error is %v, but got error is %v", gotErr, err)
		}
	})
}

func Test_ExistsRoom(t *testing.T) {
	rr := &mock.MockRoomRepository{}
	ri := NewRoomInterractor(rr)
	wantMyUserID := model.UserID(1)
	wantUserID := model.UserID(2)

	t.Run("[OK]DMの存在を確認", func(t *testing.T) {
		want := &model.Room{
			ID:        1,
			Name:      "DM room",
			IsGroup:   false,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Img:       nil,
		}
		rr.GetNoneGroupFunc = func(ctx context.Context, myID, id model.UserID) (*model.Room, error) {
			if diff := cmp.Diff(wantMyUserID, myID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			return want, nil
		}

		got, err := ri.Exists(ctx, wantMyUserID, wantUserID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want, +got)\n%s", diff)
		}
	})

	t.Run("[NG]DMの存在を確認 - 何かエラー発生", func(t *testing.T) {
		wantErr := errors.New("test error")
		rr.GetNoneGroupFunc = func(ctx context.Context, myID, id model.UserID) (*model.Room, error) {
			if diff := cmp.Diff(wantMyUserID, myID); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			return nil, wantErr
		}

		if _, err := ri.Exists(ctx, wantMyUserID, wantUserID); !errors.Is(err, wantErr) {
			t.Errorf("want error is %v, but got error %v", wantErr, err)
		}
	})
}

func Test_Show(t *testing.T) {
	rr := &mock.MockRoomRepository{}
	ri := NewRoomInterractor(rr)
	wantMyUserID := model.UserID(1)
	wantRoomID := model.RoomID(2)

	t.Run("[OK]ルームの詳細を返す", func(t *testing.T) {
		want := &model.RoomUser{
			Room: model.Room{
				ID:        wantRoomID,
				Name:      "test room",
				IsGroup:   false,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       nil,
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
		rr.ShowFunc = func(ctx context.Context, id model.RoomID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return want, nil
		}

		got, err := ri.Show(ctx, wantRoomID, wantMyUserID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want got)\n%s", diff)
		}
	})

	t.Run("[OK]ルームの詳細を返す - グループで自分が存在しない時", func(t *testing.T) {
		want := &model.RoomUser{
			Room: model.Room{
				ID:        wantRoomID,
				Name:      "test room",
				IsGroup:   true,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       pointer.Ptr("test img"),
			},
			Users: []*model.ShowUser{
				{
					ID:        2,
					Name:      "test user 2",
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
		rr.ShowFunc = func(ctx context.Context, id model.RoomID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return want, nil
		}

		got, err := ri.Show(ctx, wantRoomID, wantMyUserID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want got)\n%s", diff)
		}
	})

	t.Run("[NG]ルームの詳細を返す - ルームの閲覧権限がない時", func(t *testing.T) {
		want := &model.RoomUser{
			Room: model.Room{
				ID:        wantRoomID,
				Name:      "test room",
				IsGroup:   false,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       nil,
			},
			Users: []*model.ShowUser{
				{
					ID:        2,
					Name:      "test user 2",
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
		rr.ShowFunc = func(ctx context.Context, id model.RoomID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}

			return want, nil
		}

		if _, err := ri.Show(ctx, wantRoomID, wantMyUserID); !errors.Is(err, myerrors.ErrBadRequestNoPermission) {
			t.Errorf("error is -want %v, +got %v", myerrors.ErrBadRequestNoPermission, err)
		}
	})

	t.Run("[NG]ルームの詳細を返す - 想定外エラー発生", func(t *testing.T) {
		wantErr := errors.New("test err")
		rr.ShowFunc = func(ctx context.Context, id model.RoomID) (*model.RoomUser, error) {
			if diff := cmp.Diff(wantRoomID, id); diff != "" {
				t.Errorf("mismatch (-want got)\n%s", diff)
			}
			return nil, wantErr
		}

		if _, err := ri.Show(ctx, wantRoomID, wantMyUserID); !errors.Is(err, wantErr) {
			t.Errorf("error is -want %v, +got %v", wantErr, err)
		}
	})
}
