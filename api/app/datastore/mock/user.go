package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.UserRepository = (*MockUserRepository)(nil)

type MockUserRepository struct {
	GetUserFunc        func(ctx context.Context, id model.UserID) (*model.User, error)
	GetUserByEmailFunc func(ctx context.Context, email string) (*model.User, error)
	CreateUserFunc     func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error)
}

func (m *MockUserRepository) GetUser(ctx context.Context, id model.UserID) (*model.User, error) {
	return m.GetUserFunc(ctx, id)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return m.GetUserByEmailFunc(ctx, email)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
	return m.CreateUserFunc(ctx, name, email, password, avatar)
}
