package mock

import "github.com/kod-source/docker-goa-next/app/service"

var _ service.GoogleService = (*MockGoogleService)(nil)

type MockGoogleService struct {
	GetLoginURLFunc func(state string) string
}

func (m *MockGoogleService) GetLoginURL(state string) string {
	return m.GetLoginURLFunc(state)
}
