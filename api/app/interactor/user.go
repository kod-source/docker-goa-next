package interactor

import (
	"context"
	"database/sql"
	"time"

	"github.com/kod-source/docker-goa-next/app/model"
)

type UserInteractor interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, name, email, password string) (*int, error)
}

type userInteractor struct {
	db *sql.DB
}

func NewUserInteractor(db *sql.DB) UserInteractor {
	return userInteractor{
		db: db,
	}
}

func (u userInteractor) GetUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	err := u.db.QueryRow(
		"SELECT `id`, `name`, `email`, `password`, `created_at` FROM `users` WHERE `id` = ?", id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userInteractor) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := u.db.QueryRow(
		"SELECT `id`, `name`, `email`, `password`, `created_at` FROM `users` WHERE `email` = ?",
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userInteractor) CreateUser(ctx context.Context, name, email, passowrd string) (*int, error) {
	ins, err := u.db.Prepare(
		"INSERT INTO users(`name`,`email`,`password`,`created_at`) VALUES(?,?,?,?)",
	)
	if err != nil {
		return nil, err
	}
	res, err := ins.Exec(name, email, passowrd, time.Now())
	if err != nil {
		return nil, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	id := int(lastId)

	return &id, nil
}
