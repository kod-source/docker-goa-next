package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type GoogleUsecase interface {
	GetLoginURL(state string) string
	GetOrCreateUserInfo(ctx context.Context, code string) (*model.User, error)
}
