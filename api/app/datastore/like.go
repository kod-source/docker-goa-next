package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
)

var _ repository.LikeRepository = (*likeDatastore)(nil)

var LikeDatastoreSet = wire.NewSet(
	NewLikeDatastore,
	wire.Bind(new(repository.LikeRepository), new(*likeDatastore)),
)

type likeDatastore struct {
}

func NewLikeDatastore() *likeDatastore {
	return &likeDatastore{}
}

func (l *likeDatastore) Create(ctx context.Context, tx *sql.Tx, userID, postID int) (*model.Like, error) {
	ins, err := tx.Prepare(
		"INSERT INTO `like`(`user_id`, `post_id`) VALUES(?, ?)",
	)
	if err != nil {
		return nil, err
	}
	defer ins.Close()

	res, err := ins.Exec(userID, postID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	var like schema.Like
	err = tx.QueryRow(
		"SELECT `id`, `user_id`, `post_id` FROM `like` WHERE `id` = ?", id,
	).Scan(
		&like.ID,
		&like.UserID,
		&like.PostID,
	)
	if err != nil {
		return nil, err
	}

	return &model.Like{
		ID:     int(like.ID),
		UserID: int(like.UserID),
		PostID: int(like.PostID),
	}, nil
}

func (l *likeDatastore) Delete(ctx context.Context, tx *sql.Tx, userID, postID int) error {
	stmt, err := tx.Prepare("DELETE FROM `like` WHERE `user_id` = ? AND `post_id` = ?")
	if err != nil {
		return err
	}
	r, err := stmt.Exec(userID, postID)
	if err != nil {
		return err
	}
	i, err := r.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (l *likeDatastore) GetPostIDs(ctx context.Context, tx *sql.Tx, userID int) ([]int, error) {
	rows, err := tx.Query("SELECT `post_id` FROM `like` WHERE `user_id` = ?", userID)
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
