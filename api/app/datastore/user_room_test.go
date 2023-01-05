package datastore

import (
	"database/sql"
	"errors"
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
	wantDMRoomID := 40
	rs := []*schema.Room{
		{
			ID:        uint64(wantRoomID),
			Name:      "test create room",
			IsGroup:   true,
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        uint64(wantDMRoomID),
			Name:      "test create dm room",
			IsGroup:   false,
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	if err := schema.InsertRoom(ctx, testDB, rs...); err != nil {
		t.Fatal(err)
	}
	defer func() {
		rr := NewRoomDatastore(testDB, nil)
		if err := rr.Delete(ctx, model.RoomID(wantRoomID)); err != nil {
			t.Fatal(err)
		}
		if err := rr.Delete(ctx, model.RoomID(wantDMRoomID)); err != nil {
			t.Fatal(err)
		}
	}()

	t.Run("[OK]UserRoomの作成", func(t *testing.T) {
		got, err := urr.Create(ctx, model.RoomID(wantRoomID), 1)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := urr.Delete(ctx, got.ID); err != nil {
				t.Fatal(err)
			}
		}()

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

	t.Run("[OK]UserRoomの作成 - DMに対して作成したケース", func(t *testing.T) {
		got, err := urr.Create(ctx, model.RoomID(wantDMRoomID), 1)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := urr.Delete(ctx, got.ID); err != nil {
				t.Fatal(err)
			}
		}()

		ur, err := schema.SelectUserRoom(ctx, testDB, &schema.UserRoom{ID: uint64(got.ID)})
		if err != nil {
			t.Fatal(err)
		}
		want := &model.UserRoom{
			ID:         model.UserRoomID(ur.ID),
			UserID:     1,
			RoomID:     model.RoomID(wantDMRoomID),
			LastReadAt: pointer.PtrOrNil(ur.LastReadAt.Time),
			CreatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			UpdatedAt:  time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n %s", diff)
		}
	})

	t.Run("[NG]UserRoomの作成 - DMに対して作成したケース人数がおかしいとき", func(t *testing.T) {
		wantUserRoomIDs := []int{49, 50}
		urs := []*schema.UserRoom{
			{
				ID:     uint64(wantUserRoomIDs[0]),
				UserID: 1,
				RoomID: uint64(wantDMRoomID),
				LastReadAt: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:     uint64(wantUserRoomIDs[1]),
				UserID: 2,
				RoomID: uint64(wantDMRoomID),
				LastReadAt: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
		}
		if err := schema.InsertUserRoom(ctx, testDB, urs...); err != nil {
			t.Fatal(err)
		}
		defer func() {
			if err := urr.Delete(ctx, model.UserRoomID(wantUserRoomIDs[0])); err != nil {
				t.Fatal(err)
			}
			if err := urr.Delete(ctx, model.UserRoomID(wantUserRoomIDs[1])); err != nil {
				t.Fatal(err)
			}
		}()

		if _, err := urr.Create(ctx, model.RoomID(wantDMRoomID), 3); !errors.Is(err, myerrors.ErrBadRequestNoPermission) {
			t.Errorf("error is (-want %v, got %v)", myerrors.ErrBadRequestNoPermission, err)
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

func Test_DeleteUserRoom(t *testing.T) {
	urr := NewUserRoomRepository(testDB, nil)

	wantRoomID := 41
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

	t.Run("[OK]UserRoomの削除", func(t *testing.T) {
		userRoomID := 54
		if err := schema.InsertUserRoom(ctx, testDB, &schema.UserRoom{
			ID:     uint64(userRoomID),
			UserID: 1,
			RoomID: uint64(wantRoomID),
			LastReadAt: sql.NullTime{
				Time:  now,
				Valid: false,
			},
			CreatedAt: now,
			UpdatedAt: now,
		}); err != nil {
			t.Fatal(err)
		}

		if err := urr.Delete(ctx, model.UserRoomID(userRoomID)); err != nil {
			t.Fatal(err)
		}
		if _, err := schema.SelectUserRoom(ctx, testDB, &schema.UserRoom{ID: uint64(userRoomID)}); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("error is want %v, got %v", sql.ErrNoRows, err)
		}
	})
}
