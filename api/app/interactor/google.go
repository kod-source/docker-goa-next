package interactor

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
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
	ur repository.UserRepository
}

func NewGoogleUseCase(gs service.GoogleService, ur repository.UserRepository) *googleInteractor {
	return &googleInteractor{gs: gs, ur: ur}
}

func (gi *googleInteractor) GetLoginURL(state string) string {
	return gi.gs.GetLoginURL(state)
}

func (gi *googleInteractor) GetOrCreateUserInfo(ctx context.Context, code string) (*model.User, error) {
	gu, err := gi.gs.GetUserInfo(ctx, code)
	if err != nil {
		return nil, err
	}
	user, err := gi.ur.GetUserByEmail(ctx, gu.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			cu, err := gi.ur.CreateUser(ctx, gu.Name, gu.Email, "", &gu.Picture)
			if err != nil {
				return nil, err
			}
			return cu, nil
		}
		return nil, err
	}
	return user, nil
}
