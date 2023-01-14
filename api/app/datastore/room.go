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
	"github.com/shogo82148/pointer"
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
func (rd *roomDatastore) Create(ctx context.Context, name string, isGroup bool, userIDs []model.UserID, img *string) (*model.RoomUser, error) {
	tx, err := rd.db.Begin()
	if err != nil {
		return nil, nil
	}
	defer tx.Rollback()

	// Roomの作成
	ins, err := tx.PrepareContext(
		ctx,
		"INSERT INTO `room`(`name`, `is_group`, `created_at`, `updated_at` , `img`) VALUES(?,?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.ExecContext(ctx, name, isGroup, rd.tr.Now(), rd.tr.Now(), img)
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
	query := "SELECT `r`.`id`, `r`.`name`, `r`.`is_group`, `r`.`created_at`, `r`.`updated_at`, `r`.`img`, "
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
			&room.Img,
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

// Index ルームの一覧を返す
func (rd *roomDatastore) Index(ctx context.Context, id model.UserID, nextID model.RoomID) ([]*model.IndexRoom, *int, error) {
	tx, err := rd.db.Begin()
	if err != nil {
		return nil, nil, err
	}
	defer tx.Rollback()

	query := "SELECT `thr`.`id`, `thr`.`name`, `thr`.`is_group`, `thr`.`created_at`, `thr`.`updated_at`, `thr`.`img`, "
	query += "`thr`.`last_thread_at`, `thr`.`last_text`, `thr`.`user_count`, `thr`.`show_img`, `ur`.`last_read_at` "
	query += "FROM `user_room` AS `ur` "
	query += "INNER JOIN ( "
	query += "SELECT `r`.`id`, `r`.`name`, `r`.`is_group`, `r`.`created_at`, `r`.`updated_at`, `r`.`img`, "
	query += "`th`.`created_at` AS `last_thread_at`, `th`.`text` AS `last_text`, `ur`.`user_count`, `u`.`avatar` AS `show_img` "
	query += "FROM `room` AS `r` "
	query += "INNER JOIN ( "
	query += "SELECT `room_id`, COUNT(`id`) AS `user_count` "
	query += "FROM `user_room` "
	query += "GROUP BY `room_id` "
	query += ") AS `ur` "
	query += "ON `r`.`id` = `ur`.`room_id` "
	query += "LEFT JOIN ( "
	query += "SELECT `th1`.`id`, `th1`.`room_id`, `th1`.`created_at`, `th1`.`text` "
	query += "FROM `thread` AS `th1` "
	query += "INNER JOIN ( "
	query += "SELECT `room_id`, MAX(`created_at`) AS `created_at` "
	query += "FROM `thread` "
	query += "GROUP BY `room_id` "
	query += ") AS `th2` "
	query += "ON `th1`.`room_id` = `th2`.`room_id` AND `th1`.`created_at` = `th2`.`created_at` "
	query += ") AS `th` "
	query += "ON `r`.`id` = `th`.`room_id` "
	query += "LEFT JOIN `user_room` AS `ur2` "
	query += "ON `ur2`.`room_id` = `r`.`id` AND `r`.`is_group` = 0 AND `ur2`.`user_id` != ? "
	query += "LEFT JOIN `user` AS `u` "
	query += "ON `u`.`id` = `ur2`.`user_id` "
	query += ") AS `thr` "
	query += "ON `ur`.`room_id` = `thr`.`id` "
	query += "WHERE `ur`.`user_id` = ? "
	query += "ORDER BY `thr`.`last_thread_at` DESC "
	query += "LIMIT ?, ?"
	rows, err := tx.QueryContext(ctx, query, id, id, nextID, LIMIT)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var irs []*model.IndexRoom
	for rows.Next() {
		var room schema.Room
		var userRoom schema.UserRoom
		var lastThreadAt sql.NullTime
		var lastText sql.NullString
		var userCount int
		var showImg sql.NullString

		if err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.IsGroup,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.Img,
			&lastThreadAt,
			&lastText,
			&userCount,
			&showImg,
			&userRoom.LastReadAt,
		); err != nil {
			return nil, nil, err
		}

		isOpen := true
		if lastThreadAt.Valid {
			if userRoom.LastReadAt.Valid {
				isOpen = userRoom.LastReadAt.Time.After(lastThreadAt.Time)
			} else {
				isOpen = false
			}
		}

		var lt *string
		if lastText.Valid {
			lt = &lastText.String
		}
		r := model.Room{
			ID:        model.RoomID(room.ID),
			Name:      room.Name,
			IsGroup:   room.IsGroup,
			CreatedAt: room.CreatedAt,
			UpdatedAt: room.UpdatedAt,
		}
		if room.Img.Valid {
			r.Img = &room.Img.String
		}

		var si *string
		if showImg.Valid {
			si = &showImg.String
		}

		irs = append(irs, &model.IndexRoom{
			Room:      r,
			IsOpen:    isOpen,
			LastText:  lt,
			CountUser: userCount,
			ShowImg:   si,
		})
	}

	var roomID uint64
	getLastRoomIDQuery := "SELECT `thr`.`id` "
	getLastRoomIDQuery += "FROM `user_room` AS `ur` "
	getLastRoomIDQuery += "INNER JOIN ( "
	getLastRoomIDQuery += "SELECT `r`.`id`, `th`.`created_at` AS `last_thread_at` "
	getLastRoomIDQuery += "FROM `room` AS `r` "
	getLastRoomIDQuery += "LEFT JOIN ( "
	getLastRoomIDQuery += "SELECT `th1`.`id`, `th1`.`room_id`, `th1`.`created_at` "
	getLastRoomIDQuery += "FROM `thread` AS `th1` "
	getLastRoomIDQuery += "INNER JOIN ( "
	getLastRoomIDQuery += "SELECT `room_id`, MAX(`created_at`) AS `created_at` "
	getLastRoomIDQuery += "FROM `thread` "
	getLastRoomIDQuery += "GROUP BY `room_id` "
	getLastRoomIDQuery += ") AS `th2` "
	getLastRoomIDQuery += "ON `th1`.`room_id` = `th2`.`room_id` AND `th1`.`created_at` = `th2`.`created_at` "
	getLastRoomIDQuery += ") AS `th` "
	getLastRoomIDQuery += "ON `r`.`id` = `th`.`room_id` "
	getLastRoomIDQuery += ") AS `thr` "
	getLastRoomIDQuery += "ON `ur`.`room_id` = `thr`.`id` "
	getLastRoomIDQuery += "WHERE `ur`.`user_id` = ? "
	getLastRoomIDQuery += "ORDER BY `thr`.`last_thread_at` "
	getLastRoomIDQuery += "LIMIT 1"
	if err := tx.QueryRowContext(ctx, getLastRoomIDQuery, id).Scan(
		&roomID,
	); err != nil {
		return nil, nil, err
	}

	var resNextID *int
	resNextID = pointer.Int(int(nextID) + LIMIT)
	if len(irs) == 0 || irs[len(irs)-1].Room.ID == model.RoomID(roomID) {
		resNextID = nil
	}
	return irs, resNextID, tx.Commit()
}

// GetNoneGroup 指定したUserのDMのルームを取得する
func (rd *roomDatastore) GetNoneGroup(ctx context.Context, myID model.UserID, id model.UserID) (*model.Room, error) {
	tx, err := rd.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var room schema.Room
	query := "SELECT `r`.`id`, `r`.`name`, `r`.`is_group`, `r`.`created_at`, `r`.`updated_at`, `r`.`img` "
	query += "FROM `room` AS `r` "
	query += "INNER JOIN `user_room` AS `ur1` "
	query += "ON `r`.`id` = `ur1`.`room_id` "
	query += "INNER JOIN `user_room` AS `ur2` "
	query += "ON `ur1`.`room_id` = `ur2`.`room_id` "
	query += "WHERE `r`.`is_group` = 0 AND `ur1`.`user_id` = ? AND `ur2`.`user_id` = ? "
	if err := tx.QueryRowContext(ctx, query, myID, id).Scan(
		&room.ID,
		&room.Name,
		&room.IsGroup,
		&room.CreatedAt,
		&room.UpdatedAt,
		&room.Img,
	); err != nil {
		return nil, err
	}

	r := &model.Room{
		ID:        model.RoomID(room.ID),
		Name:      room.Name,
		IsGroup:   room.IsGroup,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
	if room.Img.Valid {
		r.Img = &room.Img.String
	}
	return r, tx.Commit()
}

// Show ...
func (rd *roomDatastore) Show(ctx context.Context, id model.RoomID) (*model.RoomUser, error) {
	tx, err := rd.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "SELECT `r`.`id`, `r`.`name`, `r`.`is_group`, `r`.`created_at`, `r`.`updated_at`, `r`.`img`, "
	query += "`u`.`id`, `u`.`name`, `u`.`created_at`, `u`.`avatar` "
	query += "FROM `room` AS `r` "
	query += "INNER JOIN `user_room` AS `ur` "
	query += "ON `r`.`id` = `ur`.`room_id` "
	query += "INNER JOIN `user` AS `u` "
	query += "ON `ur`.`user_id` = `u`.`id` "
	query += "WHERE `r`.`id` = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var room schema.Room
	var users []*schema.User
	for rows.Next() {
		var user schema.User
		if err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.IsGroup,
			&room.CreatedAt,
			&room.UpdatedAt,
			&room.Img,
			&user.ID,
			&user.Name,
			&user.CreatedAt,
			&user.Avatar,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if room.ID == 0 {
		return nil, sql.ErrNoRows
	}

	return rd.toModelRoomUser(room, users), tx.Commit()
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
	r := model.Room{
		ID:        model.RoomID(room.ID),
		Name:      room.Name,
		IsGroup:   room.IsGroup,
		CreatedAt: room.CreatedAt,
		UpdatedAt: room.UpdatedAt,
	}
	if room.Img.Valid {
		r.Img = &room.Img.String
	}

	return &model.RoomUser{
		Room:  r,
		Users: showUsers,
	}
}
