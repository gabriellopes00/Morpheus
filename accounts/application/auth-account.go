package usecases

import (
	"accounts/interfaces"
	"errors"
)

type authAccount struct {
	Repository interfaces.Repository
	Encrypter  interfaces.Encrypter
}

var (
	ErrUnregisteredEmail = errors.New("unregistered email")
)

func NewAuthAccount(repository interfaces.Repository, encrypter interfaces.Encrypter) *authAccount {
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

	token, err := a.Encrypter.Encrypt(account)
	if err != nil {
		return "", err
	}

	return token, nil
}
