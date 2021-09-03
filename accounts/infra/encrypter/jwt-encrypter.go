package encrypter

import (
	"accounts/config/env"
	"accounts/domain"
	"errors"
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
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(env.TOKEN_EXPIRATION_TIME)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(env.TOKEN_KEY))
	if err != nil {
		return "", err
	}

	return signed, nil
}

var (
	ErrInvalidToken = errors.New("invalid authorization token")
)

func (j *jwtEncrypter) Decrypt(token string) (*domain.Account, error) {

	payload, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(env.TOKEN_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	var account domain.Account

	claims, ok := payload.Claims.(jwt.MapClaims)
	if ok && payload.Valid {

		account.Id = claims["account_id"].(string)
		account.Email = claims["account_email"].(string)
		account.Name = claims["account_name"].(string)
		// account.CreatedAt = time.Parse(time.RFC1123) claims["account_created_at"].(string)

	}

	return &account, nil
}
