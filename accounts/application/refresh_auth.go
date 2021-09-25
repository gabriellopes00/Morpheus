package usecases

import (
	"accounts/framework/encrypter"
	"accounts/interfaces"
)

type refreshAuth struct {
	Encrypter encrypter.Encrypter
}

func NewRefreshAuth(repository interfaces.Repository, encrypter encrypter.Encrypter) *refreshAuth {
	return &refreshAuth{
		Encrypter: encrypter,
	}
}

func (a *refreshAuth) RefreshAuth(refreshToken string) (string, error) {
	token, err := a.Encrypter.RefreshAuthToken(refreshToken)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
