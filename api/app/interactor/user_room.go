package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.UserRoomUseCase = (*userRoomInteractor)(nil)

var UserRoomInteractorSet = wire.NewSet(
	NewUserRoomUsecase,
	wire.Bind(new(usecase.UserRoomUseCase), new(*userRoomInteractor)),
)

type userRoomInteractor struct {
}

func NewUserRoomUsecase() *userRoomInteractor {
	return &userRoomInteractor{}
}

func (uri *userRoomInteractor) InviteRoom(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
	return nil, nil
}
