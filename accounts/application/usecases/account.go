package usecases

import (
	"accounts/application/ports"
	"accounts/config/env"
	"accounts/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type accountUsecase struct {
	Repository ports.Repository
}

func NewAccountUsecase(Repository ports.Repository) *accountUsecase {
	return &accountUsecase{
		Repository: Repository,
	}
}

func (c *accountUsecase) Create(account domain.Account) (*domain.Account, error) {
	existingAccount, err := c.Repository.Exists(account.Email)
	if err != nil {
		return nil, err
	}

	if existingAccount {
		return nil, domain.ErrEmailAlreadyInUse
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

func (u *accountUsecase) Auth(data domain.AuthCredentials) (string, error) {
	account, err := u.Repository.FindByEmail(data.Email)
	if err != nil {
		return "", err
	}

	if account == nil {
		return "", domain.ErrUnregisteredEmail
	}

	claims := jwt.MapClaims{}

	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	claims["authorized"] = true
	claims["accountId"] = account.Id
	claims["id"] = id.String()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(env.TOKEN_EXPIRATION_TIME)).Local()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(env.TOKEN_KEY))
	if err != nil {
		return "", err
	}

	return signed, nil
}
