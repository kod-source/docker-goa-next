package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/pointer"
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

	var comment schema.Comment
	var user schema.User
	err = tx.QueryRow(`
		SELECT c.id, c.post_id, c.user_id, c.text, c.img, c.created_at, c.updated_at, u.id, u.name, u.avatar
		FROM comment as c
		INNER JOIN user as u
		ON c.user_id = u.id
		WHERE c.id = ?
	`, lastID).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Text,
		&comment.Img,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&user.ID,
		&user.Name,
		&user.Avatar,
	)
	if err != nil {
		return nil, err
	}

	return c.convertSchemaToModelCommentWithUser(&comment, &user), tx.Commit()
}

func (c *commentDatastore) ShowByPostID(ctx context.Context, postID int) ([]*model.CommentWithUser, error) {
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

	var commentsWithUsers []*model.CommentWithUser
	for rows.Next() {
		var comment schema.Comment
		var user schema.User

		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.UserID,
			&comment.Text,
			&comment.Img,
			&comment.CreatedAt,
			&comment.UpdatedAt,
			&user.ID,
			&user.Name,
			&user.Avatar,
		)
		if err != nil {
			return nil, err
		}
		commentsWithUsers = append(commentsWithUsers, c.convertSchemaToModelCommentWithUser(&comment, &user))
	}
	if len(commentsWithUsers) == 0 {
		return nil, sql.ErrNoRows
	}

	return commentsWithUsers, nil
}

func (c *commentDatastore) Update(ctx context.Context, id int, text string, img *string) (*model.Comment, error) {
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

	var comment schema.Comment
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

	return c.convertSchemaToModelComment(&comment), tx.Commit()
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

func (c *commentDatastore) convertSchemaToModelCommentWithUser(comment *schema.Comment, user *schema.User) *model.CommentWithUser {
	cu := &model.CommentWithUser{
		Comment: model.Comment{
			ID:        int(comment.ID),
			PostID:    int(comment.PostID),
			UserID:    int(comment.UserID),
			Text:      comment.Text,
			CreatedAt: pointer.PtrOrNil(comment.CreatedAt),
			UpdatedAt: pointer.PtrOrNil(comment.UpdatedAt),
		},
		User: model.User{
			ID:        int(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt,
		},
	}
	if comment.Img.Valid {
		cu.Comment.Img = &comment.Img.String
	}
	if user.Avatar.Valid {
		cu.User.Avatar = &user.Avatar.String
	}

	return cu
}

func (c *commentDatastore) convertSchemaToModelComment(comment *schema.Comment) *model.Comment {
	mc := &model.Comment{
		ID:        int(comment.ID),
		PostID:    int(comment.PostID),
		UserID:    int(comment.UserID),
		Text:      comment.Text,
		CreatedAt: pointer.PtrOrNil(comment.CreatedAt),
		UpdatedAt: pointer.PtrOrNil(comment.UpdatedAt),
	}
	if comment.Img.Valid {
		mc.Img = &comment.Img.String
	}

	return mc
}
