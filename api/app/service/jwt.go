package service

import "context"

type JWTService interface {
	CreateJWTToken(ctx context.Context, id int, name string) (*string, error)
}
