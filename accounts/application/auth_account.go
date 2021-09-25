package usecases

import (
	"accounts/domain"
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

func NewAuthAccount(
	repository interfaces.Repository,
	encrypter encrypter.Encrypter,
) *authAccount {
	return &authAccount{
		Repository: repository,
		Encrypter:  encrypter,
	}
}

func (a *authAccount) Auth(email, password string) (domain.AuthModel, error) {
	account, err := a.Repository.FindByEmail(email)
	if err != nil {
		return domain.AuthModel{}, err
	}

	if account == nil {
		return domain.AuthModel{}, ErrUnregisteredEmail
	}

	token, err := a.Encrypter.EncryptAuthToken(account.Id)
	if err != nil {
		return domain.AuthModel{}, err
	}

	authModel := domain.AuthModel{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return authModel, nil
}
