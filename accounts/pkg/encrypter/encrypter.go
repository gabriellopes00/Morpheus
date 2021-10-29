package encrypter

import "time"

type Token struct {
	AccessToken          string
	AccessTokenId        string
	AccessTokenDuration  time.Duration
	RefreshToken         string
	RefreshTokenId       string
	RefreshTokenDuration time.Duration
}

type Encrypter interface {
	EncryptAuthToken(accountId string) (*Token, error)
	DecryptAuthToken(token string) (string, error)
	DecryptRefreshToken(token string) (*RefreshTokenClaims, error)
}
