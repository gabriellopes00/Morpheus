package usecases

import (
	"accounts/application/ports"
	"accounts/domain"
	"time"
)

type createAccount struct {
	UUIDGenerator ports.UUIDGenerator
	Hasher        ports.Hasher
	Repository    ports.Repository
}

func NewCreateAccount(
	UUIDGenerator ports.UUIDGenerator,
	Hasher ports.Hasher,
	Repository ports.Repository,
) *createAccount {

	return &createAccount{
		UUIDGenerator: UUIDGenerator,
		Hasher:        Hasher,
		Repository:    Repository,
	}
}

func (c *createAccount) Create(account domain.Account) (*domain.Account, error) {
	existingAccount, err := c.Repository.Exists(account.Email)
	if err != nil {
		return nil, err
	}

	if existingAccount {
		return nil, domain.ErrEmailAlreadyInUse
	}

	uuid, err := c.UUIDGenerator.Generate()
	if err != nil {
		return nil, err
	}

	account.Id = uuid

	hash, err := c.Hasher.GenHash(account.Password)
	if err != nil {
		return nil, err
	}

	account.Password = hash

	account.CreatedAt = time.Now().Local()

	err = c.Repository.Create(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
