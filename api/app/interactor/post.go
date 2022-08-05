package interactor

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/model"
)

type PostInteractor interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.Post, error)
	ShowAll(ctx context.Context) ([]*model.IndexPost, error)
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

func (p postInteractor) ShowAll(ctx context.Context) ([]*model.IndexPost, error) {
	var indexPosts []*model.IndexPost
	rows, err := p.db.Query(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		var user model.User

		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Title,
			&post.Img,
			&post.CreatedAt,
			&post.UpdatedAt,
			&user.Name,
			&user.Avatar,
		)
		if err != nil {
			return nil, err
		}

		indexPosts = append(indexPosts, &model.IndexPost{
		Post: model.Post{
			ID:        post.ID,
			UserID:    post.UserID,
			Title:     post.Title,
			Img:       post.Img,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		},
		User: model.User{
			Name:      user.Name,
			Avatar:    user.Avatar,
		},
	})
	}

	return indexPosts, nil
}
