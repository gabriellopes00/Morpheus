package encrypter

import (
	"accounts/config/env"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	gouuid "github.com/satori/go.uuid"
)

var (
	ErrInvalidToken        = errors.New("invalid authentication token")
	ErrExpiredToken        = errors.New("authentication token expired")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type encrypter struct {
}

func NewEncrypter() *encrypter {
	return &encrypter{}
}

func (e *encrypter) EncryptAuthToken(accountId string) (*Token, error) {
	token := &Token{}
	var err error

	// Auth Token
	accessTokenExpTime := time.Now().Add(time.Minute * time.Duration(env.AUTH_TOKEN_EXPIRATION_TIME))
	token.AccessTokenDuration = time.Until(accessTokenExpTime)
	token.AccessTokenId = gouuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["id"] = token.AccessTokenId
	atClaims["authorized"] = true
	atClaims["account_id"] = accountId
	atClaims["exp"] = accessTokenExpTime.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(env.AUTH_TOKEN_KEY))
	if err != nil {
		return nil, err
	}

	// Refresh Token
	refreshTokenExpTime := time.Now().Add(time.Hour * 24 * 7)
	token.RefreshTokenDuration = time.Until(refreshTokenExpTime)
	token.RefreshTokenId = gouuid.NewV4().String()

	rtClaims := jwt.MapClaims{}
	rtClaims["id"] = token.RefreshTokenId
	rtClaims["account_id"] = accountId
	rtClaims["exp"] = refreshTokenExpTime.Unix()

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(env.REFRESH_TOKEN_KEY))
	if err != nil {
		return nil, err
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

	accountId, ok := claims["account_id"].(string)
	if !ok {
		return "", ErrInvalidToken
	}

	if accountId == "" {
		return "", ErrInvalidToken
	}

	return accountId, nil
}

type RefreshTokenClaims struct {
	Id        string
	AccountId string
}

func (encrypter) DecryptRefreshToken(token string) (*RefreshTokenClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(env.REFRESH_TOKEN_KEY), nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	accountId, ok := claims["account_id"].(string)
	if !ok {
		return nil, ErrInvalidRefreshToken
	}
	if accountId == "" {
		return nil, ErrInvalidToken
	}

	id, ok := claims["id"].(string)
	if !ok {
		return nil, ErrInvalidRefreshToken
	}
	if id == "" {
		return nil, ErrInvalidToken
	}

	return &RefreshTokenClaims{Id: id, AccountId: accountId}, nil
}
