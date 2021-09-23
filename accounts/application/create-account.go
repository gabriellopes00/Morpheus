package usecases

import (
	"accounts/domain"
	"accounts/interfaces"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type createAccount struct {
	Repository interfaces.Repository
}

var (
	ErrEmailAlreadyInUse = errors.New("email already in use")
)

func NewCreateAccount(Repository interfaces.Repository) *createAccount {
	return &createAccount{
		Repository: Repository,
	}
}

func (c *createAccount) Create(account domain.Account) (*domain.Account, error) {
	existingAccount, err := c.Repository.Exists(account.Email)
	if err != nil {
		return nil, err
	}

	if existingAccount {
		return nil, ErrEmailAlreadyInUse
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	account.Id = id.String()

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	account.Password = string(hash)

	account.CreatedAt = time.Now().Local()

	err = c.Repository.Create(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
