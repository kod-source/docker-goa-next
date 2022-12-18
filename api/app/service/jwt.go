package service

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type JWTService interface {
	CreateJWTToken(ctx context.Context, id model.UserID, name string) (*string, error)
}
