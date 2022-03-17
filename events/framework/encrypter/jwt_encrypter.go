package encrypter

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

var (
	ErrInvalidToken        = errors.New("invalid authentication token")
	ErrExpiredToken        = errors.New("authentication token expired")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type jwtEncrypter struct {
}

func NewJwtEncrypter() *jwtEncrypter {
	return &jwtEncrypter{}
}

func (jwtEncrypter) DecryptToken(token, key string) (string, error) {

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrInvalidToken
		}

		// generate key using keyclaock public rsa key
		secretKey := "-----BEGIN CERTIFICATE-----\n" + key + "\n-----END CERTIFICATE-----"
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(secretKey))
		if err != nil {
			return "", err
		}

		return key, nil
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

	accountId, ok := claims["account_id"].(string)
	if !ok {
		return "", ErrInvalidToken
	}

	if accountId == "" {
		return "", ErrInvalidToken
	}

	return accountId, nil
}
