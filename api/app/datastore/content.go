package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
)

var _ repository.ContentRepository = (*contentDatastore)(nil)

var ContentRepositorySet = wire.NewSet(
	NewContentRepository,
	wire.Bind(new(repository.ContentRepository), new(*contentDatastore)),
)

type contentDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewContentRepository(db *sql.DB, tr repository.TimeRepository) *contentDatastore {
	return &contentDatastore{db: db, tr: tr}
}

func (cd *contentDatastore) Delete(ctx context.Context, myID model.UserID, contentID model.ContentID) error {
	tx, err := cd.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var user schema.User
	if err := tx.QueryRowContext(
		ctx,
		"SELECT `user_id` FROM `content` WHERE `id` = ?",
		contentID,
	).Scan(
		&user.ID,
	); err != nil {
		return err
	}
	if uint64(myID) != user.ID {
		return myerrors.ErrBadRequestNoPermission
	}

	stmt, err := tx.PrepareContext(ctx, "DELETE FROM `content` WHERE `id` = ?")
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, contentID)
	if err != nil {
		return err
	}
	i, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if i == 0 {
		return sql.ErrNoRows
	}

	return tx.Commit()
}

func (cd *contentDatastore) Create(ctx context.Context, text string, threadID model.ThreadID, myID model.UserID, img *string) (*model.ContentUser, error) {
	tx, err := cd.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		ctx,
		"INSERT INTO `content` (`user_id`, `thread_id`, `text`, `created_at`, `updated_at`, `img`) VALUES(?,?,?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, myID, threadID, text, cd.tr.Now(), cd.tr.Now(), img)
	if err != nil {
		return nil, err
	}
	contentID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var content schema.Content
	var user schema.User
	query := "SELECT `c`.`id`, `c`.`user_id`, `c`.`thread_id`, `c`.`text`, `c`.`created_at`, `c`.`updated_at`, `c`.`img`, " +
		"`u`.`id`, `u`.`name`, `u`.`created_at`, `u`.`avatar` " +
		"FROM `content` AS `c` " +
		"INNER JOIN `user` AS `u` " +
		"ON `c`.`user_id` = `u`.`id` " +
		"WHERE `c`.`id` = ?"
	if err := tx.QueryRowContext(ctx, query, contentID).Scan(
		&content.ID,
		&content.UserID,
		&content.ThreadID,
		&content.Text,
		&content.CreatedAt,
		&content.UpdatedAt,
		&content.Img,
		&user.ID,
		&user.Name,
		&user.CreatedAt,
		&user.Avatar,
	); err != nil {
		return nil, err
	}

	return toContentUser(content, user), tx.Commit()
}

// GetByThread ...
func (cd *contentDatastore) GetByThread(ctx context.Context, threadID model.ThreadID) ([]*model.ContentUser, error) {
	tx, err := cd.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "SELECT `c`.`id`, `c`.`user_id`, `c`.`thread_id`, `c`.`text`, `c`.`created_at`, `c`.`updated_at`, `c`.`img`, " +
		"`u`.`id`, `u`.`name`, `u`.`created_at`, `u`.`avatar` " +
		"FROM `content` AS `c` " +
		"INNER JOIN `user` AS `u` " +
		"ON `c`.`user_id` = `u`.`id` " +
		"WHERE `c`.`thread_id` = ? " +
		"ORDER BY `c`.`created_at`"
	rows, err := tx.QueryContext(ctx, query, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cus []*model.ContentUser
	for rows.Next() {
		var c schema.Content
		var u schema.User
		if err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.ThreadID,
			&c.Text,
			&c.CreatedAt,
			&c.UpdatedAt,
			&c.Img,
			&u.ID,
			&u.Name,
			&u.CreatedAt,
			&u.Avatar,
		); err != nil {
			return nil, err
		}

		cu := &model.ContentUser{
			Content: model.Content{
				ID:        model.ContentID(c.ID),
				UserID:    model.UserID(c.UserID),
				ThreadID:  model.ThreadID(c.ThreadID),
				Text:      c.Text,
				CreatedAt: c.CreatedAt,
				UpdatedAt: c.UpdatedAt,
			},
			User: model.ShowUser{
				ID:        model.UserID(u.ID),
				Name:      u.Name,
				CreatedAt: u.CreatedAt,
			},
		}
		if c.Img.Valid {
			cu.Content.Img = &c.Img.String
		}
		if u.Avatar.Valid {
			cu.User.Avatar = &u.Avatar.String
		}
		cus = append(cus, cu)
	}

	return cus, tx.Commit()
}

func toContentUser(c schema.Content, u schema.User) *model.ContentUser {
	cu := &model.ContentUser{
		Content: model.Content{
			ID:        model.ContentID(c.ID),
			UserID:    model.UserID(c.UserID),
			ThreadID:  model.ThreadID(c.ThreadID),
			Text:      c.Text,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		},
		User: model.ShowUser{
			ID:        model.UserID(u.ID),
			Name:      u.Name,
			CreatedAt: u.CreatedAt,
		},
	}
	if c.Img.Valid {
		cu.Content.Img = &c.Img.String
	}
	if u.Avatar.Valid {
		cu.User.Avatar = &u.Avatar.String
	}

	return cu
}
