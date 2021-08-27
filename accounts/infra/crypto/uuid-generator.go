package crypto

import "github.com/gofrs/uuid"

type uuidGenerator struct{}

func NewUUIDGenerator() *uuidGenerator {
	return &uuidGenerator{}
}

func (u *uuidGenerator) Generate() (string, error) {
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
