package encrypter

type Token struct {
	AccessToken  string
	RefreshToken string
	AccessId     string
	RefreshId    string
	AtExpires    int64
	RtExpires    int64
}

type Encrypter interface {
	EncryptAuthToken(accountId string) (Token, error)
	RefreshAuthToken(refreshToken string) (Token, error)
	DecryptAuthToken(token string) (string, error)
}
