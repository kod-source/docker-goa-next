package external

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
)

func Test_GetLoginURL(t *testing.T) {
	oauthStateString := "test-state"
	config := &oauth2.Config{
		ClientID:     "mock_client_id",
		ClientSecret: "mock_client_secret",
		Endpoint:     googleOAuth.Endpoint,
		Scopes:       []string{"openid"},
		RedirectURL:  "http://localhost:8080/auth/callback/google",
	}
	gs := NewGoogleService(config)

	t.Run("[OK]リダイレクトURL生成", func(t *testing.T) {
		wantURL := "https://accounts.google.com/o/oauth2/auth?client_id=mock_client_id&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fcallback%2Fgoogle&response_type=code&scope=openid&state=test-state"
		gotURL := gs.GetLoginURL(oauthStateString)
		// fmt.Println("実行")
		// fmt.Println(gotURL)
		// panic("エラーです")

		if diff := cmp.Diff(wantURL, gotURL); diff != "" {
			t.Errorf("mismatch (-want, +got)\n%s", diff)
		}
	})
}

// func Test_GetUserInfo(t *testing.T) {
// 	config := &oauth2.Config{
// 		ClientID:     "mock_client_id",
// 		ClientSecret: "mock_client_secret",
// 		Endpoint: googleOAuth.Endpoint,
// 		Scopes:   []string{"openid", "email", "profile"},
// 		// Scopes:      []string{"openid"},
// 		RedirectURL: "http://localhost:8080/auth/callback/google",
// 	}
// 	gs := NewGoogleService(config)
// 	wantCode := "4%2F0AWtgzh6GTNr-woVapzAGHlkG_NnEusbutSonN-pP_i2VG_xVRkYFuxh5a6E-vESkj5vWSA"

// 	t.Run("[OK]Googleアカウントでログインするユーザー情報の取得", func(t *testing.T) {
// 		_, err := gs.GetUserInfo(ctx, wantCode)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 	})
// }
