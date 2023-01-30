package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.ContentUsecase = (*contentInteractor)(nil)

var ContentUsecaseSet = wire.NewSet(
	NewContentUsecase,
	wire.Bind(new(usecase.ContentUsecase), new(*contentInteractor)),
)

type contentInteractor struct {
	cr repository.ContentRepository
}

func NewContentUsecase(cr repository.ContentRepository) *contentInteractor {
	return &contentInteractor{cr: cr}
}

func (ci *contentInteractor) Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
	return ci.cr.Delete(ctx, myID, contentID)
}

func (ci *contentInteractor) Create(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
	if text == "" {
		return nil, myerrors.ErrBadRequestSting
	}
	if threadID == 0 || myID == 0 {
		return nil, myerrors.ErrBadRequestInt
	}

	cu, err := ci.cr.Create(ctx, text, threadID, myID, img)
	if err != nil {
		return nil, err
	}
	return cu, nil
}
