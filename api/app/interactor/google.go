package interactor

import (
	"context"
	"time"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
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

func (gi *googleInteractor) GetOrCreateUserInfo(ctx context.Context, code string) (*model.User, error) {
	_, err := gi.gs.GetUserInfo(ctx, code)
	if err != nil {
		return nil, err
	}
	// ***
	// ToDo for me
	// - codeから検証したUserが存在する確認する
	// - 存在しない場合はユーザーを作成する
	// ***
	mockUser := &model.User{
		ID:        1000,
		Name:      "mock_user",
		Email:     "mock@example.com",
		Avatar:    nil,
		CreatedAt: time.Date(2022, 1, 1, 0, 0, 0, 0, nil),
	}
	return mockUser, nil
}
