package handlers

import (
	"encoding/json"
	"events/application"
	domainErrs "events/domain/errors"
	"events/framework/logger"
	"events/framework/queue"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type createEventHandler struct {
	createEvent  *application.CreateEvent
	messageQueue queue.MessageQueue
}

func NewCreateEventHandler(createEvent *application.CreateEvent, messageQueue queue.MessageQueue) *createEventHandler {
	return &createEventHandler{
		createEvent:  createEvent,
		messageQueue: messageQueue,
	}
}

func (handler *createEventHandler) Create(c echo.Context) error {
	var params application.CreateEventParams

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	accountId := c.Request().Header.Get("account_id")
	if accountId == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	params.OrganizerAccountId = accountId

	event, err := handler.createEvent.Create(&params)
	if err != nil {
		if !domainErrs.IsDomainError(err) {
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"error": err.Error()},
			)
		}

		logger.Logger.Error("error while publishing message to the queue", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	payload, _ := json.Marshal(event)
	err = handler.messageQueue.PublishMessage(queue.ExchangeEvents, queue.KeyEventCreated, payload)
	if err != nil {
		logger.Logger.Error("error while publishing message to the queue", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"event": event})
}
