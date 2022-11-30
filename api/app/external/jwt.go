package external

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/service"
)

var _ service.JWTService = (*JWTExternal)(nil)

var JWTDatastoreSet = wire.NewSet(
	NewJWTExternal,
	wire.Bind(new(service.JWTService), new(*JWTExternal)),
)

type JWTExternal struct {
	tr repository.TimeRepository
}

func NewJWTExternal(tr repository.TimeRepository) *JWTExternal {
	return &JWTExternal{tr: tr}
}

func (j *JWTExternal) CreateJWTToken(ctx context.Context, id int, name string) (*string, error) {
	claims := jwt.MapClaims{
		"sub":       "auth jwt",
		"user_id":   id,
		"user_name": name,
		"scope":     "api:access",
		"exp":       j.tr.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
