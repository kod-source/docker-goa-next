package interactor

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/datastore/mock"
	serviceMock "github.com/kod-source/docker-goa-next/app/external/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	mock_repository "github.com/kod-source/docker-goa-next/app/repository/mock"
	"github.com/shogo82148/pointer"
)

func Test_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := mock_repository.NewMockUserRepository(ctrl)
	ui := NewUserInteractor(ur, nil)
	wantUserID := model.UserID(1)

	t.Run("[OK]User取得", func(t *testing.T) {
		wantUser := &model.User{
			ID:        wantUserID,
			Name:      "test_user",
			Email:     "test@gmail.com",
			Password:  "$2a$10$i/vSyq8CTN3BBn2bkE.M3eHbSI1JyLp68NW3W6wTsHdSZJi2zWQkG",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		ur.EXPECT().GetUser(ctx, wantUserID).Return(wantUser, nil)

		got, err := ui.GetUser(ctx, wantUserID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(wantUser, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]User取得 - エラー発生", func(t *testing.T) {
		ur.EXPECT().GetUser(ctx, wantUserID).Return(nil, sql.ErrNoRows)

		if _, err := ui.GetUser(ctx, wantUserID); !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("want error is %v, but got error is %v", sql.ErrNoRows, err)
		}
	})
}

func Test_GetUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ur := mock_repository.NewMockUserRepository(ctrl)
	ui := NewUserInteractor(ur, nil)
	wantEmail := "test@gmail.com"

	t.Run("[OK]ログイン", func(t *testing.T) {
		wantUser := &model.User{
			ID:        1,
			Name:      "test_user",
			Email:     wantEmail,
			Password:  "$2a$10$i/vSyq8CTN3BBn2bkE.M3eHbSI1JyLp68NW3W6wTsHdSZJi2zWQkG",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		ur.EXPECT().GetUserByEmail(ctx, wantEmail).Return(wantUser, nil)

		got, err := ui.GetUserByEmail(ctx, wantEmail, "password")
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(wantUser, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]ログイン - パスワードが違うとき", func(t *testing.T) {
		wantUser := &model.User{
			ID:        1,
			Name:      "test_user",
			Email:     wantEmail,
			Password:  "$2a$10$i/vSyq8CTN3BBn2bkE.M3eHbSI1JyLp68NW3W6wTsHdSZJi2zWQkG",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		ur.EXPECT().GetUserByEmail(ctx, wantEmail).Return(wantUser, nil)

		if _, err := ui.GetUserByEmail(ctx, wantEmail, "mistake_password"); !errors.Is(err, myerrors.ErrPasswordWorng) {
			t.Errorf("want error is %v, but got error is %v", myerrors.ErrPasswordWorng, err)
		}
	})

	t.Run("[NG]ログイン - datastoreでエラー発生", func(t *testing.T) {
		wantErr := errors.New("test_error")
		ur.EXPECT().GetUserByEmail(ctx, wantEmail).Return(nil, wantErr)

		if _, err := ui.GetUserByEmail(ctx, wantEmail, "password"); !errors.Is(err, wantErr) {
			t.Errorf("want error is %v, but got error is %v", wantErr, err)
		}
	})
}

func Test_CreateJWTToken(t *testing.T) {
	js := &serviceMock.MockJWTService{}
	ui := NewUserInteractor(nil, js)
	wantUserID := model.UserID(1)
	wantName := "test_user"
	wantToken := "create_token"
	wantError := errors.New("test error")

	t.Run("[OK]JWTトークン作成", func(t *testing.T) {
		js.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return &wantToken, nil
		}

		got, err := ui.CreateJWTToken(ctx, wantUserID, wantName)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(wantToken, *got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]JWTトークン作成 - error発生", func(t *testing.T) {
		js.CreateJWTTokenFunc = func(ctx context.Context, id model.UserID, name string) (*string, error) {
			if diff := cmp.Diff(wantUserID, id); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, wantError
		}

		if _, err := ui.CreateJWTToken(ctx, wantUserID, wantName); !errors.Is(err, wantError) {
			t.Errorf("want error is %v, but got error is %v", wantError, err)
		}
	})
}

func Test_SignUp(t *testing.T) {
	ur := &mock.MockUserRepository{}
	ui := NewUserInteractor(ur, nil)
	wantName := "test_user"
	wantEmail := "test@gmail.com"
	wantPassword := "password"
	wantAvatar := pointer.Ptr("test_avatar")

	t.Run("[OK]アカウント登録", func(t *testing.T) {
		wantUser := &model.User{
			ID:        1,
			Name:      "test_user",
			Email:     "test@gmail.com",
			Password:  "password",
			CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
			Avatar:    pointer.Ptr("test_avatar"),
		}
		ur.CreateUserFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(wantName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantAvatar, avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return wantUser, nil
		}

		got, err := ui.SignUp(ctx, wantName, wantEmail, wantPassword, wantAvatar)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(wantUser, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]アカウント登録 - パスワードがHash化されているか確認", func(t *testing.T) {
		wantHashPassword := "$2a$10"
		ur.CreateUserFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(wantName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantAvatar, avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			// $????$????$?/??????????.??????????????
			ps := strings.Split(password, "$")
			if len(ps) == 3 {
				t.Errorf("want length is 3, but got is %d", len(ps))
			}
			hp := "$" + ps[1] + "$" + ps[2]
			if diff := cmp.Diff(wantHashPassword, hp); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, nil
		}

		if _, err := ui.SignUp(ctx, wantName, wantEmail, wantPassword, wantAvatar); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("[NG]アカウント登録 - エラーが返ってきた時", func(t *testing.T) {
		wantErr := errors.New("test error")
		ur.CreateUserFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(wantName, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantEmail, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(wantAvatar, avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, wantErr
		}

		if _, err := ui.SignUp(ctx, wantName, wantEmail, wantPassword, wantAvatar); !errors.Is(err, wantErr) {
			t.Errorf("want error is %v, but got error is %v", wantErr, err)
		}
	})
}
