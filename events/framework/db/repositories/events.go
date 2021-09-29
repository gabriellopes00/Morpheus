package repositories

import (
	"database/sql"
	"events/domain/entities"
)

type EventsRepository interface {
	Create(event *entities.Event) error
	GetAccountEvents(accountId string) ([]*entities.Event, error)
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

func (repo *pgEventsRepository) GetAccountEvents(accountId string) ([]*entities.Event, error) {
	stm, err := repo.Db.Prepare("SELECT * FROM events WHERE organizer_account_id = $1")
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	events := []*entities.Event{}

	rows, err := stm.Query(accountId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		event := entities.Event{}
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.IsAvailable,
			&event.OrganizerAccountId,
			&event.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}
