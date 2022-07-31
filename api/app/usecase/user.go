package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kod-source/docker-goa-next/app/model"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email, password string) (*model.User, error)
	CreateJWTToken(ctx context.Context, id int, name string) (*string, error)
}

type userUseCase struct{}

func NewUserUseCase() UserUseCase {
	return userUseCase{}
}

func (u userUseCase) GetUser(ctx context.Context, id int) (*model.User, error) {
	if model.MockUser.ID != id {
		return nil, errors.New("404 error")
	}

	return &model.MockUser, nil
}

func (u userUseCase) GetUserByEmail(ctx context.Context, email, password string) (*model.User, error) {
	if model.MockUser.Email != email || model.MockUser.Password != password {
		return nil, errors.New("404 error")
	}

	return &model.MockUser, nil
}

func (u userUseCase) CreateJWTToken(ctx context.Context, id int, name string) (*string, error) {
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
