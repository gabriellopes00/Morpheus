package repositories

import (
	"context"
	"database/sql"
	"errors"
	"events/domain/entities"
	"events/framework/logger"
	"fmt"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type EventsRepository interface {
	Create(event *entities.Event) error
	FindAccountEvents(accountId string) ([]*entities.Event, error)
	FindById(eventId string) (*entities.Event, error)
	ExistsId(eventId string) (bool, error)
	SetStatus(eventId string, status entities.EventStatus) error
	FindAll(state string, month, ageGroup int) ([]entities.Event, error)
	Update(event *entities.Event) error
}

type pgEventsRepository struct {
	Db *sql.DB
}

func NewPgEventsRepository(connection *sql.DB) *pgEventsRepository {
	return &pgEventsRepository{connection}
}

func (repo *pgEventsRepository) Create(event *entities.Event) error {

	tx, err := repo.Db.BeginTx(context.Background(), nil)
	if err != nil {
		logger.Logger.Error("error while creating a new transaction", zap.String("error_message", err.Error()))
		return err
	}

	stm, err := tx.Prepare(`
		INSERT INTO events (id,
							name,
							description,
							organizer_account_id,
							age_group,
							maximum_capacity,
							status,
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
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15,
		$16, $17, $18, $19, $20);
	`)
	if err != nil {
		logger.Logger.Error("error while preparing a query", zap.String("error_message", err.Error()))
		logger.Logger.Info("rolling back transacion", zap.Error(tx.Rollback()))

		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		event.Id,
		event.Name,
		event.Description,
		event.OrganizerAccountId,
		strconv.Itoa(event.AgeGroup),
		event.MaximumCapacity,
		event.Status,
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
		logger.Logger.Error("error while executing a query", zap.String("error_message", err.Error()))
		logger.Logger.Info("rolling back transacion", zap.String("error", tx.Rollback().Error()))

		return err
	}

	for _, tycketOption := range event.TycketOptions {

		stm, err = tx.Prepare(`
			INSERT INTO event_tycket_options (id,
											event_id,
											title,
											created_at)
			VALUES ($1, $2, $3, $4);
		`)
		if err != nil {
			logger.Logger.Error("error while preparing a query", zap.String("error_message", err.Error()))
			logger.Logger.Info("rolling back transacion", zap.String("error", tx.Rollback().Error()))

			return err
		}

		defer stm.Close()

		_, err = stm.Exec(
			tycketOption.Id,
			tycketOption.EventId,
			tycketOption.Title,
			time.Now(),
		)
		if err != nil {
			logger.Logger.Error("error while executing a query", zap.String("error_message", err.Error()))
			logger.Logger.Info("rolling back transacion", zap.String("error", tx.Rollback().Error()))

			return err
		}

		lotQuery := `
			INSERT INTO event_tycket_lots (id,
											event_tycket_option_id,
											number,
											tycket_price,
											tycket_amount,
											created_at)
			VALUES ($1, $2, $3, $4, $5, $6);
		`

		for _, lot := range tycketOption.Lots {
			stm, err = tx.Prepare(lotQuery)
			if err != nil {
				logger.Logger.Error("error while preparing a query", zap.String("error_message", err.Error()))
				logger.Logger.Info("rolling back transacion", zap.String("error", tx.Rollback().Error()))

				return err
			}
			_, err = stm.Exec(
				uuid.NewV4(),
				lot.TycketOptionId,
				lot.Number,
				lot.TycketPrice,
				lot.TycketAmount,
				time.Now(),
			)
			if err != nil {
				logger.Logger.Error("error while executing a query", zap.String("error_message", err.Error()))
				logger.Logger.Info("rolling back transacion", zap.String("error", tx.Rollback().Error()))

				return err
			}

		}
	}

	err = tx.Commit()
	fmt.Println(err)
	return err
}

func (repo *pgEventsRepository) SetStatus(eventId string, status entities.EventStatus) error {
	stm, err := repo.Db.Prepare(`
		UPDATE events SET status = $1 WHERE id = $2 AND deleted_at IS NULL;
	`)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(eventId, status)

	return err
}

func (repo *pgEventsRepository) Update(event *entities.Event) error {
	stm, err := repo.Db.Prepare(`
		UPDATE events
		SET    name = $1,
			description = $2,
			is_available = $3,
			age_group = $4,
			maximum_capacity = $5,
			duration = $7,
			ticket_price = $8,
			date = $9,
			location_street = $10,
			location_district = $11,
			location_state = $12,
			location_city = $13,
			location_postal_code = $14,
			location_description = $15,
			location_number = $16,
			location_latitude = $17,
			location_longitude = $18
		WHERE  id = $19 AND deleted_at IS NULL; 
	`)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		event.Name,
		event.Description,
		event.AgeGroup,
		event.MaximumCapacity,
		event.Duration,
		event.Date,
		event.Location.Street,
		event.Location.District,
		event.Location.State,
		event.Location.City,
		event.Location.PostalCode,
		event.Location.Description,
		event.Location.Number,
		event.Location.Latitude,
		event.Location.Longitude,
		event.Id)

	return err
}

func (repo *pgEventsRepository) FindAccountEvents(accountId string) ([]*entities.Event, error) {
	stm, err := repo.Db.Prepare(`
		SELECT id,
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
				updated_at
		FROM events
		WHERE organizer_account_id = $1 
		AND deleted_at IS NULL;
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
		var event entities.Event
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.OrganizerAccountId,
			&event.AgeGroup,
			&event.MaximumCapacity,
			&event.Status,
			&event.Date,
			&event.Duration,
			&event.Location.Street,
			&event.Location.District,
			&event.Location.State,
			&event.Location.City,
			&event.Location.PostalCode,
			&event.Location.Description,
			&event.Location.Number,
			&event.Location.Latitude,
			&event.Location.Longitude,
			&event.CreatedAt,
			&event.UpdatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}

func (repo *pgEventsRepository) FindById(eventId string) (*entities.Event, error) {
	stm, err := repo.Db.Prepare(`
		SELECT id,
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
				updated_at
		FROM events
		WHERE id = $1 
		AND deleted_at IS NULL;
	`)
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var event entities.Event

	err = stm.QueryRow(eventId).Scan(
		&event.Id,
		&event.Name,
		&event.Description,
		&event.OrganizerAccountId,
		&event.AgeGroup,
		&event.MaximumCapacity,
		&event.Status,
		&event.Date,
		&event.Duration,
		&event.Location.Street,
		&event.Location.District,
		&event.Location.State,
		&event.Location.City,
		&event.Location.PostalCode,
		&event.Location.Description,
		&event.Location.Number,
		&event.Location.Latitude,
		&event.Location.Longitude,
		&event.CreatedAt,
		&event.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &event, nil
}

func (repo *pgEventsRepository) FindAll(state string, month, ageGroup int) ([]entities.Event, error) {
	stm, err := repo.Db.Prepare(`
		SELECT id,
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
				updated_at
		FROM events
		WHERE EXTRACT(MONTH FROM date) = $1 
		OR state = $2 OR age_group = $3
		AND deleted_at IS NULL;
	`)
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	events := []entities.Event{}

	rows, err := stm.Query(state, month, ageGroup)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event entities.Event
		err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.Description,
			&event.OrganizerAccountId,
			&event.AgeGroup,
			&event.MaximumCapacity,
			&event.Status,
			&event.Date,
			&event.Duration,
			&event.Location.Street,
			&event.Location.District,
			&event.Location.State,
			&event.Location.City,
			&event.Location.PostalCode,
			&event.Location.Description,
			&event.Location.Number,
			&event.Location.Latitude,
			&event.Location.Longitude,
			&event.CreatedAt,
			&event.UpdatedAt)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (r *pgEventsRepository) ExistsId(eventId string) (bool, error) {
	stm, err := r.Db.Prepare(`
		SELECT EXISTS(SELECT 1 FROM events WHERE id = $1 AND deleted_at IS NULL);
	`)
	if err != nil {
		return false, err
	}

	defer stm.Close()

	var exists bool

	err = stm.QueryRow(eventId).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
