package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.ThreadUsecase = (*threadInteractor)(nil)

var ThreadUsecaseSet = wire.NewSet(
	NewThreadUsecase,
	wire.Bind(new(usecase.ThreadUsecase), new(*threadInteractor)),
)

type threadInteractor struct {
}

func NewThreadUsecase() *threadInteractor {
	return &threadInteractor{}
}

func (ti *threadInteractor) Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
	return nil, nil
}
