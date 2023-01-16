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

func TestGetUser(t *testing.T) {
	ud := NewUserDatastore(testDB, nil)

	t.Run("[OK]ユーザー取得", func(t *testing.T) {
		want := &model.User{
			ID:        1,
			Name:      "test1_name",
			Email:     "test1@gmail.com",
			Password:  "test1_passowrd",
			CreatedAt: now,
			Avatar:    pointer.Ptr("test1_avatar"),
		}
		got, err := ud.GetUser(ctx, want.ID)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ユーザー取得 - Avatarがnilの時", func(t *testing.T) {
		want := &model.User{
			ID:        2,
			Name:      "test2_name",
			Email:     "test2@gmail.com",
			Password:  "test2_passowrd",
			CreatedAt: now,
			Avatar:    nil,
		}
		got, err := ud.GetUser(ctx, want.ID)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ユーザーが存在しないとき", func(t *testing.T) {
		if _, err := ud.GetUser(ctx, 100); !errors.Is(err, sql.ErrNoRows) {
			t.Fatal(err)
		}
	})
}

func TestGetUserByEmail(t *testing.T) {
	ud := NewUserDatastore(testDB, nil)

	t.Run("[OK]ユーザー取得", func(t *testing.T) {
		want := &model.User{
			ID:        1,
			Name:      "test1_name",
			Email:     "test1@gmail.com",
			Password:  "test1_passowrd",
			CreatedAt: now,
			Avatar:    pointer.Ptr("test1_avatar"),
		}
		got, err := ud.GetUserByEmail(ctx, want.Email)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ユーザー取得 - Avatarがnilの時", func(t *testing.T) {
		want := &model.User{
			ID:        2,
			Name:      "test2_name",
			Email:     "test2@gmail.com",
			Password:  "test2_passowrd",
			CreatedAt: now,
			Avatar:    nil,
		}
		got, err := ud.GetUserByEmail(ctx, want.Email)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ユーザーが存在しないとき", func(t *testing.T) {
		if _, err := ud.GetUserByEmail(ctx, "unknow@gmail.com"); !errors.Is(err, sql.ErrNoRows) {
			t.Fatal(err)
		}
	})
}

func Test_CreateUser(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return time.Date(2022, 1, 1, 0, 0, 0, 0, jst)
	}
	defer func() {
		tr.NowFunc = nil
	}()
	ud := NewUserDatastore(testDB, tr)

	t.Run("[OK]ユーザー作成", func(t *testing.T) {
		got, err := ud.CreateUser(ctx, "create_name", "create@gmail.com", "create_password", pointer.Ptr("create_avatar"))
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{
			ID: uint64(got.ID),
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &model.User{
			ID:        model.UserID(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			Avatar:    pointer.PtrOrNil(user.Avatar.String),
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]ユーザー作成 - avatarがnilの時", func(t *testing.T) {
		got, err := ud.CreateUser(ctx, "create_no_avatar_name", "create_no_avatar@gmail.com", "craate_no_avatar_password", nil)
		if err != nil {
			t.Fatal(err)
		}
		user, err := schema.SelectUser(ctx, testDB, &schema.User{
			ID: uint64(got.ID),
		})
		if err != nil {
			t.Fatal(err)
		}
		want := &model.User{
			ID:        model.UserID(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
			Avatar:    pointer.PtrOrNil(user.Avatar.String),
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ユーザー作成 - Emailがユニークインデックスの確認", func(t *testing.T) {
		// 存在しているメールアドレスを指定する
		_, err := ud.CreateUser(ctx, "", "test1@gmail.com", "", nil)
		if code := myerrors.GetMySQLErrorNumber(err); code != myerrors.MySQLErrorDuplicate.Number {
			t.Fatal(err)
		}
	})
}

func Test_IndexUser(t *testing.T) {
	ur := NewUserDatastore(testDB, nil)

	wantUserID := 7
	if err := schema.InsertUser(ctx, testDB, &schema.User{
		ID:        uint64(wantUserID),
		Name:      "create_user",
		Email:     "create@exmaple.com",
		Password:  "create_pas",
		CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, jst),
		UpdatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, jst),
		Avatar: sql.NullString{
			String: "create_avatar",
			Valid:  true,
		},
	}); err != nil {
		t.Fatal(err)
	}

	t.Run("[OK]自分以外の全てのユーザーを取得", func(t *testing.T) {
		want := []*model.User{
			{
				ID:        model.UserID(wantUserID),
				Name:      "create_user",
				Email:     "create@exmaple.com",
				CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("create_avatar"),
			},
			{
				ID:        2,
				Name:      "test2_name",
				Email:     "test2@gmail.com",
				CreatedAt: now,
				Avatar:    nil,
			},
			{
				ID:        3,
				Name:      "count_user",
				Email:     "count@gmail.com",
				CreatedAt: now,
				Avatar:    nil,
			},
			{
				ID:        4,
				Name:      "create_name",
				Email:     "create@gmail.com",
				CreatedAt: now,
				Avatar:    pointer.Ptr("create_avatar"),
			},
			{
				ID:        5,
				Name:      "create_no_avatar_name",
				Email:     "create_no_avatar@gmail.com",
				CreatedAt: now,
				Avatar:    nil,
			},
		}
		got, err := ur.IndexUser(ctx, 1)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})
}
