package usecases

import (
	"accounts/domain"
	"accounts/framework/cache"
	"accounts/framework/encrypter"
	"accounts/interfaces"
	"errors"
	"time"
)

type authAccount struct {
	Repository      interfaces.Repository
	Encrypter       encrypter.Encrypter
	CacheRepository cache.CacheRepository
}

var (
	ErrUnregisteredEmail = errors.New("unregistered email")
)

func NewAuthAccount(
	repository interfaces.Repository,
	encrypter encrypter.Encrypter,
	cacheRepository cache.CacheRepository,
) *authAccount {
	return &authAccount{
		Repository:      repository,
		Encrypter:       encrypter,
		CacheRepository: cacheRepository,
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

	err = a.CacheRepository.Set(token.AccessId, token.AccessToken, time.Duration(token.AtExpires))
	if err != nil {
		return domain.AuthModel{}, err
	}

	err = a.CacheRepository.Set(token.RefreshId, token.RefreshToken, time.Duration(token.RtExpires))
	if err != nil {
		return domain.AuthModel{}, err
	}

	authModel := domain.AuthModel{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return authModel, nil
}
