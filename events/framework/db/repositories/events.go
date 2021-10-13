package repositories

import (
	"database/sql"
	"events/domain/entities"
	"strconv"
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
	stm, err := repo.Db.Prepare(`
		INSERT INTO EVENTS (id,
							name,
							description,
							is_available,
							organizer_account_id,
							age_group,
							maximum_capacity,
							status,
							ticket_price,
							date,
							duration,
							location_street,
							location_district,
							location_state,
							location_city,
							location_postal_code,
							location_description,
							location_number,
							location_latitude,
							location_longitude,
							created_at,
							updated_at)
		VALUES ($1, $2, NULLIF($3, ''), $4, $5, $6, $7, $8, $9, $10, 
		NULLIF($11, 0), $12, $13, $14, $15, $16, NULLIF($17, ''), 
		$18, NULLIF($19, 0::float), NULLIF($20, 0::float), $21, $22);
	`)
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
		strconv.Itoa(event.AgeGroup),
		event.MaximumCapacity,
		event.Status,
		event.TicketPrice,
		event.Date,
		event.Duration,
		event.Location.Street,
		event.Location.District,
		event.Location.State,
		event.Location.City,
		event.Location.PostalCode,
		event.Location.Description,
		event.Location.Number,
		event.Location.Latitude,
		event.Location.Longitude,
		event.CreatedAt,
		event.UpdatedAt)

	return err
}

func (repo *pgEventsRepository) GetAccountEvents(accountId string) ([]*entities.Event, error) {
	stm, err := repo.Db.Prepare(`
		SELECT (id,
				name,
				description,
				is_available,
				organizer_account_id,
				age_group,
				maximum_capacity,
				status,
				ticket_price,
				date,
				duration,
				location_street,
				location_district,
				location_state,
				location_city,
				location_postal_code,
				location_description,
				location_number,
				location_latitude,
				location_longitude,
				created_at,
				updated_at)
		FROM events
		WHERE organizer_account_id = $1 
		AND WHERE deleted_at IS NULL;
	`)
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
		event := &entities.Event{}
		err := rows.Scan(
			event.Id,
			event.Name,
			event.Description,
			event.IsAvailable,
			event.OrganizerAccountId,
			event.AgeGroup,
			event.MaximumCapacity,
			event.Status,
			event.TicketPrice,
			event.Date,
			event.Duration,
			event.Location.Street,
			event.Location.District,
			event.Location.State,
			event.Location.City,
			event.Location.PostalCode,
			event.Location.Description,
			event.Location.Number,
			event.Location.Latitude,
			event.Location.Longitude,
			event.CreatedAt,
			event.UpdatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
