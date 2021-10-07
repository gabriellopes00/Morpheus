package encrypter

import (
	"errors"
	"events/config/env"

	"github.com/golang-jwt/jwt"
)

type jwtEncrypter struct{}

func NewJwtEncrypter() *jwtEncrypter {
	return &jwtEncrypter{}
}

var ErrInvalidToken = errors.New("invalid authentication token")

func (jwtEncrypter) Decrypt(token string) (accountId string, err error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(env.AUTH_TOKEN_KEY), nil
	})

	if err != nil {
		return "", ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", ErrInvalidToken
	}

	if !jwtToken.Valid {
		return "", ErrInvalidToken
	}

	accountId = claims["account_id"].(string)
	if accountId == "" {
		return "", ErrInvalidToken
	}

	return accountId, nil
}
