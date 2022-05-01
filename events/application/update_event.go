package application

import (
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
	"events/framework/utils"
	"time"

	"github.com/asaskevich/govalidator"
)

type UpdateEvent struct {
	Repository repositories.EventsRepository
}

func NewUpdateEvent(repo repositories.EventsRepository) *UpdateEvent {
	return &UpdateEvent{
		Repository: repo,
	}
}

func (u *UpdateEvent) SetStatus(eventId string, status entities.EventStatus) error {
	return u.Repository.SetStatus(eventId, status)
}

type UpdateEventDTO struct {
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	CoverUrl      string         `json:"cover_url,omitempty"`
	AgeGroup      int            `json:"age_group,omitempty"`
	Location      LocationParams `json:"location,omitempty"`
	StartDateTime string         `json:"start_datetime,omitempty" gorm:"column:start_datetime"`
	EndDateTime   string         `json:"end_datetime,omitempty" gorm:"column:end_datetime"`
	CategoryId    string         `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	SubjectId     string         `json:"subject,omitempty" gorm:"foreignKey:SubjectId"`
	Visibility    string         `json:"visibility,omitempty"`
	Status        string         `json:"status,omitempty"`
}

func (u *UpdateEvent) UpdateData(eventId string, data *UpdateEventDTO) (*entities.Event, error) {

	event, err := u.Repository.FindById(eventId, false)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("event with given id does not exists")
	}

	if time.Until(event.StartDateTime) <= time.Hour*24 {
		return nil, errors.New("the event date cannot be updated less than 24 hours before the event starts")
	}

	if err := u.validate(data); err != nil {
		return nil, err
	}

	// TODO: when mergin objects, if the new start and end datetime is empty,
	// it sets the current datetime to empty wither
	err = utils.MergeObjects(event, data)
	if err != nil {
		return nil, err
	}

	event.UpdatedAt = time.Now()

	err = u.Repository.Update(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (*UpdateEvent) validate(data *UpdateEventDTO) error {

	if data.Visibility != "" {
		switch data.Visibility {
		case string(entities.VisibilityPublic),
			string(entities.VisibilityPrivate),
			string(entities.VisibilityInvitedOnly):
		default:
			return errors.New("invalid event visibility")
		}
	}

	if data.Status != "" {
		switch data.Status {
		case string(entities.StatusAvailable),
			string(entities.StatusCanceled),
			string(entities.StatusFinished),
			string(entities.StatusSoldOut):
		default:
			return errors.New("invalid event status")
		}
	}

	if data.SubjectId != "" && !govalidator.IsUUIDv4(data.SubjectId) {
		return errors.New("invalid subject id. it must be a valid uuid v4")
	}

	if data.CategoryId != "" && !govalidator.IsUUIDv4(data.CategoryId) {
		return errors.New("invalid category id. it must be a valid uuid v4")
	}

	if data.CoverUrl != "" && !govalidator.IsURL(data.CoverUrl) {
		return errors.New("invalid cover url. it must be a valid url")
	}

	switch data.AgeGroup {
	case 0, 10, 12, 14, 16, 18:
	default:
		{
			return errors.New("invalid age group. it must be either 0, 10, 12, 14, 16, 18")
		}
	}

	return nil
}
