package application

import (
	"accounts/domain"
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"accounts/pkg/encrypter"
)

type refreshAuth struct {
	Encrypter       encrypter.Encrypter
	CacheRepository cache.CacheRepository
	Repository      db.Repository
}

func NewRefreshAuth(encrypter encrypter.Encrypter, repo db.Repository, cacheRepo cache.CacheRepository) *refreshAuth {
	return &refreshAuth{
		Encrypter:       encrypter,
		Repository:      repo,
		CacheRepository: cacheRepo,
	}
}

func (r *refreshAuth) Refresh(refreshToken string) (*domain.AuthModel, error) {
	refreshTokenClaims, err := r.Encrypter.DecryptRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	existingToken, err := r.CacheRepository.Get(refreshTokenClaims.Id)
	if err != nil {
		return nil, err
	}

	if existingToken != refreshToken {
		return nil, encrypter.ErrInvalidRefreshToken
	}

	if err = r.CacheRepository.Delete(refreshTokenClaims.Id); err != nil {
		return nil, err
	}

	existingAccount, err := r.Repository.ExistsId(refreshTokenClaims.AccountId)
	if err != nil {
		return nil, err
	}

	if !existingAccount {
		return nil, encrypter.ErrInvalidRefreshToken
	}

	token, err := r.Encrypter.EncryptAuthToken(refreshTokenClaims.AccountId)
	if err != nil {
		return nil, err
	}

	authModel := &domain.AuthModel{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return authModel, nil
}
