package interactor

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	external "github.com/kod-source/docker-goa-next/app/external/mock"
)

func Test_GetLoginURL(t *testing.T) {
	gs := &external.MockGoogleService{}
	gu := NewGoogleUseCase(gs)
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
