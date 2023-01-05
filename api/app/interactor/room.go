package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.RoomUseCase = (*roomInteractor)(nil)

var RoomInteractorSet = wire.NewSet(
	NewRoomInterractor,
	wire.Bind(new(usecase.RoomUseCase), new(*roomInteractor)),
)

type roomInteractor struct {
	rr repository.RoomRepository
}

func NewRoomInterractor(rr repository.RoomRepository) *roomInteractor {
	return &roomInteractor{rr: rr}
}

func (ri *roomInteractor) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
	if len(userIDs) == 0 {
		return nil, myerrors.ErrBadRequestEmptyArray
	}

	ru, err := ri.rr.Create(ctx, name, isGroup, userIDs)
	if err != nil {
		return nil, err
	}
	return ru, nil
}

func (ri *roomInteractor) Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
	irs, nID, err := ri.rr.Index(ctx, id, nextID)
	if err != nil {
		return nil, nil, err
	}

	return irs, nID, nil
}

func (ri *roomInteractor) Exists(ctx context.Context, myID model.UserID, id model.UserID) (*model.Room, error) {
	room, err := ri.rr.GetNoneGroup(ctx, myID, id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (ri *roomInteractor) Show(ctx context.Context, id model.RoomID, myID model.UserID) (*model.RoomUser, error) {
	ru, err := ri.rr.Show(ctx, id)
	if err != nil {
		return nil, err
	}

	// DMの際は自分が存在しているルームかチェックする
	if !ru.Room.IsGroup && !ri.isFineRoom(myID, ru.Users) {
		return nil, myerrors.ErrBadRequestNoPermission
	}

	return ru, nil
}

func (ri *roomInteractor) isFineRoom(id model.UserID, users []*model.ShowUser) bool {
	for _, u := range users {
		if u.ID == id {
			return true
		}
	}
	return false
}
