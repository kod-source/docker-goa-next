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

var _ repository.ThreadRepository = (*threadDatastore)(nil)

// ThreadRepositorySet ...
var ThreadRepositorySet = wire.NewSet(
	NewThreadRepository,
	wire.Bind(new(repository.ThreadRepository), new(*threadDatastore)),
)

type threadDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewThreadRepository(db *sql.DB, tr repository.TimeRepository) *threadDatastore {
	return &threadDatastore{db: db, tr: tr}
}

// Create ...
func (td *threadDatastore) Create(ctx context.Context, text string, roomID model.RoomID, userID model.UserID, img *string) (*model.ThreadUser, error) {
	tx, err := td.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		ctx,
		"INSERT INTO `thread`(`room_id`, `user_id`, `text`, `created_at`, `updated_at`, `img`) VALUES(?,?,?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, roomID, userID, text, td.tr.Now(), td.tr.Now(), img)
	if err != nil {
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var thread schema.Thread
	var user schema.User
	query := "SELECT `th`.`id`, `th`.`room_id`, `th`.`user_id`, `th`.`text`, `th`.`created_at`, `th`.`updated_at`, `th`.`img`, " +
		"`u`.`id`, `u`.`name`, `u`.`created_at`, `u`.`avatar` " +
		"FROM `thread` AS `th` " +
		"INNER JOIN `user` AS `u` " +
		"ON `th`.`user_id` = `u`.`id` " +
		"WHERE `th`.`id` = ?"
	if err := tx.QueryRowContext(ctx, query, lastID).Scan(
		&thread.ID,
		&thread.RoomID,
		&thread.UserID,
		&thread.Text,
		&thread.CreatedAt,
		&thread.UpdatedAt,
		&thread.Img,
		&user.ID,
		&user.Name,
		&user.CreatedAt,
		&user.Avatar,
	); err != nil {
		return nil, err
	}

	return toModelThreadUser(thread, user), tx.Commit()
}

// Delete ...
func (td *threadDatastore) Delete(ctx context.Context, myID model.UserID, threadID model.ThreadID) error {
	tx, err := td.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var user schema.User
	if err := tx.QueryRowContext(
		ctx,
		"SELECT `user_id` FROM `thread` WHERE `id` = ?",
		threadID,
	).Scan(
		&user.ID,
	); err != nil {
		return err
	}
	if user.ID != uint64(myID) {
		return myerrors.ErrBadRequestNoPermission
	}

	stmt, err := tx.PrepareContext(ctx, "DELETE FROM `thread` WHERE `id` = ?")
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, threadID)
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

func toModelThreadUser(th schema.Thread, u schema.User) *model.ThreadUser {
	thu := &model.ThreadUser{
		Thread: model.Thread{
			ID:        model.ThreadID(th.ID),
			UserID:    model.UserID(th.UserID),
			RoomID:    model.RoomID(th.RoomID),
			Text:      th.Text,
			CreatedAt: th.CreatedAt,
			UpdatedAt: th.UpdatedAt,
		},
		User: model.ShowUser{
			ID:        model.UserID(u.ID),
			Name:      u.Name,
			CreatedAt: u.CreatedAt,
		},
	}
	if th.Img.Valid {
		thu.Thread.Img = &th.Img.String
	}
	if u.Avatar.Valid {
		thu.User.Avatar = &u.Avatar.String
	}

	return thu
}
