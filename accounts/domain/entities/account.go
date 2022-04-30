package entities

import (
	app_error "accounts/domain/errors"
	"time"

	"github.com/asaskevich/govalidator"
	gouuid "github.com/satori/go.uuid"
)

// Available gender options
const (
	Male        = "male"
	Female      = "female"
	Unspecified = "unspecified"
)

// Account data model
type Account struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Document    string    `json:"document,omitempty"`
	AvatarUrl   string    `json:"avatar_url,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// New Account returns a new account model or a validation error
func NewAccount(name, email, avatarUrl, phoneNumber, gender, birthDate, document string) (*Account, error) {
	account := &Account{
		Id:          gouuid.NewV4().String(),
		Name:        name,
		Email:       email,
		Document:    document,
		PhoneNumber: phoneNumber,
		Gender:      gender,
		AvatarUrl:   avatarUrl,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
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

// validate validates account's data
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

	documentPattern := `^([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})$`
	if !govalidator.Matches(account.Document, documentPattern) {
		return app_error.NewAppError("Invalid input", "invalid document format")
	}

	// phoneNumberPattern := `^\(\d{2,}\) \d{5,}\-\d{4}$` TODO: update regexp
	phoneNumberPattern := `^\(\d{2,}\) \d{4,}\-\d{4}$`
	if !govalidator.Matches(account.PhoneNumber, phoneNumberPattern) {
		return app_error.NewAppError("Invalid input", "invalid phone number format")
	}

	if account.Gender != Male && account.Gender != Unspecified && account.Gender != Female {
		err = app_error.NewAppError("Invalid input", "gender must be either: \"male\", \"female\" or \"unspecified\"")
	}

	return err
}
