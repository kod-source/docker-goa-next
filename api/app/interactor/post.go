package interactor

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

type PostInteractor interface {
	CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error)
	ShowAll(ctx context.Context) ([]*model.IndexPost, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error)
}

type postInteractor struct {
	db *sql.DB
}

func NewPostInteractor(db *sql.DB) PostInteractor {
	return postInteractor{db: db}
}

func (p postInteractor) CreatePost(ctx context.Context, userID int, title string, img *string) (*model.IndexPost, error) {
	var indexPost model.IndexPost
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
	err = tx.QueryRow(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
		WHERE p.id = ?
	`, lastID).Scan(
		&indexPost.Post.ID,
		&indexPost.Post.UserID,
		&indexPost.Post.Title,
		&indexPost.Post.Img,
		&indexPost.Post.CreatedAt,
		&indexPost.Post.UpdatedAt,
		&indexPost.User.Name,
		&indexPost.User.Avatar,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &indexPost, nil
}

func (p postInteractor) ShowAll(ctx context.Context) ([]*model.IndexPost, error) {
	var indexPosts []*model.IndexPost
	rows, err := p.db.Query(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
		ORDER BY p.created_at DESC
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
				Name:   user.Name,
				Avatar: user.Avatar,
			},
		})
	}

	return indexPosts, nil
}

func (p postInteractor) Delete(ctx context.Context, id int) error {
	stmt, err := p.db.Prepare("DELETE FROM `posts` WHERE `id` = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (p postInteractor) Update(ctx context.Context, id int, title string, img *string) (*model.IndexPost, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	upd, err := tx.Prepare("UPDATE `posts` set `title` = ?, `img` = ?, `updated_at` = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	ti := repository.NewTimeRepositoy()
	result, err := upd.Exec(title, img, ti.Now(), id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	var indexPost model.IndexPost
	err = tx.QueryRow(`
		SELECT p.id, p.user_id, p.title, p.img, p.created_at, p.updated_at, u.name, u.avatar
		FROM posts as p
		INNER JOIN users as u
		ON p.user_id = u.id
		WHERE p.id = ?
	`, rowsAffected).Scan(
		&indexPost.Post.ID,
		&indexPost.Post.UserID,
		&indexPost.Post.Title,
		&indexPost.Post.Img,
		&indexPost.Post.CreatedAt,
		&indexPost.Post.UpdatedAt,
		&indexPost.User.Name,
		&indexPost.User.Avatar,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &indexPost, nil
}
