package encrypter

type Encrypter interface {
	Decrypt(token string) (accountId string, err error)
}
