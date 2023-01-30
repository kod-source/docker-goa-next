package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.ThreadUsecase = (*threadInteractor)(nil)

var ThreadUsecaseSet = wire.NewSet(
	NewThreadUsecase,
	wire.Bind(new(usecase.ThreadUsecase), new(*threadInteractor)),
)

type threadInteractor struct {
	tr repository.ThreadRepository
}

func NewThreadUsecase(tr repository.ThreadRepository) *threadInteractor {
	return &threadInteractor{tr: tr}
}

func (ti *threadInteractor) Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
	if text == "" {
		return nil, myerrors.ErrBadRequestSting
	}
	if roomID == 0 || userID == 0 {
		return nil, myerrors.ErrBadRequestInt
	}

	return ti.tr.Create(ctx, text, roomID, userID, img)
}

func (ti *threadInteractor) Delete(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
	if err := ti.tr.Delete(ctx, myID, threadID); err != nil {
		return err
	}
	return nil
}

func (ti *threadInteractor) GetThreadsByRoom(ctx context.Context, roomID model.RoomID, nextID model.ThreadID) ([]*model.IndexThread, *int, error) {
	its, nID, err := ti.tr.GetThreadsByRoom(ctx, roomID, nextID)
	if err != nil {
		return nil, nil, err
	}

	return its, nID, nil
}
