package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.GoogleUsecase = (*MockGoogleUsecase)(nil)

type MockGoogleUsecase struct {
	GetLoginURLFunc         func(state string) string
	GetOrCreateUserInfoFunc func(ctx context.Context, code string) (*model.User, error)
}

func (m *MockGoogleUsecase) GetLoginURL(state string) string {
	return m.GetLoginURLFunc(state)
}

func (m *MockGoogleUsecase) GetOrCreateUserInfo(ctx context.Context, code string) (*model.User, error) {
	return m.GetOrCreateUserInfoFunc(ctx, code)
}
