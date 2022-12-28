package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/kod-source/docker-goa-next/app/model"
	myerrors "github.com/kod-source/docker-goa-next/app/my_errors"
	"github.com/kod-source/docker-goa-next/app/repository"
	"github.com/kod-source/docker-goa-next/app/service"
	"github.com/kod-source/docker-goa-next/app/usecase"
	"golang.org/x/crypto/bcrypt"
)

var _ usecase.UserUseCase = (*userInteractor)(nil)

var UserInteractorSet = wire.NewSet(
	NewUserInteractor,
	wire.Bind(new(usecase.UserUseCase), new(*userInteractor)),
)

type userInteractor struct {
	ur repository.UserRepository
	js service.JWTService
}

func NewUserInteractor(ur repository.UserRepository, js service.JWTService) *userInteractor {
	return &userInteractor{ur: ur, js: js}
}

func (u *userInteractor) GetUser(ctx context.Context, id model.UserID) (*model.User, error) {
	user, err := u.ur.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userInteractor) GetUserByEmail(ctx context.Context, email, password string) (*model.User, error) {
	user, err := u.ur.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if err = compareHashAndPassword(user.Password, password); err != nil {
		return nil, myerrors.ErrPasswordWorng
	}

	return user, nil
}

func (u *userInteractor) CreateJWTToken(ctx context.Context, id model.UserID, name string) (*string, error) {
	token, err := u.js.CreateJWTToken(ctx, id, name)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (u *userInteractor) SignUp(ctx context.Context, name, email, password string, avatar *string) (*model.User, error) {
	p, err := passwordEncrypt(password)
	if err != nil {
		return nil, err
	}
	user, err := u.ur.CreateUser(ctx, name, email, p, avatar)
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
