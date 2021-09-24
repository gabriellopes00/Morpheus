package encrypter

import (
	"errors"
	"events/config/env"
	"events/domain/entities"

	"github.com/golang-jwt/jwt"
)

type Encrypter interface {
	Decrypt(token string) (*entities.Account, error)
}

type jwtEncrypter struct{}

func NewJwtEncrypter() *jwtEncrypter {
	return &jwtEncrypter{}
}

var (
	ErrInvalidToken = errors.New("invalid authorization token")
)

func (j *jwtEncrypter) Decrypt(token string) (*entities.Account, error) {

	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(env.TOKEN_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	var account entities.Account

	claims, ok := payload.Claims.(jwt.MapClaims)
	if ok && payload.Valid {
		account.Id = claims["account_id"].(string)
	}

	return &account, nil
}
