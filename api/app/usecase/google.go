package usecase

type GoogleUsecase interface {
	GetLoginURL(state string) string
}
