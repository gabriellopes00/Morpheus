package usecases

import "accounts/domain/entities"

type UpdateAccountDTO struct {
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}

type UpdateAccount interface {
	Update(accountId string, data *UpdateAccountDTO) (*entities.Account, error)
}
