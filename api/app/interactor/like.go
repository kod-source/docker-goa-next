package interactor

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
)

type LikeInteractor interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
}

type likeInteractor struct {
	db *sql.DB
}

func NewLikeInteractor(db *sql.DB) LikeInteractor {
	return &likeInteractor{db: db}
}

func (l *likeInteractor) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	var like model.Like
	tx, err := l.db.Begin()
	if err != nil {
		return nil, err
	}
	ins, err := tx.Prepare(
		"INSERT INTO likes(`user_id`, `post_id`) VALUES(?, ?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.Exec(userID, postID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(
		"SELECT `id`, `user_id`, `post_id` FROM `likes` WHERE `id` = ?", id,
	).Scan(
		&like.ID,
		&like.UserID,
		&like.PostID,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &like, tx.Commit()
}
