package interactor

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

type CommentInteractor interface {
	Create(ctx context.Context, postID int, text string, img *string) (*model.Comment, error)
}

type commentInteractor struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewCommentInteractor(db *sql.DB, tr repository.TimeRepository) CommentInteractor {
	return &commentInteractor{db: db, tr: tr}
}

func (c *commentInteractor) Create(ctx context.Context, postID int, text string, img *string) (*model.Comment, error) {
	var comment model.Comment
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}
	ins, err := tx.Prepare(
		"INSERT INTO comments(`post_id`, `text`, `img`, `created_at`, `updated_at`) VALUES(?,?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.Exec(postID, text, img, c.tr.Now(), c.tr.Now())
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(
		"SELECT `id`, `post_id`, `text`, `img`, `created_at`, `updated_at` FROM comments WHERE id = ?", lastID,
	).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.Text,
		&comment.Img,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &comment, tx.Commit()
}
