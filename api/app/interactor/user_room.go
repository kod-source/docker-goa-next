package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.UserRoomUseCase = (*userRoomInteractor)(nil)

var UserRoomInteractorSet = wire.NewSet(
	NewUserRoomUsecase,
	wire.Bind(new(usecase.UserRoomUseCase), new(*userRoomInteractor)),
)

type userRoomInteractor struct {
	urr repository.UserRoomRepository
}

func NewUserRoomUsecase(urr repository.UserRoomRepository) *userRoomInteractor {
	return &userRoomInteractor{urr: urr}
}

func (uri *userRoomInteractor) InviteRoom(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
	if roomID == 0 || userID == 0 {
		return nil, myerrors.ErrBadRequestInt
	}
	ur, err := uri.urr.Create(ctx, roomID, userID)
	if err != nil {
		return nil, err
	}
	return ur, nil
}
