package usecases

import (
	"accounts/entities"
	"accounts/ports"
	"errors"
	"time"
)

type CreateAccount struct {
	UUIDGenerator ports.UUIDGenerator
	Hasher        ports.Hasher
	Repository    ports.Repository
}

func (c *CreateAccount) Create(account entities.Account) (*entities.Account, error) {
	existing, err := c.Repository.Exists(account.Email)
	if err != nil {
		return nil, err
	}

	if existing {
		return nil, errors.New("Email already in use")
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
