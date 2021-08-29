package crypto

import "golang.org/x/crypto/bcrypt"

type bcryptHasher struct{}

func NewBcryptHasher() *bcryptHasher {
	return &bcryptHasher{}
}

func (b *bcryptHasher) GenHash(payload string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(payload), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
