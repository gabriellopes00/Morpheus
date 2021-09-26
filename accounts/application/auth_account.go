package usecases

import (
	"accounts/domain"
	"accounts/pkg/db"
	"accounts/pkg/encrypter"
	"errors"
)

type authAccount struct {
	Repository db.Repository
	Encrypter  encrypter.Encrypter
}

var (
	ErrUnregisteredEmail = errors.New("unregistered email")
)

func NewAuthAccount(
	repository db.Repository,
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
