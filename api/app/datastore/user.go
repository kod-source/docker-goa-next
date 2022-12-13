package datastore

import (
	"context"
	"database/sql"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	"github.com/kod-source/docker-goa-next/app/repository"
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

func (u *userDatastore) GetUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := u.db.QueryRow(
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
	return &user, nil
}

func (u *userDatastore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.QueryRow(
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

	return &user, nil
}

func (u *userDatastore) CreateUser(ctx context.Context, name, email, passowrd string, avatar *string) (*model.User, error) {
	tx, err := u.db.Begin()
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
	res, err := ins.Exec(name, email, passowrd, u.tr.Now(), u.tr.Now(), avatar)
	if err != nil {
		return nil, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	var user model.User
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

	return &user, nil
}
