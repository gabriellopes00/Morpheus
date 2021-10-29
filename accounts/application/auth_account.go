package usecases

import (
	"accounts/domain"
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"accounts/pkg/encrypter"
	"errors"
)

type authAccount struct {
	Repository      db.Repository
	Encrypter       encrypter.Encrypter
	CacheRepository cache.CacheRepository
}

var (
	ErrUnregisteredEmail = errors.New("unregistered email")
)

func NewAuthAccount(
	repository db.Repository,
	encrypter encrypter.Encrypter,
	cacheRepository cache.CacheRepository,
) *authAccount {
	return &authAccount{
		Repository:      repository,
		Encrypter:       encrypter,
		CacheRepository: cacheRepository,
	}
}

func (a *authAccount) Auth(email, password string) (*domain.AuthModel, error) {
	account, err := a.Repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, ErrUnregisteredEmail
	}

	token, err := a.Encrypter.EncryptAuthToken(account.Id)
	if err != nil {
		return nil, err
	}

	err = a.CacheRepository.Set(token.AccessTokenId, token.AccessToken, token.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	err = a.CacheRepository.Set(token.RefreshTokenId, token.RefreshToken, token.RefreshTokenDuration)
	if err != nil {
		return nil, err
	}

	authModel := &domain.AuthModel{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return authModel, nil
}
