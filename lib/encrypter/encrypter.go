package main

type Encrypter interface {
	DecryptToken(token string) (string, error)
}
