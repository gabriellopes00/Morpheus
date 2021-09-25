package usecases

import (
	"accounts/domain"
	"accounts/framework/cache"
	"accounts/framework/encrypter"
	"time"
)

type refreshAuth struct {
	Encrypter       encrypter.Encrypter
	CacheRepository cache.CacheRepository
}

func NewRefreshAuth(encrypter encrypter.Encrypter, cacheRepository cache.CacheRepository) *refreshAuth {
	return &refreshAuth{
		Encrypter:       encrypter,
		CacheRepository: cacheRepository,
	}
}

func (r *refreshAuth) Refresh(refreshToken string) (domain.AuthModel, error) {
	token, err := r.Encrypter.RefreshAuthToken(refreshToken)
	if err != nil {
		return domain.AuthModel{}, err
	}

	err = r.CacheRepository.Set(token.AccessId, token.AccessToken, time.Duration(token.AtExpires))
	if err != nil {
		return domain.AuthModel{}, nil
	}

	err = r.CacheRepository.Set(token.RefreshId, token.RefreshToken, time.Duration(token.RtExpires))
	if err != nil {
		return domain.AuthModel{}, nil
	}

	authModel := domain.AuthModel{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return authModel, nil
}
