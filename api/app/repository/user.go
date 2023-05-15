package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

//go:generate mockgen -source=./user.go -package=mock_repository -destination=./mock/user_repository.go

type UserRepository interface {
	GetUser(ctx context.Context, id model.UserID) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, name, email, password string, avatar *string) (*model.User, error)
	IndexUser(ctx context.Context, myID model.UserID) ([]*model.User, error)
}
