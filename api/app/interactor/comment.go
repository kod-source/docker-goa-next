package interactor

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

type CommentInteractor interface {
	Create(ctx context.Context, postID, userID int, text string, img *string) (*model.CommentWithUser, error)
	ShowByPostID(ctx context.Context, postID int) ([]*model.Comment, error)
	Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error)
	Delete(ctx context.Context, id int) error
}

type commentInteractor struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewCommentInteractor(db *sql.DB, tr repository.TimeRepository) CommentInteractor {
	return &commentInteractor{db: db, tr: tr}
}

func (c *commentInteractor) Create(ctx context.Context, postID, userID int, text string, img *string) (*model.CommentWithUser, error) {
	var commentWithUser model.CommentWithUser
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}
	ins, err := tx.Prepare(
		"INSERT INTO comments(`post_id`, `user_id`, `text`, `img`, `created_at`, `updated_at`) VALUES(?,?,?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.Exec(postID, userID, text, img, c.tr.Now(), c.tr.Now())
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(`
		SELECT c.id, c.post_id, c.user_id, c.text, c.img, c.created_at, c.updated_at, u.id, u.name, u.avatar
		FROM comments as c
		INNER JOIN users as u
		ON c.user_id = u.id
		WHERE c.id = ?
	`, lastID).Scan(
		&commentWithUser.Comment.ID,
		&commentWithUser.Comment.PostID,
		&commentWithUser.Comment.UserID,
		&commentWithUser.Comment.Text,
		&commentWithUser.Comment.Img,
		&commentWithUser.Comment.CreatedAt,
		&commentWithUser.Comment.UpdatedAt,
		&commentWithUser.User.ID,
		&commentWithUser.User.Name,
		&commentWithUser.User.Avatar,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &commentWithUser, tx.Commit()
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

func (c *commentInteractor) Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error) {
	var comment model.Comment
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}
	upd, err := tx.Prepare("UPDATE `comments` set `text` = ?, `img` = ?, `updated_at` = ? WHERE `id` = ?")
	if err != nil {
		return nil, err
	}
	_, err = upd.Exec(text, img, c.tr.Now(), id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.QueryRow(
		"SELECT `id`, `post_id`, `text`, `img`, `created_at`, `updated_at` FROM `comments` WHERE `id` = ?", id,
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

func (c *commentInteractor) Delete(ctx context.Context, id int) error {
	stmt, err := c.db.Prepare("DELETE FROM `comments` WHERE `id` = ?")
	if err != nil {
		return err
	}
	r, err := stmt.Exec(id)
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
