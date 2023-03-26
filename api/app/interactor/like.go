package interactor

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/usecase"
)

var _ usecase.LikeUsecase = (*likeInteractor)(nil)

var LikeInteractorSet = wire.NewSet(
	NewLikeInteractor,
	wire.Bind(new(usecase.LikeUsecase), new(*likeInteractor)),
)

type likeInteractor struct {
	lr repository.LikeRepository
	db *sql.DB
}

func NewLikeInteractor(lr repository.LikeRepository, db *sql.DB) *likeInteractor {
	return &likeInteractor{lr: lr, db: db}
}

func (l *likeInteractor) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	if userID == 0 || postID == 0 {
		return nil, myerrors.ErrBadRequestInt
	}
	like, err := l.lr.Create(ctx, tx, userID, postID)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return like, nil
}

func (l *likeInteractor) Delete(ctx context.Context, userID, postID int) error {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if userID == 0 || postID == 0 {
		return myerrors.ErrBadRequestInt
	}
	err = l.lr.Delete(ctx, tx, userID, postID)
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (l *likeInteractor) GetPostIDs(ctx context.Context, userID int) ([]int, error) {
	tx, err := l.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	postIDs, err := l.lr.GetPostIDs(ctx, tx, userID)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return postIDs, nil
}
