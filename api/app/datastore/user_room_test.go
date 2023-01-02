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

func Test_CreateUserRoom(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return time.Date(2022, 1, 1, 0, 0, 0, 0, jst)
	}
	urr := NewUserRoomRepository(testDB, tr)

	wantRoomID := 39
	if err := schema.InsertRoom(ctx, testDB, &schema.Room{
		ID:        uint64(wantRoomID),
		Name:      "test create room",
		IsGroup:   true,
		CreatedAt: now,
		UpdatedAt: now,
	}); err != nil {
		t.Fatal(err)
	}
	defer func() {
		rr := NewRoomDatastore(testDB, nil)
		if err := rr.Delete(ctx, model.RoomID(wantRoomID)); err != nil {
			t.Fatal(err)
		}
	}()

	t.Run("[OK]UserRoomの作成", func(t *testing.T) {
		got, err := urr.Create(ctx, model.RoomID(wantRoomID), 1)
		if err != nil {
			t.Fatal(err)
		}

		ur, err := schema.SelectUserRoom(ctx, testDB, &schema.UserRoom{ID: uint64(got.ID)})
		if err != nil {
			t.Fatal(err)
		}
		want := &model.UserRoom{
			ID:         model.UserRoomID(ur.ID),
			UserID:     1,
			RoomID:     model.RoomID(wantRoomID),
			LastReadAt: pointer.PtrOrNil(ur.LastReadAt.Time),
			CreatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n %s", diff)
		}
	})

	t.Run("[NG]UserRoomの作成 - 存在しないルームIDを指定", func(t *testing.T) {
		_, err := urr.Create(ctx, 1000, 1)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error %v, but got error is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})

	t.Run("[NG]UserRoomの作成 - 存在しないユーザーIDを指定", func(t *testing.T) {
		_, err := urr.Create(ctx, model.RoomID(wantRoomID), 1000)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorAddOrUpdateForeignKey.Number {
			t.Errorf("want error %v, but got error is %v", myerrors.MySQLErrorAddOrUpdateForeignKey.Number, code)
		}
	})
}
