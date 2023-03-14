package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/service"
)

var _ service.GoogleService = (*MockGoogleService)(nil)

type MockGoogleService struct {
	GetLoginURLFunc func(state string) string
	GetUserInfoFunc func(ctx context.Context, code string) (*model.GoogleUser, error)
}

func (m *MockGoogleService) GetLoginURL(state string) string {
	return m.GetLoginURLFunc(state)
}

func (m *MockGoogleService) GetUserInfo(ctx context.Context, code string) (*model.GoogleUser, error) {
	return m.GetUserInfoFunc(ctx, code)
}
