package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id model.UserID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email, password string) (*model.User, error)
	CreateJWTToken(ctx context.Context, id model.UserID, name string) (*string, error)
	SignUp(ctx context.Context, name, email, password string, avatar *string) (*model.User, error)
	IndexUser(ctx context.Context, id model.UserID) ([]*model.User, error)
}
