package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email, password string) (*model.User, error)
	CreateJWTToken(ctx context.Context, id int, name string) (*string, error)
}

type userUseCase struct {
	ui interactor.UserInteractor
	ji interactor.JWTInteractor
}

func NewUserUseCase(ui interactor.UserInteractor, ji interactor.JWTInteractor) UserUseCase {
	return userUseCase{ui: ui, ji: ji}
}

func (u userUseCase) GetUser(ctx context.Context, id int) (*model.User, error) {
	user, err := u.ui.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUseCase) GetUserByEmail(ctx context.Context, email, password string) (*model.User, error) {
	user, err := u.ui.GetUserByEmail(ctx, email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUseCase) CreateJWTToken(ctx context.Context, id int, name string) (*string, error) {
	token, nil := u.ji.CreateJWTToken(ctx, id, name)
	return token, nil
}
