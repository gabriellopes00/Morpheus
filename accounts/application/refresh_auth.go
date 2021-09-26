package usecases

import (
	"accounts/domain"
	"accounts/pkg/encrypter"
)

type refreshAuth struct {
	Encrypter encrypter.Encrypter
}

func NewRefreshAuth(encrypter encrypter.Encrypter) *refreshAuth {
	return &refreshAuth{
		Encrypter: encrypter,
	}
}

func (r *refreshAuth) Refresh(refreshToken string) (domain.AuthModel, error) {
	token, err := r.Encrypter.RefreshAuthToken(refreshToken)
	if err != nil {
		return domain.AuthModel{}, err
	}

	authModel := domain.AuthModel{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return authModel, nil
}
