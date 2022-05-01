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
	FindByLocation(state, city string) ([]entities.Event, error)
	ExistsId(eventId string) (bool, error)
	SetStatus(eventId string, status entities.EventStatus) error
	FindAll(state string, month, ageGroup, limit, offset int) ([]entities.Event, error)
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
	query := repo.Db.Model(&entities.Event{})
	query.Where("id = ?", eventId)
	query.Update("status", status)
	return query.Error
}

func (repo *pgEventsRepository) Update(event *entities.Event) error {
	query := repo.Db.Table("events AS event")
	query.Where("event.id = ?", event.Id)
	query.Updates(
		map[string]interface{}{
			"name":           event.Name,
			"description":    event.Description,
			"cover_url":      event.CoverUrl,
			"age_group":      event.AgeGroup,
			"status":         event.Status,
			"start_datetime": event.StartDateTime,
			"end_datetime":   event.EndDateTime,
			"category_id":    event.CategoryId,
			"subject_id":     event.SubjectId,
			"visibility":     event.Visibility,
			"updated_at":     event.UpdatedAt,
		},
	)

	return query.Error
}

func (repo *pgEventsRepository) FindAccountEvents(accountId string) ([]entities.Event, error) {

	var models []entities.Event

	query := repo.Db.Table("events AS event")
	query.Where("event.organizer_account_id = ?", accountId)
	query.Preload("TicketOptions.Lots").Preload(clause.Associations).Find(&models)
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
	query.Where("event.id = ?", eventId)
	query.Preload("TicketOptions.Lots").Preload(clause.Associations).Find(&model)
	err := query.Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model, nil
}

func (repo *pgEventsRepository) FindByLocation(state, city string) ([]entities.Event, error) {

	var models []entities.Event

	fmt.Println(city)

	query := repo.Db.Table("events")
	query.Where("events.status != ?", entities.StatusCanceled) // ignore canceled events
	query.Preload("Location")
	query.Joins("JOIN event_locations Location ON events.id = Location.event_id AND Location.city = ?", city)

	query.Limit(100)

	query.Find(&models)
	err := query.Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return models, nil
}

func (repo *pgEventsRepository) FindAll(state string, month, ageGroup, limit, offset int) ([]entities.Event, error) {

	var models []entities.Event

	query := repo.Db.Table("events")

	query.Preload("Location")
	query.Where(&entities.Event{Location: &entities.EventLocation{State: state}})
	query.Where("events.status != ?", entities.StatusCanceled) // ignore canceled events

	if ageGroup > 0 {
		query.Where("events.age_group <= ?", ageGroup)
	}

	if month > 0 {
		query.Where("EXTRACT(MONTH FROM events.start_datetime) = ?", month)
	}

	query.Order("events.start_datetime ASC")
	query.Order("events.age_group DESC")

	query.Limit(limit)
	query.Offset(offset)

	query.Find(&models)
	err := query.Error
	if err != nil {
		return nil, err
	}

	return models, nil

}

func (r *pgEventsRepository) ExistsId(eventId string) (bool, error) {
	var exists bool
	err := r.Db.Raw("SELECT EXISTS(SELECT 1 FROM events WHERE id = ?)", eventId).Scan(&exists).Error
	if err != nil {
		return false, err
	}

	return exists, nil
}
