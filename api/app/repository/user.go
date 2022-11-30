package repository

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, name, email, password string, avatar *string) (*model.User, error)
}
