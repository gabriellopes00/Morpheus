package infra

import "github.com/gofrs/uuid"

type UUIDGenerator struct{}

func (u *UUIDGenerator) Generate() (string, error) {
	result, err := uuid.NewV4()
	if err != nil {
		return "", nil
	}

	return result.String(), nil
}
