package interactor

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTInteractor interface {
	CreateJWTToken(ctx context.Context, id int, name string) (*string, error)
}

type jwtInteractor struct{}

func NewJWTInteractor() JWTInteractor {
	return jwtInteractor{}
}

func (j jwtInteractor) CreateJWTToken(ctx context.Context, id int, name string) (*string, error) {
	claims := jwt.MapClaims{
		"sub":       "auth jwt",
		"user_id":   id,
		"user_name": name,
		"scope":     "api:access",
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &token, nil
}
