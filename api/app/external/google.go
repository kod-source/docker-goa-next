package external

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/service"
	"github.com/shogo82148/pointer"
	"golang.org/x/oauth2"
)

var _ service.GoogleService = (*googleExternal)(nil)

var GoogleServiceSet = wire.NewSet(
	NewGoogleService,
	wire.Bind(new(service.GoogleService), new(*googleExternal)),
)

type googleExternal struct {
	Config *oauth2.Config
}

func NewGoogleService(c *oauth2.Config) *googleExternal {
	gex := &googleExternal{
		Config: c,
	}
	return gex
}

// GetLoginURL リダイレクトURLを取得する
func (g *googleExternal) GetLoginURL(state string) string {
	return g.Config.AuthCodeURL(state)
}

// GetUserInfo codeからユーザー情報を返す
func (g *googleExternal) GetUserInfo(ctx context.Context, code string) (*model.User, error) {
	token, err := g.Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	gu, err := g.getUserInfo(ctx, token)
	if err != nil {
		return nil, err
	}

	return &model.User{
		Name:      gu.Name,
		Email:     gu.Email,
		CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, nil),
		Avatar:    pointer.PtrOrNil(gu.Picture),
	}, nil
}

func (g *googleExternal) getUserInfo(ctx context.Context, token *oauth2.Token) (*model.GoogleUser, error) {
	client := g.Config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gu model.GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&gu); err != nil {
		return nil, err
	}

	return &gu, nil
}
