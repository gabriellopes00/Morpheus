package infra

import "github.com/gofrs/uuid"

type uuidGenerator struct{}

func NewUUIDGenerator() *uuidGenerator {
	return &uuidGenerator{}
}

func (u *uuidGenerator) Generate() (string, error) {
	result, err := uuid.NewV4()
	if err != nil {
		return "", nil
	}

	return result.String(), nil
}
