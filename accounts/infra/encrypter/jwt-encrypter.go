package encrypter

import (
	"accounts/config/env"
	"accounts/domain"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt"
)

type jwtEncrypter struct{}

func NewJwtEncrypter() *jwtEncrypter {
	return &jwtEncrypter{}
}

func (j *jwtEncrypter) Encrypt(payload *domain.Account) (string, error) {
	claims := jwt.MapClaims{}

	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	claims["id"] = id.String()
	claims["authorized"] = true
	claims["account_id"] = payload.Id
	claims["account_email"] = payload.Email
	claims["account_name"] = payload.Name
	claims["account_created_at"] = payload.CreatedAt
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(env.TOKEN_EXPIRATION_TIME)).Local()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(env.TOKEN_KEY))
	if err != nil {
		return "", err
	}

	return signed, nil
}
