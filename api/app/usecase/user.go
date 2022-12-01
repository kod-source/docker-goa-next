package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email, password string) (*model.User, error)
	CreateJWTToken(ctx context.Context, id int, name string) (*string, error)
	SignUp(ctx context.Context, name, email, password string, avatar *string) (*model.User, error)
}
