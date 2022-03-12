package entities

import (
	app_error "accounts/domain/errors"
	"time"

	"github.com/asaskevich/govalidator"
	gouuid "github.com/satori/go.uuid"
)

// Account data model
type Account struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Document  string    `json:"document,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func NewAccount(name, email, avatarUrl, birthDate, document string) (*Account, error) {
	account := &Account{
		Id:        gouuid.NewV4().String(),
		Name:      name,
		Email:     email,
		Document:  document,
		AvatarUrl: avatarUrl,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	parsed, err := time.Parse(time.RFC3339, birthDate)
	if err != nil {
		return nil, app_error.NewAppError("Invalid input", "invalid account birth date format")
	}

	if time.Since(parsed) <= 0 {
		return nil, app_error.NewAppError("Invalid input", "invalid account birth date")
	}

	account.BirthDate = parsed

	err = account.validate()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (account *Account) validate() error {
	var err error = nil

	if len(account.Name) <= 4 || len(account.Name) > 255 {
		err = app_error.NewAppError("Invalid input", "account's name must have at least of 4 characters and at most of 255")
	}

	if !govalidator.IsURL(account.AvatarUrl) {
		err = app_error.NewAppError("Invalid input", "invalid avatar url format")
	}

	if !govalidator.IsEmail(account.Email) {
		err = app_error.NewAppError("Invalid input", "invalid email format")
	}

	documentPattern := `([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})`
	if !govalidator.Matches(account.Document, documentPattern) {
		return app_error.NewAppError("Invalid input", "invalid document format")
	}

	return err
}
