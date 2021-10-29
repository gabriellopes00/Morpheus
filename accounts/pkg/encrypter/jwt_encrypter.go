package encrypter

import (
	"accounts/config/env"
	"accounts/pkg/cache"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	gouuid "github.com/satori/go.uuid"
)

var (
	ErrInvalidToken        = errors.New("invalid authentication token")
	ErrExpiredToken        = errors.New("authentication token expired")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type encrypter struct {
	cache cache.CacheRepository
}

func NewEncrypter(cacheRepository cache.CacheRepository) *encrypter {
	return &encrypter{cache: cacheRepository}
}

func (e *encrypter) EncryptAuthToken(accountId string) (Token, error) {
	token := Token{}
	var err error

	atDuration := time.Now().Add(time.Minute * time.Duration(env.AUTH_TOKEN_EXPIRATION_TIME))
	token.AtExpires = atDuration.Unix()
	token.AccessId = gouuid.NewV4().String()

	rtDuration := time.Now().Add(time.Hour * 24 * 7)
	token.RtExpires = rtDuration.Unix()
	token.RefreshId = gouuid.NewV4().String()

	// Auth Token
	atClaims := jwt.MapClaims{}
	atClaims["id"] = token.AccessId
	atClaims["authorized"] = true
	atClaims["account_id"] = accountId
	atClaims["exp"] = token.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(env.AUTH_TOKEN_KEY))
	if err != nil {
		return Token{}, err
	}

	err = e.cache.Set(token.AccessId, token.AccessToken, time.Until(atDuration))
	if err != nil {
		return Token{}, err
	}

	// Auth Token
	rtClaims := jwt.MapClaims{}
	rtClaims["id"] = token.RefreshId
	rtClaims["account_id"] = accountId
	rtClaims["exp"] = token.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(env.REFRESH_TOKEN_KEY))
	if err != nil {
		return Token{}, err
	}

	err = e.cache.Set(token.RefreshId, token.RefreshToken, time.Until(rtDuration))
	if err != nil {
		return Token{}, err
	}

	return token, nil
}

func (e *encrypter) RefreshAuthToken(refreshToken string) (Token, error) {
	token := Token{}
	var err error

	jwtToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(env.REFRESH_TOKEN_KEY), nil
	})

	if err != nil {
		return Token{}, ErrInvalidToken
	}

	if !jwtToken.Valid {
		return Token{}, ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return Token{}, ErrInvalidToken
	}

	refreshTokenId, ok := claims["id"].(string)
	if !ok {
		return Token{}, ErrInvalidToken
	}

	exitingToken, err := e.cache.Get(refreshTokenId)
	if err != nil && !errors.Is(err, redis.Nil) {
		return Token{}, err
	}

	if exitingToken != refreshToken {
		return Token{}, ErrInvalidRefreshToken
	}

	if err = e.cache.Delete(refreshTokenId); err != nil {
		return Token{}, err
	}

	accountId, ok := claims["account_id"].(string)
	if !ok {
		return Token{}, ErrInvalidToken
	}

	token, err = e.EncryptAuthToken(accountId) // TODO: cannot call encrypt token without find this account id in db
	if err != nil {
		return Token{}, nil
	}

	return token, nil

}

func (encrypter) DecryptAuthToken(token string) (string, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(env.AUTH_TOKEN_KEY), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	if !jwtToken.Valid {
		return "", ErrInvalidToken
	}

	accountId := claims["account_id"].(string)
	if accountId == "" {
		return "", ErrInvalidToken
	}

	return accountId, nil
}
