package external

import (
	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/service"
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

func (g *googleExternal) GetLoginURL(state string) string {
	return g.Config.AuthCodeURL(state)
}
