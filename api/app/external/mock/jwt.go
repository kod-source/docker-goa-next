package mock

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/service"
)

var _ service.JWTService = (*MockJWTService)(nil)

type MockJWTService struct {
	CreateJWTTokenFunc func(ctx context.Context, id model.UserID, name string) (*string, error)
}

func (m *MockJWTService) CreateJWTToken(ctx context.Context, id model.UserID, name string) (*string, error) {
	return m.CreateJWTTokenFunc(ctx, id, name)
}
