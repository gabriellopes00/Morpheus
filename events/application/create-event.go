package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type createEventUsecase struct {
	Repository repositories.EventsRepository
}

func NewCreateEventUsecase(repo repositories.EventsRepository) *createEventUsecase {
	return &createEventUsecase{
		Repository: repo,
	}
}

func (c *createEventUsecase) Create(params *entities.Event) (*entities.Event, error) {

	event, err := entities.NewEvent(
		params.Name,
		params.Description,
		params.OrganizerAccountId,
		params.IsAvailable,
	)
	if err != nil {
		return nil, err
	}

	if err = c.Repository.Create(event); err != nil {
		return nil, err
	}

	return event, nil
}
