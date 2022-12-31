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

func Test_Login(t *testing.T) {
	srv := testApp.srv
	uu := &mock.MockUserUsecase{}
	ac := NewAuthController(srv, uu)
	wantUserID := model.UserID(1)
	wantUserName := "test_user"
	wantEmail := "test@gmail.com"
	wantPassword := "passowrd"
	wantToken := "token"
	hashPass := "hash_password"

	t.Run("[OK]ログイン", func(t *testing.T) {
		user := &model.User{
			ID:        wantUserID,
			Name:      wantUserName,
			Email:     wantEmail,
			Password:  hashPass,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		uu.GetUserByEmailFunc = func(ctx context.Context, email, password string) (*model.User, error) {
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return user, nil
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return &wantToken, nil
		}
		defer func() {
			uu.GetUserByEmailFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		want := &app.Token{
			Token: wantToken,
			User: &app.User{
				CreatedAt: pointer.Ptr(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)),
				Email:     &wantEmail,
				ID:        int(wantUserID),
				Name:      &wantUserName,
				Password:  &hashPass,
				Avatar:    pointer.Ptr("test_avatar"),
			},
		}

		_, got := test.LoginAuthOK(t, ctx, srv, ac, &app.LoginAuthPayload{
			Email:    wantEmail,
			Password: wantPassword,
		})
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ログイン - メールアドレスが間違っているとき", func(t *testing.T) {
		uu.GetUserByEmailFunc = func(ctx context.Context, email, password string) (*model.User, error) {
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, sql.ErrNoRows
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return &wantToken, nil
		}
		defer func() {
			uu.GetUserByEmailFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		test.LoginAuthNotFound(t, ctx, srv, ac, &app.LoginAuthPayload{
			Email:    wantEmail,
			Password: wantPassword,
		})
	})

	t.Run("[NG]ログイン - パスワードが間違っているとき", func(t *testing.T) {
		uu.GetUserByEmailFunc = func(ctx context.Context, email, password string) (*model.User, error) {
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.ErrPasswordWorng
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return &wantToken, nil
		}
		defer func() {
			uu.GetUserByEmailFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		test.LoginAuthBadRequest(t, ctx, srv, ac, &app.LoginAuthPayload{
			Email:    wantEmail,
			Password: wantPassword,
		})
	})

	t.Run("[NG]ログイン - トークン作成に失敗した時", func(t *testing.T) {
		user := &model.User{
			ID:        wantUserID,
			Name:      wantUserName,
			Email:     wantEmail,
			Password:  hashPass,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		uu.GetUserByEmailFunc = func(ctx context.Context, email, password string) (*model.User, error) {
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return user, nil
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, errors.New("test_error")
		}
		defer func() {
			uu.GetUserByEmailFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		test.LoginAuthInternalServerError(t, ctx, srv, ac, &app.LoginAuthPayload{
			Email:    wantEmail,
			Password: wantPassword,
		})
	})
}

func Test_SignUp(t *testing.T) {
	srv := testApp.srv
	uu := &mock.MockUserUsecase{}
	ac := NewAuthController(srv, uu)
	wantUserName := "test_user"
	wantEmail := "test@gmail.com"
	wantPassword := "passowrd"
	wantAvatar := pointer.Ptr("test_avatar")
	wantToken := "token"
	hashPass := "hash_password"

	t.Run("[OK]アカウント登録", func(t *testing.T) {
		user := &model.User{
			ID:        1,
			Name:      wantUserName,
			Email:     wantEmail,
			Password:  hashPass,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    wantAvatar,
		}
		uu.SignUpFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantAvatar, avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return user, nil
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			return &wantToken, nil
		}
		defer func() {
			uu.SignUpFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		want := &app.Token{
			Token: wantToken,
			User: &app.User{
				CreatedAt: pointer.Ptr(time.Date(2022, 1, 1, 0, 0, 0, 0, jst)),
				Email:     &wantEmail,
				ID:        1,
				Name:      &wantUserName,
				Password:  &hashPass,
				Avatar:    wantAvatar,
			},
		}

		_, got := test.SignUpAuthCreated(t, ctx, srv, ac, &app.SignUpAuthPayload{
			Avatar:   wantAvatar,
			Email:    wantEmail,
			Name:     wantUserName,
			Password: wantPassword,
		})

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]アカウント作成 - メアドが使用済み", func(t *testing.T) {
		uu.SignUpFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantAvatar, avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, myerrors.MySQLErrorDuplicate
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			return &wantToken, nil
		}
		defer func() {
			uu.SignUpFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		test.SignUpAuthBadRequest(t, ctx, srv, ac, &app.SignUpAuthPayload{
			Avatar:   wantAvatar,
			Email:    wantEmail,
			Name:     wantUserName,
			Password: wantPassword,
		})
	})

	t.Run("[NG]アカウント作成 - トークン作成エラー", func(t *testing.T) {
		user := &model.User{
			ID:        1,
			Name:      wantUserName,
			Email:     wantEmail,
			Password:  hashPass,
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    wantAvatar,
		}
		uu.SignUpFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantAvatar, avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return user, nil
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			return nil, errors.New("test_error")
		}
		defer func() {
			uu.SignUpFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		test.SignUpAuthInternalServerError(t, ctx, srv, ac, &app.SignUpAuthPayload{
			Avatar:   wantAvatar,
			Email:    wantEmail,
			Name:     wantUserName,
			Password: wantPassword,
		})
	})

	t.Run("[NG]アカウント作成 - ユーザー作成エラー", func(t *testing.T) {
		uu.SignUpFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(wantUserName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantPassword, password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantAvatar, avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, errors.New("test_error")
		}
		uu.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			return nil, errors.New("test_error")
		}
		defer func() {
			uu.SignUpFunc = nil
			uu.CreateJWTTokenFunc = nil
		}()

		test.SignUpAuthInternalServerError(t, ctx, srv, ac, &app.SignUpAuthPayload{
			Avatar:   wantAvatar,
			Email:    wantEmail,
			Name:     wantUserName,
			Password: wantPassword,
		})
	})
}
