package mock

import "github.com/kod-source/docker-goa-next/app/usecase"

var _ usecase.GoogleUsecase = (*MockGoogleUsecase)(nil)

type MockGoogleUsecase struct {
	GetLoginURLFunc func(state string) string
}

func (m *MockGoogleUsecase) GetLoginURL(state string) string {
	return m.GetLoginURLFunc(state)
}
