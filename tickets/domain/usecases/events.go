package usecases

import (
	"time"
)

type CreateEvent interface {
	Create(eventId string, date time.Time) error
}
