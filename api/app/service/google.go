package service

type GoogleService interface {
	// GetLoginURL ログインするためのリダレクトURLを取得
	GetLoginURL(state string) (string)
}
