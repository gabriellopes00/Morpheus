package repositories

import (
	"events/domain/entities"
	"fmt"
	"strings"

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

	event.Location.PostalCode = strings.ReplaceAll(event.Location.PostalCode, "-", "")

	err := repo.Db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&event).Error
	})

	return err
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

	var models []entities.Event

	query := repo.Db.Table("events AS event")
	query.Where("event.organizer_account_id = ?", accountId).Find(&models)
	query.Preload("TycketOptions.Lots").Preload(clause.Associations).Find(&models)
	err := query.Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return models, nil
}

func (repo *pgEventsRepository) FindById(eventId string) (*entities.Event, error) {

	var model entities.Event

	query := repo.Db.Table("events AS event")
	query.Where("event.id = ?", eventId).First(&model)
	// query.Preload("Location").Find(&model)
	// query.Preload("TycketOptions").Find(&model)
	// query.Preload("TycketOptions.Lots").Find(&model)
	query.Preload("TycketOptions.Lots").Preload(clause.Associations).Find(&model)
	err := query.Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model, nil
}

func (repo *pgEventsRepository) FindAll(state string, month, ageGroup int) ([]entities.Event, error) {

	var models []entities.Event

	query := repo.Db.Table("events")
	query.Where("events.age_group = ?", ageGroup).Find(&models)
	query.Preload("Location").Find(&models)
	query.Or(&entities.Event{Location: &entities.EventLocation{State: state}}).Find(&models)
	err := query.Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return models, nil

}

func (r *pgEventsRepository) ExistsId(eventId string) (bool, error) {
	var exists bool
	err := r.Db.Raw("SELECT EXISTS(SELECT 1 FROM events WHERE id = ?);", eventId).Scan(&exists).Error
	if err != nil {
		return false, err
	}

	return exists, nil
}
