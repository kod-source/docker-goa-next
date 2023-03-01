package interactor

import (
	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/service"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.GoogleUsecase = (*googleInteractor)(nil)

var GoogleUsecaseSet = wire.NewSet(
	NewGoogleUseCase,
	wire.Bind(new(usecase.GoogleUsecase), new(*googleInteractor)),
)

type googleInteractor struct {
	gs service.GoogleService
}

func NewGoogleUseCase(gs service.GoogleService) *googleInteractor {
	return &googleInteractor{gs: gs}
}

func (gi *googleInteractor) GetLoginURL(state string) string {
	return gi.gs.GetLoginURL(state)
}
