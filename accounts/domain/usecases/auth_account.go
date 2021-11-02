package usecases

type AuthModel struct {
	AccessToken  string
	RefreshToken string
}

type AuthAccount interface {
	Auth(email, password string) (*AuthModel, error)
}

type RefreshAuth interface {
	Refresh(refreshToken string) (*AuthModel, error)
}
