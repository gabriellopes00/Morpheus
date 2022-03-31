package entities

import (
	"errors"
	"events/domain"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Subject struct {
	domain.Entity
	Name string `json:"name,omitempty"`
}

func NewSubject(name string) (*Subject, error) {
	subject := new(Subject)

	subject.Id = uuid.NewV4().String()
	subject.CreatedAt = time.Now()
	subject.UpdatedAt = time.Now()
	subject.Name = name

	err := subject.validate()
	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (s *Subject) validate() error {
	if len(s.Name) == 0 || len(s.Name) > 255 {
		return errors.New("invalid subject name")
	}

	return nil
}

// gorm required
func (Subject) TableName() string {
	return "subjects"
}
