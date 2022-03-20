package repositories

import (
	"events/domain/entities"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EventsRepository interface {
	Create(event *entities.Event) error
	FindAccountEvents(accountId string) ([]entities.Event, error)
	FindById(eventId string) (*entities.Event, error)
	ExistsId(eventId string) (bool, error)
	SetStatus(eventId string, status entities.EventStatus) error
	FindAll(state string, month, ageGroup int) ([]entities.Event, error)
	Update(event *entities.Event) error
}

type pgEventsRepository struct {
	Db *gorm.DB
}

func NewPgEventsRepository(connection *gorm.DB) *pgEventsRepository {
	return &pgEventsRepository{connection}
}

func (repo *pgEventsRepository) Create(event *entities.Event) error {

	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&event).Error
	})

	if err != nil {
		return err
	}

	return nil
}

func (repo *pgEventsRepository) SetStatus(eventId string, status entities.EventStatus) error {
	// stm, err := repo.Db.Prepare(`
	// 	UPDATE events SET status = $1 WHERE id = $2;
	// `)
	// if err != nil {
	// 	return err
	// }

	// defer stm.Close()

	// _, err = stm.Exec(eventId, status)

	// return err
	return nil
}

func (repo *pgEventsRepository) Update(event *entities.Event) error {
	// stm, err := repo.Db.Prepare(`
	// 	UPDATE events
	// 	SET     name = $1,
	// 			description = $2,
	// 			age_group = $4,
	// 			maximum_capacity = $5,
	// 			duration = $7,
	// 			date = $9,
	// 			location_street = $10,
	// 			location_district = $11,
	// 			location_state = $12,
	// 			location_city = $13,
	// 			location_postal_code = $14,
	// 			location_description = $15,
	// 			location_number = $16,
	// 			location_latitude = $17,
	// 			location_longitude = $18
	// 	WHERE  id = $19;
	// `)
	// if err != nil {
	// 	return err
	// }

	// defer stm.Close()

	// _, err = stm.Exec(
	// 	event.Name,
	// 	event.Description,
	// 	event.AgeGroup,
	// 	event.MaximumCapacity,
	// 	event.Duration,
	// 	event.Date,
	// 	event.Location.Street,
	// 	event.Location.District,
	// 	event.Location.State,
	// 	event.Location.City,
	// 	event.Location.PostalCode,
	// 	event.Location.Description,
	// 	event.Location.Number,
	// 	event.Location.Latitude,
	// 	event.Location.Longitude,
	// 	event.Id)

	// return err
	return nil
}

func (repo *pgEventsRepository) FindAccountEvents(accountId string) ([]entities.Event, error) {

	var events []entities.Event
	err := repo.Db.Model(&events).Preload(clause.Associations).Find(&events).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(events)

	return events, nil
}

func (repo *pgEventsRepository) FindById(eventId string) (*entities.Event, error) {

	var model entities.Event

	query := repo.Db.Table("events AS event")
	query.Joins("INNER JOIN event_locations AS location ON location.event_id = event.id").Scan(&model.Location)
	query.Joins("INNER JOIN event_tycket_options AS tckt_opts ON tckt_opts.event_id = event.id").Scan(&model.TycketOptions)
	query.Where("event.id = ?", eventId)
	query.Find(&model)

	if query.Error != nil {
		return nil, query.Error
	}

	fmt.Println(model)

	// stm, err := repo.Db.Prepare(`
	// 	SELECT (
	// 		event.id,
	// 		event.name,
	// 		event.description,
	// 		event.organizer_account_id,
	// 		event.age_group,
	// 		event.maximum_capacity,
	// 		event.status,
	// 		event.date,
	// 		event.duration,
	// 		event.location_street,
	// 		event.location_district,
	// 		event.location_state,
	// 		event.location_city,
	// 		event.location_postal_code,
	// 		event.location_description,
	// 		event.location_number,
	// 		event.location_latitude,
	// 		event.location_longitude,
	// 		event.created_at,
	// 		event.updated_attckt_opt.id,

	// 		tckt_opt.event_id,
	// 		tckt_opt.title,
	// 		-- tckt_opt.created_at,

	// 		tckt_lot.id,
	// 		tckt_lot.event_tycket_option_id,
	// 		tckt_lot.number,
	// 		tckt_lot.tycket_price,
	// 		tckt_lot.tycket_amount
	// 		-- tckt_lot.created_at,
	// 	)
	// 	FROM events AS event

	// 	INNER JOIN event_tycket_options AS tckt_opt
	// 	ON tckt_opt.event_id = event.id

	// 	INNER JOIN event_tycket_lots AS tckt_lot
	// 	ON tckt_lot.event_tycket_option_id = tckt_opt.id

	// 	WHERE event.id = $1;
	// `)
	// if err != nil {
	// 	return nil, err
	// }

	// defer stm.Close()

	// rows, err := stm.Query(eventId)
	// if err != nil {
	// 	return nil, err
	// }

	// var event entities.Event

	// for rows.Next() {

	// 	var eventTycketOption entities.TycketOption
	// 	var eventTycketLot entities.TycketLot

	// 	err = rows.Scan(
	// 		&event.Id,
	// 		&event.Name,
	// 		&event.Description,
	// 		&event.OrganizerAccountId,
	// 		&event.AgeGroup,
	// 		&event.MaximumCapacity,
	// 		&event.Status,
	// 		&event.Date,
	// 		&event.Duration,
	// 		&event.Location.Street,
	// 		&event.Location.District,
	// 		&event.Location.State,
	// 		&event.Location.City,
	// 		&event.Location.PostalCode,
	// 		&event.Location.Description,
	// 		&event.Location.Number,
	// 		&event.Location.Latitude,
	// 		&event.Location.Longitude,
	// 		&event.CreatedAt,
	// 		&event.UpdatedAt,

	// 		// &eventTycketOption.Id,
	// 		// &eventTycketOption.EventId,
	// 		// &eventTycketOption.Title,

	// 		// &eventTycketLot.Id,
	// 		// &eventTycketLot.TycketOptionId,
	// 		// &eventTycketLot.Number,
	// 		// &eventTycketLot.TycketPrice,
	// 		// &eventTycketLot.TycketAmount,
	// 	)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	eventTycketOption.Lots = append(eventTycketOption.Lots, eventTycketLot)
	// 	event.TycketOptions = append(event.TycketOptions, eventTycketOption)
	// }

	// // err = stm.QueryRow(eventId).Scan(
	// // 	&event.Id,
	// // 	&event.Name,
	// // 	&event.Description,
	// // 	&event.OrganizerAccountId,
	// // 	&event.AgeGroup,
	// // 	&event.MaximumCapacity,
	// // 	&event.Status,
	// // 	&event.Date,
	// // 	&event.Duration,
	// // 	&event.Location.Street,
	// // 	&event.Location.District,
	// // 	&event.Location.State,
	// // 	&event.Location.City,
	// // 	&event.Location.PostalCode,
	// // 	&event.Location.Description,
	// // 	&event.Location.Number,
	// // 	&event.Location.Latitude,
	// // 	&event.Location.Longitude,
	// // 	&event.CreatedAt,
	// // 	&event.UpdatedAt)
	// // if err != nil {
	// // 	if errors.Is(err, gorm.ErrNoRows) {
	// // 		return nil, nil
	// // 	}

	// // 	return nil, err
	// // }

	// return &event, nil
	return &model, nil
}

func (repo *pgEventsRepository) FindAll(state string, month, ageGroup int) ([]entities.Event, error) {
	// stm, err := repo.Db.Prepare(`
	// 	SELECT id,
	// 			name,
	// 			description,
	// 			organizer_account_id,
	// 			age_group,
	// 			maximum_capacity,
	// 			status,
	// 			date,
	// 			duration,
	// 			location_street,
	// 			location_district,
	// 			location_state,
	// 			location_city,
	// 			location_postal_code,
	// 			location_description,
	// 			location_number,
	// 			location_latitude,
	// 			location_longitude,
	// 			created_at,
	// 			updated_at
	// 	FROM events
	// 	WHERE EXTRACT(MONTH FROM date) = $1
	// 	OR state = $2 OR age_group = $3;
	// `)
	// if err != nil {
	// 	return nil, err
	// }

	// defer stm.Close()

	// events := []entities.Event{}

	// rows, err := stm.Query(state, month, ageGroup)
	// if err != nil {
	// 	return nil, err
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var event entities.Event
	// 	err := rows.Scan(
	// 		&event.Id,
	// 		&event.Name,
	// 		&event.Description,
	// 		&event.OrganizerAccountId,
	// 		&event.AgeGroup,
	// 		&event.MaximumCapacity,
	// 		&event.Status,
	// 		&event.Date,
	// 		&event.Duration,
	// 		&event.Location.Street,
	// 		&event.Location.District,
	// 		&event.Location.State,
	// 		&event.Location.City,
	// 		&event.Location.PostalCode,
	// 		&event.Location.Description,
	// 		&event.Location.Number,
	// 		&event.Location.Latitude,
	// 		&event.Location.Longitude,
	// 		&event.CreatedAt,
	// 		&event.UpdatedAt)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	events = append(events, event)
	// }

	// return events, nil
	return nil, nil
}

func (r *pgEventsRepository) ExistsId(eventId string) (bool, error) {
	var exists bool
	err := r.Db.Raw("SELECT EXISTS(SELECT 1 FROM events WHERE id = ?);", eventId).Scan(&exists).Error
	if err != nil {
		return false, err
	}

	return exists, nil
}
