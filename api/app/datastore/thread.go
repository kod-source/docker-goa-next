package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
	"github.com/shogo82148/pointer"
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

// GetThreadsByRoom ...
func (td *threadDatastore) GetThreadsByRoom(ctx context.Context, roomID model.RoomID, nextID model.ThreadID) ([]*model.IndexThread, *int, error) {
	tx, err := td.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	query := "SELECT `th`.`id`, `th`.`user_id`, `th`.`room_id`, `th`.`text`, `th`.`created_at`, `th`.`updated_at`, `th`.`img`, " +
		"`u`.`id`, `u`.`name`, `u`.`created_at`, `u`.`avatar`, `c`.`count` " +
		"FROM `thread` AS `th` " +
		"INNER JOIN `user` AS `u` " +
		"ON `th`.`user_id` = `u`.`id` " +
		"LEFT JOIN (" +
		"SELECT `thread_id`, COUNT(`thread_id`) AS `count` " +
		"FROM `content` " +
		"GROUP BY thread_id" +
		") AS `c` " +
		"ON `th`.`id` = `c`.`thread_id` " +
		"WHERE `th`.`room_id` = ? " +
		"ORDER BY `th`.`created_at` DESC " +
		"LIMIT ?, ?"
	rows, err := tx.QueryContext(ctx, query, roomID, nextID, LIMIT)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var its []*model.IndexThread
	for rows.Next() {
		var thread schema.Thread
		var user schema.User
		var count sql.NullInt64
		if err := rows.Scan(
			&thread.ID,
			&thread.UserID,
			&thread.RoomID,
			&thread.Text,
			&thread.CreatedAt,
			&thread.UpdatedAt,
			&thread.Img,
			&user.ID,
			&user.Name,
			&user.CreatedAt,
			&user.Avatar,
			&count,
		); err != nil {
			return nil, nil, err
		}

		it := &model.IndexThread{
			ThreadUser: model.ThreadUser{
				Thread: model.Thread{
					ID:        model.ThreadID(thread.ID),
					UserID:    model.UserID(thread.UserID),
					RoomID:    model.RoomID(thread.RoomID),
					Text:      thread.Text,
					CreatedAt: thread.CreatedAt,
					UpdatedAt: thread.UpdatedAt,
					Img:       nil,
				},
				User: model.ShowUser{
					ID:        model.UserID(user.ID),
					Name:      user.Name,
					CreatedAt: user.CreatedAt,
					Avatar:    nil,
				},
			},
			CountContent: nil,
		}
		if thread.Img.Valid {
			it.ThreadUser.Thread.Img = &thread.Img.String
		}
		if user.Avatar.Valid {
			it.ThreadUser.User.Avatar = &user.Avatar.String
		}
		if count.Valid {
			it.CountContent = pointer.Ptr(int(count.Int64))
		}
		its = append(its, it)
	}

	var threadID model.ThreadID
	if err := tx.QueryRowContext(
		ctx,
		"SELECT `id` FROM `thread` WHERE `room_id` = ? ORDER BY `created_at` LIMIT 1",
		roomID,
	).Scan(
		&threadID,
	); err != nil {
		return nil, nil, err
	}
	var resNextID *int
	resNextID = pointer.Int(int(nextID) + LIMIT)
	if len(its) == 0 || its[len(its)-1].ThreadUser.Thread.ID == threadID {
		resNextID = nil
	}

	return its, resNextID, tx.Commit()
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
