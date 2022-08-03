package interactor

import (
	"context"
	"database/sql"

	"github.com/kod-source/docker-goa-next/app/model"
)

type UserInteractor interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email, password string) (*model.User, error)
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

func (u userInteractor) GetUserByEmail(ctx context.Context, email, password string) (*model.User, error) {
	var user model.User
	err := u.db.QueryRow(
		"SELECT `id`, `name`, `email`, `password`, `created_at` FROM `users` WHERE `email` = ? AND `password` = ?",
		email,
		password,
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
