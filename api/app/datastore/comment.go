package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
)

var _ repository.CommentRepository = (*commentDatastore)(nil)

var CommentDatastoreSet = wire.NewSet(
	NewCommentDatastore,
	wire.Bind(new(repository.CommentRepository), new(*commentDatastore)),
)

type commentDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewCommentDatastore(db *sql.DB, tr repository.TimeRepository) *commentDatastore {
	return &commentDatastore{db: db, tr: tr}
}

func (c *commentDatastore) Create(ctx context.Context, postID, userID int, text string, img *string) (*model.CommentWithUser, error) {
	var commentWithUser model.CommentWithUser
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	ins, err := tx.Prepare(
		"INSERT INTO comment(`post_id`, `user_id`, `text`, `img`, `created_at`, `updated_at`) VALUES(?,?,?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.Exec(postID, userID, text, img, c.tr.Now(), c.tr.Now())
	if err != nil {
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(`
		SELECT c.id, c.post_id, c.user_id, c.text, c.img, c.created_at, c.updated_at, u.id, u.name, u.avatar
		FROM comment as c
		INNER JOIN user as u
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
		return nil, err
	}

	return &commentWithUser, tx.Commit()
}

func (c *commentDatastore) ShowByPostID(ctx context.Context, postID int) ([]*model.CommentWithUser, error) {
	var commentsWithUsers []*model.CommentWithUser
	rows, err := c.db.Query(`
		SELECT c.id, c.post_id, c.user_id, c.text, c.img, c.created_at, c.updated_at, u.id, u.name, u.avatar
		FROM comment as c
		INNER JOIN user as u
		ON c.user_id = u.id
		WHERE c.post_id = ?
		ORDER BY c.created_at DESC
	`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commentWithUser model.CommentWithUser

		err := rows.Scan(
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
			return nil, err
		}
		commentsWithUsers = append(commentsWithUsers, &commentWithUser)
	}
	if len(commentsWithUsers) == 0 {
		return nil, sql.ErrNoRows
	}

	return commentsWithUsers, nil
}

func (c *commentDatastore) Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error) {
	var comment model.Comment
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	upd, err := tx.Prepare("UPDATE `comment` set `text` = ?, `img` = ?, `updated_at` = ? WHERE `id` = ?")
	if err != nil {
		return nil, err
	}
	_, err = upd.Exec(text, img, c.tr.Now(), id)
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(
		"SELECT `id`, `post_id`, `user_id`, `text`, `img`, `created_at`, `updated_at` FROM `comment` WHERE `id` = ?", id,
	).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Text,
		&comment.Img,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &comment, tx.Commit()
}

func (c *commentDatastore) Delete(ctx context.Context, id int) error {
	stmt, err := c.db.Prepare("DELETE FROM `comment` WHERE `id` = ?")
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
