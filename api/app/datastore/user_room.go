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

var _ (repository.UserRoomRepository) = (*userRoomDatastore)(nil)

var UserRoomRepositorySet = wire.NewSet(
	NewUserRoomRepository,
	wire.Bind(new(repository.UserRoomRepository), new(*userRoomDatastore)),
)

type userRoomDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewUserRoomRepository(db *sql.DB, tr repository.TimeRepository) *userRoomDatastore {
	return &userRoomDatastore{db: db, tr: tr}
}

func (urd *userRoomDatastore) Create(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
	tx, err := urd.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(
		ctx,
		"INSERT INTO `user_room`(`user_id`, `room_id`, `created_at`, `updated_at`) VALUES(?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, userID, roomID, urd.tr.Now(), urd.tr.Now())
	if err != nil {
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var userRoom schema.UserRoom
	var room schema.Room
	var userCount int
	query := "SELECT `ur`.`id`, `ur`.`user_id`, `ur`.`room_id`, `ur`.`last_read_at`, `ur`.`created_at`, `ur`.`updated_at`, "
	query += "`r`.`is_group`, `cur`.`user_count` "
	query += "FROM `user_room` AS `ur` "
	query += "INNER JOIN `room` AS `r` "
	query += "ON `ur`.`room_id` = `r`.`id` "
	query += "INNER JOIN ( "
	query += "SELECT `room_id`, COUNT(`id`) AS `user_count` "
	query += "FROM `user_room` "
	query += "GROUP BY `room_id` "
	query += ") AS `cur` "
	query += "ON `ur`.`room_id` = `cur`.`room_id` "
	query += "WHERE `ur`.`id` = ?"
	if err := tx.QueryRowContext(
		ctx,
		query,
		lastID,
	).Scan(
		&userRoom.ID,
		&userRoom.UserID,
		&userRoom.RoomID,
		&userRoom.LastReadAt,
		&userRoom.CreatedAt,
		&userRoom.UpdatedAt,
		&room.IsGroup,
		&userCount,
	); err != nil {
		return nil, err
	}

	if !room.IsGroup && userCount > DMMaxCount {
		return nil, myerrors.ErrBadRequestNoPermission
	}

	return urd.toModelUserRoom(userRoom), tx.Commit()
}

func (urd *userRoomDatastore) Delete(ctx context.Context, id model.UserRoomID) error {
	stmt, err := urd.db.PrepareContext(ctx, "DELETE FROM `user_room` WHERE `id` = ?")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (urd *userRoomDatastore) toModelUserRoom(sur schema.UserRoom) *model.UserRoom {
	ur := &model.UserRoom{
		ID:        model.UserRoomID(sur.ID),
		UserID:    model.UserID(sur.UserID),
		RoomID:    model.RoomID(sur.RoomID),
		CreatedAt: sur.CreatedAt,
		UpdatedAt: sur.UpdatedAt,
	}
	if sur.LastReadAt.Valid {
		ur.LastReadAt = &sur.LastReadAt.Time
	}

	return ur
}
