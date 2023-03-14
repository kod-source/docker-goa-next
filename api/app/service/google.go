package service

import (
	"context"

	"github.com/kod-source/docker-goa-next/app/model"
)

type GoogleService interface {
	// GetLoginURL ログインするためのリダレクトURLを取得
	GetLoginURL(state string) string
	// GetUserInfo codeからUserの情報を返す
	GetUserInfo(ctx context.Context, code string) (*model.GoogleUser, error)
}
