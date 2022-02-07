package usecases

import "accounts/domain/entities"

type CreateAccountDTO struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Document  string `json:"document,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}

type CreateAccount interface {
	Create(data *CreateAccountDTO) (*entities.Account, error)
}
