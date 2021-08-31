package interfaces

import "accounts/domain"

type Encrypter interface {
	Encrypt(payload *domain.Account) (string, error)
	// Decrypt(token string) (*domain.Account, error)
}
