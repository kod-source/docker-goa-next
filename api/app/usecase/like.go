package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
)

type LikeUsecase interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
}

type likeUsecase struct {
	li interactor.LikeInteractor
}

func NewLikeUsecase(li interactor.LikeInteractor) LikeUsecase {
	return &likeUsecase{li: li}
}

func (l *likeUsecase) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	if userID == 0 || postID == 0 {
		return nil, myerrors.BadRequestIntError
	}
	like, err := l.li.Create(ctx, userID, postID)
	if err != nil {
		return nil, err
	}

	return like, nil
}
