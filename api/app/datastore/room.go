package datastore

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
)

var _ repository.RoomRepository = (*roomDatastore)(nil)

// RoomDatastoreSet ...
var RoomDatastoreSet = wire.NewSet(
	NewRoomDatastore,
	wire.Bind(new(repository.RoomRepository), new(*roomDatastore)),
)

type roomDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewRoomDatastore(db *sql.DB, tr repository.TimeRepository) *roomDatastore {
	return &roomDatastore{db: db, tr: tr}
}

// Create ルームの作成 DB処理
func (rd *roomDatastore) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID) (*model.RoomUser, error) {
	tx, err := rd.db.Begin()
	if err != nil {
		return nil, nil
	}
	defer tx.Rollback()

	// Roomの作成
	ins, err := tx.PrepareContext(
		ctx,
		"INSERT INTO `room`(`name`, `is_group`, `created_at`, `updated_at`) VALUES(?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.ExecContext(ctx, name, isGroup, rd.tr.Now(), rd.tr.Now())
	if err != nil {
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	// UserRoomの作成
	stmt, err := tx.PrepareContext(
		ctx,
		"INSERT INTO `user_room`(`user_id`, `room_id`, `created_at`, `updated_at`) VALUES(?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var wgError error
	errChan := make(chan error, len(userIDs))

	wg.Add(len(userIDs))
	func() {
		for _, userID := range userIDs {
			go func(id model.UserID, errChan chan error) {
				defer func() {
					if err := <-errChan; err != nil {
						log.Println(err)
						wgError = err
					}
					wg.Done()
				}()

				_, err = stmt.ExecContext(ctx, id, lastID, rd.tr.Now(), rd.tr.Now())
				errChan <- err
			}(userID, errChan)
		}
	}()
	wg.Wait()
	close(errChan)
	if wgError != nil {
		return nil, wgError
	}

	var room schema.Room
	var users []*schema.User
	query := "SELECT `r`.`id`, `r`.`name`, `r`.`is_group`, `r`.`created_at`, `r`.`updated_at`, "
	query += "`u`.`id`, `u`.`name`, `u`.`avatar`, `u`.`created_at` "
	query += "FROM `room` AS `r` "
	query += "INNER JOIN ( "
	query += "SELECT `ur`.`room_id`, `u`.`id`, `u`.`name`, `u`.`avatar`, `u`.`created_at` "
	query += "FROM `user_room` AS `ur` "
	query += "INNER JOIN `user` AS `u` "
	query += "ON `ur`.`user_id` = `u`.`id` "
	query += ") AS `u` "
	query += "ON `r`.`id` = `u`.`room_id` "
	query += "WHERE `r`.`id` = ?"
	rows, err := tx.QueryContext(ctx, query, lastID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user schema.User

		if err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.IsGroup,
			&room.CreatedAt,
			&room.UpdatedAt,
			&user.ID,
			&user.Name,
			&user.Avatar,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return rd.toModelRoomUser(room, users), tx.Commit()
}

// Delete ルームの削除
func (rd *roomDatastore) Delete(ctx context.Context, id model.RoomID) error {
	stmt, err := rd.db.Prepare("DELETE FROM `room` WHERE `id` = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (rd *roomDatastore) toModelRoomUser(room schema.Room, users []*schema.User) *model.RoomUser {
	var showUsers []*model.ShowUser
	for _, u := range users {
		showUser := &model.ShowUser{
			ID:        model.UserID(u.ID),
			Name:      u.Name,
			CreatedAt: u.CreatedAt,
		}
		if u.Avatar.Valid {
			showUser.Avatar = &u.Avatar.String
		}

		showUsers = append(showUsers, showUser)
	}

	return &model.RoomUser{
		Room: model.Room{
			ID:        model.RoomID(room.ID),
			Name:      room.Name,
			IsGroup:   room.IsGroup,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		},
		Users: showUsers,
	}
}
