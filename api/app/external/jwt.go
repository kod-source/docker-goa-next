package external

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/service"
)

var _ service.JWTService = (*jwtExternal)(nil)

var JWTDatastoreSet = wire.NewSet(
	NewJWTExternal,
	wire.Bind(new(service.JWTService), new(*jwtExternal)),
)

type jwtExternal struct {
	tr repository.TimeRepository
}

func NewJWTExternal(tr repository.TimeRepository) *jwtExternal {
	return &jwtExternal{tr: tr}
}

func (j *jwtExternal) CreateJWTToken(ctx context.Context, id model.UserID, name string) (*string, error) {
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
