package handlers

import (
	"encoding/json"
	"events/application"
	"events/domain/entities"
	"events/framework/logger"
	"events/framework/queue"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type CancelEventHandler struct {
	updateEvent  *application.UpdateEvent
	findEvent    *application.FindEvents
	messageQueue queue.MessageQueue
}

func NewCancelEventHandler(updateEvent *application.UpdateEvent, messageQueue queue.MessageQueue) *CancelEventHandler {
	return &CancelEventHandler{
		updateEvent:  updateEvent,
		messageQueue: messageQueue,
	}
}

func (handler *CancelEventHandler) Handle(c echo.Context) error {

	accountId := c.Request().Header.Get("account_id")
	if accountId == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	eventId := c.Param("event_id")

	event, err := handler.findEvent.FindEventById(eventId)
	if err != nil {
		logger.Logger.Error("error while finding an event", zap.String("error_msg", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	if event.OrganizerAccountId != accountId {
		return c.NoContent(http.StatusForbidden)
	}

	err = handler.updateEvent.SetStatus(event.Id, entities.StatusCanceled)
	if err != nil {
		logger.Logger.Error("error while updating an event", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	event.Status = entities.StatusCanceled
	payload, _ := json.Marshal(event)
	err = handler.messageQueue.PublishMessage(queue.ExchangeEvents, queue.KeyEventCanceled, payload)
	if err != nil {
		logger.Logger.Error("error while publishing message to the queue", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"event": event})
}
