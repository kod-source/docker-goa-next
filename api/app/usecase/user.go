package usecase

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/interactor"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	GetUser(ctx context.Context, id int) (*model.User, error)
	GetUserByEmail(ctx context.Context, email, password string) (*model.User, error)
	CreateJWTToken(ctx context.Context, id int, name string) (*string, error)
	SignUp(ctx context.Context, name, email, password string) (*model.User, error)
}

type userUseCase struct {
	ui interactor.UserInteractor
	ji interactor.JWTInteractor
}

func NewUserUseCase(ui interactor.UserInteractor, ji interactor.JWTInteractor) UserUseCase {
	return userUseCase{ui: ui, ji: ji}
}

func (u userUseCase) GetUser(ctx context.Context, id int) (*model.User, error) {
	user, err := u.ui.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUseCase) GetUserByEmail(ctx context.Context, email, password string) (*model.User, error) {
	user, err := u.ui.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if err = compareHashAndPassword(user.Password, password); err != nil {
		return nil, myerrors.PasswordWorngError
	}

	return user, nil
}

func (u userUseCase) CreateJWTToken(ctx context.Context, id int, name string) (*string, error) {
	token, nil := u.ji.CreateJWTToken(ctx, id, name)
	return token, nil
}

func (u userUseCase) SignUp(ctx context.Context, name, email, password string) (*model.User, error) {
	p, err := passwordEncrypt(password)
	if err != nil {
		return nil, err
	}
	id, err := u.ui.CreateUser(ctx, name, email, p)
	if err != nil {
		return nil, err
	}
	user, err := u.ui.GetUser(ctx, *id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 暗号(Hash)化
func passwordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func compareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
