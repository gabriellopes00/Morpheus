package usecases

import (
	"accounts/framework/encrypter"
	"accounts/interfaces"
	"errors"
)

type authAccount struct {
	Repository interfaces.Repository
	Encrypter  encrypter.Encrypter
}

var (
	ErrUnregisteredEmail = errors.New("unregistered email")
)

func NewAuthAccount(repository interfaces.Repository, encrypter encrypter.Encrypter) *authAccount {
	return &authAccount{
		Repository: repository,
		Encrypter:  encrypter,
	}
}

func (a *authAccount) Auth(email, password string) (string, error) {
	account, err := a.Repository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if account == nil {
		return "", ErrUnregisteredEmail
	}

	token, err := a.Encrypter.EncryptAuthToken(account.Id)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}
