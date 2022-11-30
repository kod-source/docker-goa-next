package datastore

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ JWTDatastore = (*jwtDatastore)(nil)

var JWTDatastoreSet = wire.NewSet(
	NewJWTDatastore,
	wire.Bind(new(JWTDatastore), new(*jwtDatastore)),
)

type JWTDatastore interface {
	CreateJWTToken(ctx context.Context, id int, name string) (*string, error)
}

type jwtDatastore struct {
	tr repository.TimeRepository
}

func NewJWTDatastore(tr repository.TimeRepository) *jwtDatastore {
	return &jwtDatastore{tr: tr}
}

func (j *jwtDatastore) CreateJWTToken(ctx context.Context, id int, name string) (*string, error) {
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
