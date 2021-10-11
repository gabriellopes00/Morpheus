package usecases

import (
	"accounts/domain"
	"accounts/pkg/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type createAccount struct {
	Repository db.Repository
}

var (
	ErrEmailAlreadyInUse = errors.New("email already in use")
)

func NewCreateAccount(Repository db.Repository) *createAccount {
	return &createAccount{
		Repository: Repository,
	}
}

func (c *createAccount) Create(account *domain.CreateAccountDTO) (*domain.Account, error) {
	existingAccount, err := c.Repository.Exists(account.Email)
	if err != nil {
		return nil, err
	}

	if existingAccount {
		return nil, ErrEmailAlreadyInUse
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	a, err := domain.NewAccount(
		account.Name,
		account.Email,
		string(hash),
		account.AvatarUrl,
		account.RG,
		account.BirthDate)
	if err != nil {
		return nil, err
	}

	err = c.Repository.Create(a)
	if err != nil {
		return nil, err
	}

	return a, nil
}
