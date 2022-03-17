package encrypter

type Encrypter interface {
	DecryptToken(token, key string) (accountId string, err error)
}
