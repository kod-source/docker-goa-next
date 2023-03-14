package interactor

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	repository "github.com/kod-source/docker-goa-next/app/datastore/mock"
	external "github.com/kod-source/docker-goa-next/app/external/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/shogo82148/pointer"
)

func Test_GetLoginURL(t *testing.T) {
	gs := &external.MockGoogleService{}
	gu := NewGoogleUseCase(gs, nil)
	testState := "test-state"
	wantURL := "https://accounts.google.com/o/oauth2/auth?client_id=mock_client_id&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fcallback%2Fgoogle&response_type=code&scope=openid&state=test-state"

	t.Run("[OK]GoogleアカウントログインリダイレクトURL取得", func(t *testing.T) {
		gs.GetLoginURLFunc = func(state string) string {
			if diff := cmp.Diff(testState, state); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			return wantURL
		}
		got := gu.GetLoginURL(testState)

		if diff := cmp.Diff(wantURL, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})
}

func Test_GetOrCreateUserInfo(t *testing.T) {
	gs := &external.MockGoogleService{}
	ur := &repository.MockUserRepository{}
	gu := NewGoogleUseCase(gs, ur)
	wantCode := "want_code"
	mockGoogleUser := &model.GoogleUser{
		Name:    "google user",
		Picture: "https://google_user.com",
		Email:   "google_user@gmail.com",
	}
	wantUser := &model.User{
		ID:        1,
		Name:      "google user",
		Email:     "google_user@gmail.com",
		Password:  "",
		CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, jst),
		Avatar:    pointer.Ptr("https://google_user.com"),
	}
	wantErr := errors.New("test error")

	t.Run("[OK]Googleアカウント情報からユーザー作成", func(t *testing.T) {
		gs.GetUserInfoFunc = func(ctx context.Context, code string) (*model.GoogleUser, error) {
			if diff := cmp.Diff(wantCode, code); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return mockGoogleUser, nil
		}
		ur.GetUserByEmailFunc = func(ctx context.Context, email string) (*model.User, error) {
			if diff := cmp.Diff(mockGoogleUser.Email, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, sql.ErrNoRows
		}
		ur.CreateUserFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(mockGoogleUser.Name, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(mockGoogleUser.Email, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(mockGoogleUser.Picture, *avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff("", password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return wantUser, nil
		}

		defer func() {
			gs.GetUserInfoFunc = nil
			ur.GetUserByEmailFunc = nil
			ur.CreateUserFunc = nil
		}()

		got, err := gu.GetOrCreateUserInfo(ctx, wantCode)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(wantUser, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[OK]Googleアカウント情報からユーザー取得", func(t *testing.T) {
		gs.GetUserInfoFunc = func(ctx context.Context, code string) (*model.GoogleUser, error) {
			if diff := cmp.Diff(wantCode, code); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return mockGoogleUser, nil
		}
		ur.GetUserByEmailFunc = func(ctx context.Context, email string) (*model.User, error) {
			if diff := cmp.Diff(mockGoogleUser.Email, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return wantUser, nil
		}

		defer func() {
			gs.GetUserInfoFunc = nil
			ur.GetUserByEmailFunc = nil
		}()

		got, err := gu.GetOrCreateUserInfo(ctx, wantCode)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(wantUser, got); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})

	t.Run("[NG]Googleアカウント情報からユーザー取得 - 不正なcodeの時", func(t *testing.T) {
		batCode := "bat code"
		gs.GetUserInfoFunc = func(ctx context.Context, code string) (*model.GoogleUser, error) {
			if diff := cmp.Diff(batCode, code); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, wantErr
		}

		defer func() {
			gs.GetUserInfoFunc = nil
		}()

		if _, err := gu.GetOrCreateUserInfo(ctx, batCode); !errors.Is(wantErr, err) {
			t.Errorf("mismatch error (-want %v, +got %v)", wantErr, err)
		}
	})
	t.Run("[NG]Googleアカウント情報からユーザー取得 - GetUserがエラー時", func(t *testing.T) {
		gs.GetUserInfoFunc = func(ctx context.Context, code string) (*model.GoogleUser, error) {
			if diff := cmp.Diff(wantCode, code); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return mockGoogleUser, nil
		}
		ur.GetUserByEmailFunc = func(ctx context.Context, email string) (*model.User, error) {
			if diff := cmp.Diff(mockGoogleUser.Email, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, wantErr
		}

		defer func() {
			gs.GetUserInfoFunc = nil
			ur.GetUserByEmailFunc = nil
		}()

		if _, err := gu.GetOrCreateUserInfo(ctx, wantCode); !errors.Is(wantErr, err) {
			t.Errorf("mismatch error (-want %v, +got %v)", wantErr, err)
		}
	})
	t.Run("[NG]Googleアカウント情報からユーザー取得 - CreateUserがエラー時", func(t *testing.T) {
		gs.GetUserInfoFunc = func(ctx context.Context, code string) (*model.GoogleUser, error) {
			if diff := cmp.Diff(wantCode, code); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return mockGoogleUser, nil
		}
		ur.GetUserByEmailFunc = func(ctx context.Context, email string) (*model.User, error) {
			if diff := cmp.Diff(mockGoogleUser.Email, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, sql.ErrNoRows
		}
		ur.CreateUserFunc = func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
			if diff := cmp.Diff(mockGoogleUser.Name, name); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(mockGoogleUser.Email, email); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff(mockGoogleUser.Picture, *avatar); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}
			if diff := cmp.Diff("", password); diff != "" {
				t.Errorf("mismatch (-want +got)\n%s", diff)
			}

			return nil, wantErr
		}

		defer func() {
			gs.GetUserInfoFunc = nil
			ur.GetUserByEmailFunc = nil
			ur.CreateUserFunc = nil
		}()

		if _, err := gu.GetOrCreateUserInfo(ctx, wantCode); !errors.Is(wantErr, err) {
			t.Errorf("mismatch error (-want %v, +got %v)", wantErr, err)
		}
	})
}
