package encrypter

type Encrypter interface {
	DecryptToken(token string) (string, error)
}
