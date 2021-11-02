package entities

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	gouuid "github.com/satori/go.uuid"
)

// Account data model
type Account struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewAccount(name, email, password, avatarUrl, birthDate string) (*Account, error) {
	account := &Account{
		Id:        gouuid.NewV4().String(),
		Name:      name,
		Email:     email,
		Password:  password,
		AvatarUrl: avatarUrl,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	parsed, err := time.Parse(time.RFC3339, birthDate)
	if err != nil {
		return nil, errors.New("invalid account birth date format")
	}

	if time.Since(parsed) <= 0 {
		return nil, errors.New("invalid birth date")
	}

	if time.Now().Year()-parsed.Year() < 12 {
		return nil, errors.New("minimun age required to create an account is 12y")
	}

	account.BirthDate = parsed

	err = account.validate()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (account *Account) validate() error {
	var err error

	if len(account.Name) <= 4 || len(account.Name) > 255 {
		err = errors.New("account's name must have at least of 4 characters and at most of 255")
	}

	if len(account.Password) <= 4 || len(account.Password) > 255 {
		err = errors.New("account's password must have at least of 4 characters and at most of 255")
	}

	if !govalidator.IsURL(account.AvatarUrl) {
		err = errors.New("invalid avatar url format")
	}

	if !govalidator.IsEmail(account.Email) {
		err = errors.New("invalid email format")
	}

	return err
}
