package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/schema"
)

var _ repository.UserRepository = (*userDatastore)(nil)

var UserDatastoreSet = wire.NewSet(
	NewUserDatastore,
	wire.Bind(new(repository.UserRepository), new(*userDatastore)),
)

type userDatastore struct {
	db *sql.DB
	tr repository.TimeRepository
}

func NewUserDatastore(db *sql.DB, tr repository.TimeRepository) *userDatastore {
	return &userDatastore{
		db: db,
		tr: tr,
	}
}

func (ud *userDatastore) GetUser(ctx context.Context, id model.UserID) (*model.User, error) {
	var user schema.User
	err := ud.db.QueryRow(
		"SELECT `id`, `name`, `email`, `password`, `created_at`, `avatar` FROM `user` WHERE `id` = ?", id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.Avatar,
	)
	if err != nil {
		return nil, err
	}
	return ud.convetSchemaToModelUser(&user), nil
}

func (ud *userDatastore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user schema.User
	err := ud.db.QueryRow(
		"SELECT `id`, `name`, `email`, `password`, `created_at`, `avatar` FROM `user` WHERE `email` = ?",
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.Avatar,
	)
	if err != nil {
		return nil, err
	}

	return ud.convetSchemaToModelUser(&user), nil
}

func (ud *userDatastore) CreateUser(ctx context.Context, name, email, passowrd string, avatar *string) (*model.User, error) {
	tx, err := ud.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	ins, err := tx.Prepare(
		"INSERT INTO user(`name`,`email`,`password`,`created_at`, `updated_at`, `avatar`) VALUES(?,?,?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.Exec(name, email, passowrd, ud.tr.Now(), ud.tr.Now(), avatar)
	if err != nil {
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	var user schema.User
	err = tx.QueryRow(
		"SELECT `id`, `name`, `email`, `password`, `created_at`, `avatar` FROM `user` WHERE `id` = ?", lastID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.Avatar,
	)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ud.convetSchemaToModelUser(&user), nil
}

// IndexUser ...
func (ud *userDatastore) IndexUser(ctx context.Context, myID model.UserID) ([]*model.User, error) {
	tx, err := ud.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	query := "SELECT `id`, `name`, `email`, `created_at`, `avatar` "
	query += "FROM `user` "
	query += "WHERE `id` != ? "
	query += "ORDER BY `created_at` DESC"
	rows, err := tx.QueryContext(ctx, query, myID)
	if err != nil {
		return nil, err
	}

	var users []*schema.User
	for rows.Next() {
		var u schema.User

		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.CreatedAt,
			&u.Avatar,
		); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return ud.convertSchemaToUsers(users), tx.Commit()
}

func (ud *userDatastore) convertSchemaToUsers(users []*schema.User) []*model.User {
	us := []*model.User{}

	for _, u := range users {
		us = append(us, ud.convetSchemaToModelUser(u))
	}
	return us
}

func (ud *userDatastore) convetSchemaToModelUser(user *schema.User) *model.User {
	u := &model.User{
		ID:        model.UserID(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
	if user.Avatar.Valid {
		u.Avatar = &user.Avatar.String
	}
	return u
}
