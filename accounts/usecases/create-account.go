package usecases

import (
	"accounts/entities"
	"accounts/ports"
	"errors"
	"time"
)

var (
	ErrEmailAlerayInUse = errors.New("Email already in use")
)

type createAccount struct {
	UUIDGenerator ports.UUIDGenerator
	Hasher        ports.Hasher
	Repository    ports.Repository
}

func NewCreateAccount(UUIDGenerator ports.UUIDGenerator, Hasher ports.Hasher, Repository ports.Repository) *createAccount {
	return &createAccount{
		UUIDGenerator: UUIDGenerator,
		Hasher:        Hasher,
		Repository:    Repository,
	}
}

func (c *createAccount) Create(account entities.Account) (*entities.Account, error) {
	existing, err := c.Repository.Exists(account.Email)
	if err != nil {
		return nil, err
	}

	if existing {
		return nil, ErrEmailAlerayInUse
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
