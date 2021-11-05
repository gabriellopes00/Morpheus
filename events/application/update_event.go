package application

import (
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
)

type updateEventUsecase struct {
	Repository repositories.EventsRepository
}

func NewUpdateEventUsecase(repo repositories.EventsRepository) *updateEventUsecase {
	return &updateEventUsecase{
		Repository: repo,
	}
}

func (u *updateEventUsecase) UpdateStatus(eventId string, status interface{}) error {
	eventExists, err := u.Repository.ExistsId(eventId)
	if err != nil {
		return err
	}

	if !eventExists {
		return errors.New("event with given id does not exists")
	}

	stts, ok := status.(entities.EventStatus)
	if !ok {
		return nil
	}

	return u.Repository.SetStatus(eventId, stts)
}
