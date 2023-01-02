package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
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
	res, err := stmt.ExecContext(ctx, userID, roomID, urd.tr.Now(), urd.tr.Now())
	if err != nil {
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	var userRoom schema.UserRoom
	if err := tx.QueryRowContext(
		ctx,
		"SELECT `id`, `user_id`, `room_id`, `last_read_at`, `created_at`, `updated_at` FROM `user_room` WHERE `id` = ?",
		lastID,
	).Scan(
		&userRoom.ID,
		&userRoom.UserID,
		&userRoom.RoomID,
		&userRoom.LastReadAt,
		&userRoom.CreatedAt,
		&userRoom.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return urd.toModelUserRoom(userRoom), tx.Commit()
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
