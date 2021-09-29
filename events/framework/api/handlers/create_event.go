package handlers

import (
	"encoding/json"
	"errors"
	"events/domain/entities"
	domainErrs "events/domain/errors"
	"events/domain/usecases"
	"events/framework/queue"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createEventHandler struct {
	Usecase      usecases.CreateEvent
	MessageQueue queue.MessageQueue
}

func NewCreateEventHandler(usecase usecases.CreateEvent, messageQueue queue.MessageQueue) *createEventHandler {
	return &createEventHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
	}
}

var (
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServer      = errors.New("unexpected internal server error")
)

func (handler *createEventHandler) Create(c echo.Context) error {
	var params entities.Event

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"error": ErrUnprocessableEntity.Error()},
		)
	}

	accountId := c.Request().Header.Get("account_id")
	if accountId == "" {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{"error": ErrUnauthorized.Error()},
		)
	}

	params.OrganizerAccountId = accountId

	event, err := handler.Usecase.Create(&params)
	if err != nil {
		if !domainErrs.IsDomainError(err) {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"error": err.Error()},
			)
		}

		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)

	}

	payload, err := json.Marshal(event)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	err = handler.MessageQueue.PublishMessage(queue.ExchangeEvents, queue.QueueEventCreated, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{"event": event},
	)
}
