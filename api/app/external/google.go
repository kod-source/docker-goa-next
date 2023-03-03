package external

import (
	"context"
	"fmt"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/service"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
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
	fmt.Println("実行")
	fmt.Println(code)
	token, err := g.Config.Exchange(ctx, code)
	if err != nil {
		fmt.Println("エラー1")
		fmt.Println(err)
		return nil, err
	}

	client := g.Config.Client(ctx, token)
	service, err := v2.New(client)
	// service, err := v2.NewService(ctx, option.WithClientCertSource(client))
	if err != nil {
		fmt.Println("エラー2")
		fmt.Println(err)
		return nil, err
	}
	userInfo, err := service.Tokeninfo().AccessToken(token.AccessToken).Context(ctx).Do()
	if err != nil {
		fmt.Println("エラー3")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(userInfo)

	return nil, nil
}
