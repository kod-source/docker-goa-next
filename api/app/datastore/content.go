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
