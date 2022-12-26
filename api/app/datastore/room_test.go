package datastore

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/pointer"
)

func Test_CreateRoom(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return time.Date(2022, 1, 1, 0, 0, 0, 0, jst)
	}
	rd := NewRoomDatastore(testDB, tr)

	t.Run("[OK]ルーム作成", func(t *testing.T) {
		got, err := rd.Create(ctx, "test_room", false, []model.UserID{1, 2})
		if err != nil {
			t.Fatal(err)
		}

		room, err := schema.SelectRoom(ctx, testDB, &schema.Room{ID: uint64(got.Room.ID)})
		if err != nil {
			t.Fatal(err)
		}
		gotUserRooms, err := schema.SelectAllUserRoom(ctx, testDB)
		if err != nil {
			t.Fatal(err)
		}
		for _, ur := range gotUserRooms {
			if ur.RoomID != room.ID {
				t.Errorf("want user_room room_id is %v, but got user_room room_id is %v", room.ID, ur.RoomID)
			}
			if !(ur.UserID == 1 || ur.UserID == 2) {
				t.Errorf("want user_room user_id id is %v or %v, but got user_room user_id is %v", 1, 2, ur.UserID)
			}
			if ur.LastReadAt.Valid {
				t.Errorf("want user_room last_read_at is null, but got user_room last_read_at is %v", ur.LastReadAt.Time)
			}
			if !ur.CreatedAt.Equal(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)) {
				t.Errorf("want user_room created_at is %v, but got user_room created_at is %v", time.Date(2022, 1, 1, 0, 0, 0, 0, jst), ur.CreatedAt)
			}
			if !ur.UpdatedAt.Equal(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)) {
				t.Errorf("want user_room updated_at is %v, but got user_room updated_at is %v", time.Date(2022, 1, 1, 0, 0, 0, 0, jst), ur.UpdatedAt)
			}
		}

		want := &model.RoomUser{
			Room: model.Room{
				ID:        model.RoomID(room.ID),
				Name:      "test_room",
				IsGroup:   false,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			},
			Users: []*model.ShowUser{
				{
					ID:        1,
					Name:      "test1_name",
					CreatedAt: now,
					Avatar:    pointer.Ptr("test1_avatar"),
				},
				{
					ID:        2,
					Name:      "test2_name",
					CreatedAt: now,
					Avatar:    nil,
				},
			},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got\n%s", diff)
		}
	})

	t.Run("[NG]ルーム作成 - 存在しないUserIDを指定した時", func(t *testing.T) {
		_, err := rd.Create(ctx, "test_room", true, []model.UserID{1, 1000})
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error code is %v, but got error code is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})
}
