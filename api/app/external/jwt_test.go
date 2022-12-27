package external

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-cmp/cmp"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/shogo82148/pointer"
)

func Test_CreateJWTToken(t *testing.T) {
	tr := &repository.MockTimeRepository{}
	tr.NowFunc = func() time.Time {
		return time.Date(2100, 1, 1, 0, 0, 0, 0, jst)
	}
	je := NewJWTExternal(tr)
	wantUserID := model.UserID(1)
	wantName := "test_user"
	wantScop := "api:access"
	wantSub := "auth jwt"

	t.Run("[OK]JWTトークン作成", func(t *testing.T) {
		gotToken, err := je.CreateJWTToken(ctx, wantUserID, wantName)
		if err != nil {
			t.Fatal(err)
		}

		token, err := jwt.Parse(pointer.Value(gotToken), func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			t.Fatal(err)
		}

		claims := token.Claims.(jwt.MapClaims)
		userID, ok := claims["user_id"].(float64)
		if !ok {
			t.Errorf("claims user_is does not exit")
		}
		userName, ok := claims["user_name"].(string)
		if !ok {
			t.Errorf("claims user_name does not exit")
		}
		scop, ok := claims["scope"].(string)
		if !ok {
			t.Errorf("claims scope does not exit")
		}
		sub, ok := claims["sub"].(string)
		if !ok {
			t.Errorf("claims sub does not exit")
		}

		if diff := cmp.Diff(wantUserID, model.UserID(userID)); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if diff := cmp.Diff(wantName, userName); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if diff := cmp.Diff(wantScop, scop); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
		if diff := cmp.Diff(wantSub, sub); diff != "" {
			t.Errorf("mismatch (-want +got)\n%s", diff)
		}
	})
}
