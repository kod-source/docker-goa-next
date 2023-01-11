package datastore

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
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
	wantImg := pointer.Ptr("test img")

	t.Run("[OK]ルーム作成", func(t *testing.T) {
		got, err := rd.Create(ctx, "test_room", false, []model.UserID{1, 2}, nil)
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
				continue
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
				Img:       nil,
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
		if err := rd.Delete(ctx, model.RoomID(room.ID)); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]ルーム作成 - グループ作成", func(t *testing.T) {
		got, err := rd.Create(ctx, "test_group_room", true, []model.UserID{1, 2}, wantImg)
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
				continue
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
				Name:      "test_group_room",
				IsGroup:   true,
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Img:       wantImg,
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
		if err := rd.Delete(ctx, model.RoomID(room.ID)); err != nil {
			t.Fatal(err)
		}
	})
}

func Test_DeleteRoom(t *testing.T) {
	rd := NewRoomDatastore(testDB, nil)

	t.Run("[OK]ルームの削除", func(t *testing.T) {
		roomID := 5

		if err := schema.InsertRoom(ctx, testDB, &schema.Room{
			ID:        uint64(roomID),
			Name:      "delete_room",
			IsGroup:   true,
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		}); err != nil {
			t.Fatal(err)
		}
		room, err := schema.SelectRoom(ctx, testDB, &schema.Room{ID: uint64(roomID)})
		if err != nil {
			t.Fatal(err)
		}

		userRooms := []*schema.UserRoom{
			{
				ID:     7,
				UserID: 1,
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:     8,
				UserID: 2,
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  now,
					Valid: true,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
		}
		if err := schema.InsertUserRoom(ctx, testDB, userRooms...); err != nil {
			t.Fatal(err)
		}

		if err := rd.Delete(ctx, model.RoomID(room.ID)); err != nil {
			t.Fatal(err)
		}
		if _, err := schema.SelectRoom(ctx, testDB, &schema.Room{ID: room.ID}); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error is %v, but got error is %v", sql.ErrNoRows, err)
		}
		// ON DELETE CASCADEの確認
		for _, ur := range userRooms {
			if _, err := schema.SelectRoom(ctx, testDB, &schema.Room{ID: ur.ID}); !errors.Is(err, sql.ErrNoRows) {
				t.Errorf("want error is %v, but got error is %v", sql.ErrNoRows, err)
			}
		}
	})
}

func Test_IndexRoom(t *testing.T) {
	rd := NewRoomDatastore(testDB, nil)

	t.Run("[OK]ルーム表示", func(t *testing.T) {
		want := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        2,
					Name:      "test2_room",
					IsGroup:   false,
					CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					Img:       nil,
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread5"),
				CountUser: 2,
				ShowImg:   pointer.Ptr("test1_avatar"),
			},
			{
				Room: model.Room{
					ID:        1,
					Name:      "test1_room",
					IsGroup:   true,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       pointer.Ptr("test1_img"),
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread3"),
				CountUser: 2,
				ShowImg:   nil,
			},
		}
		got, gotNextID, err := rd.Index(ctx, 2, 0)
		if err != nil {
			t.Fatal(err)
		}

		if gotNextID != nil {
			t.Errorf("want nextID is nil but got is %d", *gotNextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ルーム表示 - 空のスレッドがある時", func(t *testing.T) {
		roomID := 6
		if err := schema.InsertRoom(ctx, testDB, &schema.Room{
			ID:        uint64(roomID),
			Name:      "test_create_room",
			IsGroup:   true,
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "test_create_img",
				Valid:  true,
			},
		}); err != nil {
			t.Fatal(err)
		}
		userRooms := []*schema.UserRoom{
			{
				ID:     5,
				UserID: 1,
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
					Valid: false,
				},
				CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			},
			{
				ID:     6,
				UserID: 2,
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
					Valid: true,
				},
				CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			},
		}
		if err := schema.InsertUserRoom(ctx, testDB, userRooms...); err != nil {
			t.Fatal(err)
		}

		want := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        2,
					Name:      "test2_room",
					IsGroup:   false,
					CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					Img:       nil,
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread5"),
				CountUser: 2,
				ShowImg:   pointer.Ptr("test1_avatar"),
			},
			{
				Room: model.Room{
					ID:        1,
					Name:      "test1_room",
					IsGroup:   true,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       pointer.Ptr("test1_img"),
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread3"),
				CountUser: 2,
				ShowImg:   nil,
			},
			{
				Room: model.Room{
					ID:        model.RoomID(roomID),
					Name:      "test_create_room",
					IsGroup:   true,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       pointer.Ptr("test_create_img"),
				},
				IsOpen:    true,
				LastText:  nil,
				CountUser: 2,
				ShowImg:   nil,
			},
		}
		got, gotNextID, err := rd.Index(ctx, 2, 0)
		if err != nil {
			t.Fatal(err)
		}

		if gotNextID != nil {
			t.Errorf("want nextID is nil but got is %d", *gotNextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := rd.Delete(ctx, model.RoomID(roomID)); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]ルーム表示 - NextIDを指定", func(t *testing.T) {
		want := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        1,
					Name:      "test1_room",
					IsGroup:   true,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       pointer.Ptr("test1_img"),
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread3"),
				CountUser: 2,
				ShowImg:   nil,
			},
		}
		got, gotNextID, err := rd.Index(ctx, 1, 1)
		if err != nil {
			t.Fatal(err)
		}

		if gotNextID != nil {
			t.Errorf("want nextID is nil but got is %d", *gotNextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ルーム表示 - 既読つきがある時", func(t *testing.T) {
		roomID := 7
		if err := schema.InsertRoom(ctx, testDB, &schema.Room{
			ID:        uint64(roomID),
			Name:      "test_is_opoen_room",
			IsGroup:   false,
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		}); err != nil {
			t.Fatal(err)
		}
		if err := schema.InsertUserRoom(ctx, testDB, &schema.UserRoom{
			ID:     7,
			UserID: 1,
			RoomID: uint64(roomID),
			LastReadAt: sql.NullTime{
				Time:  time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
				Valid: true,
			},
			CreatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
		}); err != nil {
			t.Fatal(err)
		}
		if err := schema.InsertThread(ctx, testDB, &schema.Thread{
			ID:        6,
			UserID:    1,
			RoomID:    uint64(roomID),
			Text:      "test_thread",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "test_img",
				Valid:  true,
			},
		}); err != nil {
			t.Fatal(err)
		}
		want := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        2,
					Name:      "test2_room",
					IsGroup:   false,
					CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					Img:       nil,
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread5"),
				CountUser: 2,
				ShowImg:   nil,
			},
			{
				Room: model.Room{
					ID:        1,
					Name:      "test1_room",
					IsGroup:   true,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       pointer.Ptr("test1_img"),
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread3"),
				CountUser: 2,
				ShowImg:   nil,
			},
			{
				Room: model.Room{
					ID:        model.RoomID(roomID),
					Name:      "test_is_opoen_room",
					IsGroup:   false,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       nil,
				},
				IsOpen:    true,
				LastText:  pointer.Ptr("test_thread"),
				CountUser: 1,
				ShowImg:   nil,
			},
		}
		got, gotNextID, err := rd.Index(ctx, 1, 0)
		if err != nil {
			t.Fatal(err)
		}

		if gotNextID != nil {
			t.Errorf("want nextID is nil but got is %d", *gotNextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := rd.Delete(ctx, model.RoomID(roomID)); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]ルーム表示 - データがたくさんある時", func(t *testing.T) {
		var rooms []*schema.Room
		var userRooms []*schema.UserRoom
		var threads []*schema.Thread
		for i := 0; i < 30; i++ {
			roomID := i + 8
			rooms = append(rooms, &schema.Room{
				ID:        uint64(roomID),
				Name:      fmt.Sprintf("create_room_%d", roomID),
				IsGroup:   false,
				CreatedAt: time.Date(2023, 3, roomID, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2023, 3, roomID, 0, 0, 0, 0, jst),
				Img: sql.NullString{
					String: "",
					Valid:  false,
				},
			})
			userRooms = append(userRooms, &schema.UserRoom{
				ID:     uint64(i),
				UserID: 1,
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  time.Date(2022, 3, 1, 0, 0, 0, 0, jst),
					Valid: false,
				},
				CreatedAt: time.Date(2023, 3, roomID, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2023, 3, roomID, 0, 0, 0, 0, jst),
			})
			threads = append(threads, &schema.Thread{
				ID:        uint64(i),
				UserID:    1,
				RoomID:    uint64(roomID),
				Text:      fmt.Sprintf("thread_%d", roomID),
				CreatedAt: time.Date(2023, 3, roomID, 0, 0, 0, 0, jst),
				UpdatedAt: time.Date(2023, 3, roomID, 0, 0, 0, 0, jst),
				Img: sql.NullString{
					String: "",
					Valid:  false,
				},
			})
		}
		if err := schema.InsertRoom(ctx, testDB, rooms...); err != nil {
			t.Fatal(err)
		}
		if err := schema.InsertUserRoom(ctx, testDB, userRooms...); err != nil {
			t.Fatal(err)
		}
		if err := schema.InsertThread(ctx, testDB, threads...); err != nil {
			t.Fatal(err)
		}

		wantNextID := 25
		_, gotNextID, err := rd.Index(ctx, 1, 5)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(pointer.Ptr(wantNextID), gotNextID); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		for _, r := range rooms {
			if err := rd.Delete(ctx, model.RoomID(r.ID)); err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("[OK]ルーム表示 - CountUserが多い時", func(t *testing.T) {
		roomID := 38
		if err := schema.InsertRoom(ctx, testDB, &schema.Room{
			ID:        uint64(roomID),
			Name:      "count_user_room",
			IsGroup:   true,
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "",
				Valid:  false,
			},
		}); err != nil {
			t.Fatal(err)
		}
		countUserID := 3
		if err := schema.InsertUser(ctx, testDB, &schema.User{
			ID:        uint64(countUserID),
			Name:      "count_user",
			Email:     "count@gmail.com",
			Password:  "count_pass",
			CreatedAt: now,
			UpdatedAt: now,
			Avatar: sql.NullString{
				String: "",
				Valid:  false,
			},
		}); err != nil {
			t.Fatal(err)
		}
		userRooms := []*schema.UserRoom{
			{
				ID:     8,
				UserID: 1,
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  now,
					Valid: false,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:     8,
				UserID: 2,
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  now,
					Valid: false,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:     9,
				UserID: uint64(countUserID),
				RoomID: uint64(roomID),
				LastReadAt: sql.NullTime{
					Time:  now,
					Valid: false,
				},
				CreatedAt: now,
				UpdatedAt: now,
			},
		}
		if err := schema.InsertUserRoom(ctx, testDB, userRooms...); err != nil {
			t.Fatal(err)
		}
		if err := schema.InsertThread(ctx, testDB, &schema.Thread{
			ID:        7,
			UserID:    1,
			RoomID:    uint64(roomID),
			Text:      "test_thread",
			CreatedAt: now,
			UpdatedAt: now,
			Img: sql.NullString{
				String: "test_img",
				Valid:  true,
			},
		}); err != nil {
			t.Fatal(err)
		}
		want := []*model.IndexRoom{
			{
				Room: model.Room{
					ID:        2,
					Name:      "test2_room",
					IsGroup:   false,
					CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
					Img:       nil,
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread5"),
				CountUser: 2,
				ShowImg:   pointer.Ptr("test1_avatar"),
			},
			{
				Room: model.Room{
					ID:        1,
					Name:      "test1_room",
					IsGroup:   true,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       pointer.Ptr("test1_img"),
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("thread3"),
				CountUser: 2,
				ShowImg:   nil,
			},
			{
				Room: model.Room{
					ID:        model.RoomID(roomID),
					Name:      "count_user_room",
					IsGroup:   true,
					CreatedAt: now,
					UpdatedAt: now,
					Img:       nil,
				},
				IsOpen:    false,
				LastText:  pointer.Ptr("test_thread"),
				CountUser: 3,
				ShowImg:   nil,
			},
		}
		got, gotNextID, err := rd.Index(ctx, 2, 0)
		if err != nil {
			t.Fatal(err)
		}

		if gotNextID != nil {
			t.Errorf("want nextID is nil but got is %d", *gotNextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if err := rd.Delete(ctx, model.RoomID(roomID)); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[OK]ルーム表示 - データが存在しない時", func(t *testing.T) {
		var want []*model.IndexRoom
		got, gotNextID, err := rd.Index(ctx, 1, 1000)
		if err != nil {
			t.Fatal(err)
		}
		if gotNextID != nil {
			t.Errorf("want nextID is nil but got is %v", *gotNextID)
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルーム表示 - 存在しないUserIDを指定した時", func(t *testing.T) {
		_, gotNextID, err := rd.Index(ctx, 1000, 0)
		if !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error is %v, but got error is %v", sql.ErrNoRows, err)
		}
		if gotNextID != nil {
			t.Errorf("want nextID is nil but got is %v", *gotNextID)
		}
	})
}

func Test_GetNoneGroup(t *testing.T) {
	rd := NewRoomDatastore(testDB, nil)

	t.Run("[OK]DMのルームを取得", func(t *testing.T) {
		want := &model.Room{
			ID:        2,
			Name:      "test2_room",
			IsGroup:   false,
			CreatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			UpdatedAt: time.Date(2022, 2, 1, 0, 0, 0, 0, jst),
			Img:       nil,
		}
		got, err := rd.GetNoneGroup(ctx, 1, 2)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]DMのルームを取得 - ルームが存在しない時", func(t *testing.T) {
		if _, err := rd.GetNoneGroup(ctx, 1, 1000); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("error is want %v, got %v", sql.ErrNoRows, err)
		}
	})

	t.Run("[NG]DMのルームを取得 - ルームが存在しない時②", func(t *testing.T) {
		if _, err := rd.GetNoneGroup(ctx, 1000, 2); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("error is want %v, got %v", sql.ErrNoRows, err)
		}
	})
}

func Test_ShowRoom(t *testing.T) {
	rd := NewRoomDatastore(testDB, nil)
	wantRoomID := model.RoomID(1)

	t.Run("[OK]ルームの詳細を取得する", func(t *testing.T) {
		want := &model.RoomUser{
			Room: model.Room{
				ID:        wantRoomID,
				Name:      "test1_room",
				IsGroup:   true,
				CreatedAt: now,
				UpdatedAt: now,
				Img:       pointer.Ptr("test1_img"),
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
		got, err := rd.Show(ctx, wantRoomID)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ルームの詳細を取得する - ルームが存在しないケース", func(t *testing.T) {
		if _, err := rd.Show(ctx, 1000); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("error is (-want %v, got %v)", sql.ErrNoRows, err)
		}
	})
}
