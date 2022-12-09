package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.UserUseCase = (*MockUserUsecase)(nil)

type MockUserUsecase struct {
	GetUserFunc        func(ctx context.Context, id int) (*model.User, error)
	GetUserByEmailFunc func(ctx context.Context, email, password string) (*model.User, error)
	CreateJWTTokenFunc func(ctx context.Context, id int, name string) (*string, error)
	SignUpFunc         func(ctx context.Context, name, email, password string, avatar *string) (*model.User, error)
}

func (m *MockUserUsecase) GetUser(ctx context.Context, id int) (*model.User, error) {
	return m.GetUserFunc(ctx, id)
}

func (m *MockUserUsecase) GetUserByEmail(ctx context.Context, email, password string) (*model.User, error) {
	return m.GetUserByEmailFunc(ctx, email, password)
}

func (m *MockUserUsecase) CreateJWTToken(ctx context.Context, id int, name string) (*string, error) {
	return m.CreateJWTTokenFunc(ctx, id, name)
}

func (m *MockUserUsecase) SignUp(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
	return m.SignUpFunc(ctx, name, email, password, avatar)
}
