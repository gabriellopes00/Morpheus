package infra

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct{}

func (b *BcryptHasher) GenHash(payload string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)
	return string(hash), err
}
