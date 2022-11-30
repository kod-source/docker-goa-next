package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
)

var _ LikeDatastore = (*likeDatastore)(nil)

var LikeDatastoreSet = wire.NewSet(
	NewLikeDatastore,
	wire.Bind(new(LikeDatastore), new(*likeDatastore)),
)

type LikeDatastore interface {
	Create(ctx context.Context, userID, postID int) (*model.Like, error)
	Delete(ctx context.Context, userID, postID int) error
	GetPostIDs(ctx context.Context, userID int) ([]int, error)
}

type likeDatastore struct {
	db *sql.DB
}

func NewLikeDatastore(db *sql.DB) LikeDatastore {
	return &likeDatastore{db: db}
}

func (l *likeDatastore) Create(ctx context.Context, userID, postID int) (*model.Like, error) {
	var like model.Like
	tx, err := l.db.Begin()
	if err != nil {
		return nil, err
	}
	ins, err := tx.Prepare(
		"INSERT INTO `like`(`user_id`, `post_id`) VALUES(?, ?)",
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
		"SELECT `id`, `user_id`, `post_id` FROM `like` WHERE `id` = ?", id,
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

func (l *likeDatastore) Delete(ctx context.Context, userID, postID int) error {
	stmt, err := l.db.Prepare("DELETE FROM `like` WHERE `user_id` = ? AND `post_id` = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(userID, postID)
	if err != nil {
		return err
	}

	return nil
}

func (l *likeDatastore) GetPostIDs(ctx context.Context, userID int) ([]int, error) {
	rows, err := l.db.Query("SELECT `post_id` FROM `like` WHERE `user_id` = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var postIDs []int
	for rows.Next() {
		var postID int
		err = rows.Scan(
			&postID,
		)
		if err != nil {
			return nil, err
		}

		postIDs = append(postIDs, postID)
	}

	return postIDs, nil
}
