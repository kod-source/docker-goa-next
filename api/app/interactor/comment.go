package interactor

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

type CommentInteractor interface {
	Create(ctx context.Context, postID int, text string, img *string) (*model.Comment, error)
	ShowByPostID(ctx context.Context, postID int) ([]*model.Comment, error)
	Update(ctx context.Context, id int, text, img string) (*model.Comment, error)
	Delete(ctx context.Context, id int) error
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

func (c *commentInteractor) ShowByPostID(ctx context.Context, postID int) ([]*model.Comment, error) {
	var comments []*model.Comment
	rows, err := c.db.Query(
		"SELECT `id`, `post_id`, `text`, `img`, `created_at`, `updated_at` FROM comments WHERE post_id = ?", postID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment model.Comment

		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.Text,
			&comment.Img,
			&comment.CreatedAt,
			&comment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if len(comments) == 0 {
		return nil, sql.ErrNoRows
	}

	return comments, nil
}

func (c *commentInteractor) Update(ctx context.Context, id int, text, img string) (*model.Comment, error) {
	return nil, nil
}

func (c *commentInteractor) Delete(ctx context.Context, id int) error {
	return nil
}
