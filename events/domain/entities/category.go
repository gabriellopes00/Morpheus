package entities

import (
	"errors"
	"events/domain"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Category struct {
	domain.Entity
	Name string `json:"name,omitempty"`
}

func NewCategory(name string) (*Category, error) {
	category := new(Category)

	category.Id = uuid.NewV4().String()
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()
	category.Name = name

	err := category.validate()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Category) validate() error {
	if len(s.Name) == 0 || len(s.Name) > 255 {
		return errors.New("invalid category name")
	}

	return nil
}

// gorm required
func (Category) TableName() string {
	return "categories"
}
