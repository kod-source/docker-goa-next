package external

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
)

func Test_GetLoginURL(t *testing.T) {
	oauthStateString := "pseudo-random"
	config := &oauth2.Config{
		ClientID:     "mock_client_id",
		ClientSecret: "mock_client_secret",
		Endpoint:     googleOAuth.Endpoint,
		Scopes:       []string{"openid"},
		RedirectURL:  "http://localhost:8080/auth/callback/google",
	}
	gs := NewGoogleService(config)

	t.Run("[OK]リダイレクトURL生成", func(t *testing.T) {
		wantURL := "https://accounts.google.com/o/oauth2/auth?client_id=mock_client_id&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2Fauth%2Fcallback%2Fgoogle&response_type=code&scope=openid&state=pseudo-random"
		gotURL := gs.GetLoginURL(oauthStateString)

		if diff := cmp.Diff(wantURL, gotURL); diff != "" {
			t.Errorf("mismatch (-want, +got)\n%s", diff)
		}
	})
}
