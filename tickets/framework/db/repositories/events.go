package repositories

import (
	"database/sql"
	"errors"
	"tickets/domain/entities"
)

type EventsRepository interface {
	FindById(eventId string) (*entities.Event, error)
	Create(event *entities.Event) error
}

type pgEventsRepository struct {
	db *sql.DB
}

func NewPgEventsRepository(db *sql.DB) *pgEventsRepository {
	return &pgEventsRepository{
		db: db,
	}
}

func (repo *pgEventsRepository) Create(event *entities.Event) error {
	stm, err := repo.db.Prepare(`
		INSERT INTO events (id,	date) VALUES ($1, $2, $3);
	`)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		event.Id,
		event.Date)

	return err
}

func (repo *pgEventsRepository) FindById(eventId string) (*entities.Event, error) {
	stm, err := repo.db.Prepare(`
		SELECT id, date FROM events WHERE id = $1;
	`)
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var event entities.Event

	err = stm.QueryRow(eventId).Scan(&event.Id, &event.Date)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &event, nil
}
