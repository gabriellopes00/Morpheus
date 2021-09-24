package repositories

import (
	"database/sql"
	"events/domain/entities"
)

type EventsRepository interface {
	Create(event *entities.Event) error
}

type pgEventsRepository struct {
	Db *sql.DB
}

func NewPgEventsRepository(connection *sql.DB) *pgEventsRepository {
	return &pgEventsRepository{connection}
}

func (repo *pgEventsRepository) Create(event *entities.Event) error {
	stm, err := repo.Db.Prepare("INSERT INTO events VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		event.Id,
		event.Name,
		event.Description,
		event.IsAvailable,
		event.OrganizerAccountId,
		event.CreatedAt,
	)

	return err
}
