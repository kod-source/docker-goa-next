package interactor

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/model"
)

type PostInteractor interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.Post, error)
}

type postInteractor struct {
	db *sql.DB
}

func NewPostInteractor(db *sql.DB) PostInteractor {
	return postInteractor{db: db}
}

func (p postInteractor) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.Post, error) {
	var post model.Post
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	ins, err := tx.Prepare(
		"INSERT INTO posts(`user_id`, `title`, `img`, `created_at`, `updated_at`) VALUES(?,?,?,?,?)",
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	res, err := ins.Exec(userID, title, img, time.Now(), time.Now())
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(
		"SELECT `id`, `user_id`, `title`, `img`, `created_at`, `updated_at` FROM `posts` WHERE `id` = ?", lastID,
	).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Img,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &post, nil
}
