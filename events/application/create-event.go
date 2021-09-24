package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type CreateEventUsecase struct {
	Repository repositories.EventsRepository
}

type CreateEventParams struct {
	Name               string
	Description        string
	IsAvailable        bool
	OrganizerAccountId string
}

func (c CreateEventUsecase) Create(params *CreateEventParams) error {

	event, err := entities.NewEvent(
		params.Name,
		params.Description,
		params.OrganizerAccountId,
		params.IsAvailable,
	)
	if err != nil {
		return err
	}

	return c.Repository.Create(event)
}
