package webapi

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	interactor "github.com/kod-source/docker-goa-next/app/interactor/mock"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/webapi/app"
	"github.com/kod-source/docker-goa-next/webapi/app/test"
	goa "github.com/shogo82148/goa-v1"
	"github.com/shogo82148/pointer"
)

func Test_Login(t *testing.T) {
	srv := testApp.srv
	uu := &interactor.MockUserUsecase{}
	ac := NewAuthController(srv, uu, nil)
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
	uu := &interactor.MockUserUsecase{}
	ac := NewAuthController(srv, uu, nil)
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

func Test_newAuthMiddleware(t *testing.T) {
	testHelper := func(middleware goa.Middleware, req *http.Request, h goa.Handler) (*httptest.ResponseRecorder, error) {
		t.Helper()
		service := goa.New("kd")
		service.Encoder = goa.NewHTTPEncoder()
		service.Encoder.Register(goa.NewJSONEncoder, "*/*")

		rw := httptest.NewRecorder()
		if err := middleware(h)(ctx, rw, req); err != nil {
			return nil, err
		}
		rw.Flush()
		return rw, nil
	}

	t.Run("[NG]認証ヘッダーなしの場合HTTPステータスが400を返すこと", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
		if err != nil {
			t.Fatal(err)
		}
		handler := func(ctx context.Context, _ http.ResponseWriter, _ *http.Request) error {
			t.Fatal("must not reach here")
			panic("認証エラーなので、エラーハンドラーは呼ばれない")
		}

		_, err = testHelper(newAuthMiddleware(), req, handler)
		var goaErr goa.ServiceError
		if !errors.As(err, &goaErr) {
			t.Errorf("want goa.ServiceError, got %t", err)
		}
		if goaErr.ResponseStatus() != http.StatusBadRequest {
			t.Errorf("unexpected status code: want %d, got %d", http.StatusBadRequest, goaErr.ResponseStatus())
		}
	})

	t.Run("[NG]不正なトークンの場合HTTPステータスが400を返すこと", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
		if err != nil {
			t.Fatal(err)
		}
		handler := func(ctx context.Context, _ http.ResponseWriter, _ *http.Request) error {
			t.Fatal("must not reach here")
			panic("認証エラーなので、エラーハンドラーは呼ばれない")
		}

		// JWTではないトークン
		req.Header.Set("Authorization", "Bearer VERY-VERY-SECRET")

		_, err = testHelper(newAuthMiddleware(), req, handler)
		var goaErr goa.ServiceError
		if !errors.As(err, &goaErr) {
			t.Errorf("want goa.ServiceError, got %t", err)
		}
		if goaErr.ResponseStatus() != http.StatusBadRequest {
			t.Errorf("unexpected status code: want %d, got %d", http.StatusBadRequest, goaErr.ResponseStatus())
		}
	})

	// トークンの有効期限が切れるとテストに落ちるためコメント
	// t.Run("[OK]正常なトークンの場合エラーが発生しないこと", func(t *testing.T) {
	// 	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	wantUserID := 1
	// 	handler := func(ctx context.Context, _ http.ResponseWriter, _ *http.Request) error {
	// 		if diff := cmp.Diff(wantUserID, getUserIDCode(ctx)); diff != "" {
	// 			t.Errorf("mismatch (-want +got)\n%s", diff)
	// 		}
	// 		return nil
	// 	}

	// 	// 正常なトークン
	// 	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzYzNzk1ODMsInNjb3BlIjoiYXBpOmFjY2VzcyIsInN1YiI6ImF1dGggand0IiwidXNlcl9pZCI6MSwidXNlcl9uYW1lIjoi44Gp44KTIn0.Eu6fa77kpzZ-M19dUYY08efzxkxBDVv6Z6R9hJsyL9c")

	// 	if _, err = testHelper(newAuthMiddleware(), req, handler); err != nil {
	// 		t.Fatal(err)
	// 	}
	// })
}

func Test_GoogleLogin(t *testing.T) {
	srv := testApp.srv
	uu := &interactor.MockGoogleUsecase{}
	ac := NewAuthController(srv, nil, uu)
	wantRedirectURL := "https://accounts.google.com/o/oauth2/auth?client_id=mock_client_id&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fcallback%2Fgoogle&response_type=code&scope=openid&state=test-state"

	t.Run("[OK]GoogleアカウントログインリダイレクトURL取得", func(t *testing.T) {
		uu.GetLoginURLFunc = func(state string) string {
			return wantRedirectURL
		}
		defer func() {
			uu.GetLoginURLFunc = nil
		}()

		want := &app.RedirectURI{
			URL: wantRedirectURL,
		}
		_, got := test.GoogleLoginAuthOK(t, ctx, srv, ac)
		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("mismatch (-want, +got)\n%s", diff)
		}
	})
}
