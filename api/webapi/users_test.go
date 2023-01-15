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
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/kod-source/docker-goa-next/webapi/app/test"
	"github.com/shogo82148/pointer"
)

func TestGetCurrentUser(t *testing.T) {
	srv := testApp.srv
	uu := &mock.MockUserUsecase{}
	u := NewUsersController(srv, uu)

	wantUserID := model.UserID(1234)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantUserID))

	t.Run("[OK] ログインしているユーザー取得", func(t *testing.T) {
		user := &model.User{
			ID:        wantUserID,
			Name:      "テスト",
			Email:     "test@gmail.com",
			Password:  "hash_password",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		uu.GetUserFunc = func(ctx context.Context, id model.UserID) (*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("argument mismasch `user_id` (-want +got)\n%s", diff)
			}
			return user, nil
		}
		defer func() {
			uu.GetUserFunc = nil
		}()

		want := &app.User{
			ID:        1234,
			Name:      pointer.Ptr("テスト"),
			Email:     pointer.Ptr("test@gmail.com"),
			Password:  pointer.Ptr("hash_password"),
			CreatedAt: pointer.Ptr(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)),
			Avatar:    pointer.Ptr("test_avatar"),
		}

		_, got := test.GetCurrentUserUsersOK(t, ctx, srv, u)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("response mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG] ユーザーが存在しないとき", func(t *testing.T) {
		uu.GetUserFunc = func(ctx context.Context, id model.UserID) (*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("argument mismasch `user_id` (-want +got)\n%s", diff)
			}
			return nil, sql.ErrNoRows
		}
		defer func() {
			uu.GetUserFunc = nil
		}()

		test.GetCurrentUserUsersNotFound(t, ctx, srv, u)
	})

	t.Run("[NG] 500エラーのとき", func(t *testing.T) {
		uu.GetUserFunc = func(ctx context.Context, id model.UserID) (*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("argument mismasch `user_id` (-want +got)\n%s", diff)
			}
			return nil, errors.New("test error")
		}
		defer func() {
			uu.GetUserFunc = nil
		}()

		test.GetCurrentUserUsersInternalServerError(t, ctx, srv, u)
	})
}

func TestShowUser(t *testing.T) {
	srv := testApp.srv
	uu := &mock.MockUserUsecase{}
	u := NewUsersController(srv, uu)
	wantUserID := model.UserID(1234)

	t.Run("[OK] ユーザー取得", func(t *testing.T) {
		user := &model.User{
			ID:        wantUserID,
			Name:      "テスト",
			Email:     "test@gmail.com",
			Password:  "hash_password",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		uu.GetUserFunc = func(ctx context.Context, id model.UserID) (*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("argument mismasch `user_id` (-want +got)\n%s", diff)
			}
			return user, nil
		}
		defer func() {
			uu.GetUserFunc = nil
		}()

		want := &app.ShowUser{
			ID:        1234,
			Name:      "テスト",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}

		_, got := test.ShowUserUsersOK(t, ctx, srv, u, int(wantUserID))
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("response mismach (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG] ユーザーが存在しないとき", func(t *testing.T) {
		uu.GetUserFunc = func(ctx context.Context, id model.UserID) (*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("argument mismasch `user_id` (-want +got)\n%s", diff)
			}
			return nil, sql.ErrNoRows
		}
		defer func() {
			uu.GetUserFunc = nil
		}()

		test.ShowUserUsersNotFound(t, ctx, srv, u, int(wantUserID))
	})

	t.Run("[NG] 500エラーのとき", func(t *testing.T) {
		uu.GetUserFunc = func(ctx context.Context, id model.UserID) (*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("argument mismasch `user_id` (-want +got)\n%s", diff)
			}
			return nil, errors.New("test error")
		}
		defer func() {
			uu.GetUserFunc = nil
		}()

		test.ShowUserUsersInternalServerError(t, ctx, srv, u, int(wantUserID))
	})
}

func Test_IndexUser(t *testing.T) {
	srv := testApp.srv
	uu := &mock.MockUserUsecase{}
	u := NewUsersController(srv, uu)
	wantUserID := model.UserID(1234)
	ctx = context.WithValue(ctx, userIDCodeKey, int(wantUserID))

	t.Run("[OK]全ユーザー取得", func(t *testing.T) {
		users := []*model.User{
			{
				ID:        1,
				Name:      "test1",
				Email:     "test1@gmail.com",
				CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
				Avatar:    pointer.Ptr("test1_avatar"),
			},
			{
				ID:        2,
				Name:      "test2",
				Email:     "test2@gmail.com",
				CreatedAt: time.Date(2022, 2, 2, 0, 0, 0, 0, jst),
				Avatar:    nil,
			},
		}
		uu.IndexUserFunc = func(ctx context.Context, id model.UserID) ([]*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return users, nil
		}
		defer func() {
			uu.IndexUserFunc = nil
		}()

		want := app.UserCollection{
			{
				ID:        1,
				Name:      pointer.Ptr("test1"),
				Email:     pointer.Ptr("test1@gmail.com"),
				CreatedAt: pointer.Ptr(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)),
				Avatar:    pointer.Ptr("test1_avatar"),
			},
			{
				ID:        2,
				Name:      pointer.Ptr("test2"),
				Email:     pointer.Ptr("test2@gmail.com"),
				CreatedAt: pointer.Ptr(time.Date(2022, 2, 2, 0, 0, 0, 0, jst)),
				Avatar:    nil,
			},
		}
		_, got := test.IndexUsersOK(t, ctx, srv, u)

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]全ユーザー取得 - ユーザーが存在しないケース", func(t *testing.T) {
		uu.IndexUserFunc = func(ctx context.Context, id model.UserID) ([]*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, sql.ErrNoRows
		}
		defer func() {
			uu.IndexUserFunc = nil
		}()

		test.IndexUsersNotFound(t, ctx, srv, u)
	})

	t.Run("[NG]全ユーザー取得 - 想定外エラー発生", func(t *testing.T) {
		uu.IndexUserFunc = func(ctx context.Context, id model.UserID) ([]*model.User, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return nil, errors.New("test error")
		}
		defer func() {
			uu.IndexUserFunc = nil
		}()

		test.IndexUsersInternalServerError(t, ctx, srv, u)
	})
}
