package handlers

import (
	"encoding/json"
	"events/application"
	"events/domain/usecases"
	"events/framework/logger"
	"events/framework/queue"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type updateEventHandler struct {
	usecase      usecases.UpdateEvents
	MessageQueue queue.MessageQueue
}

func NewUpdateEventHandler(usecase usecases.UpdateEvents, messageQueue queue.MessageQueue) *updateEventHandler {
	return &updateEventHandler{
		usecase:      usecase,
		MessageQueue: messageQueue,
	}
}

func (handler *updateEventHandler) Handle(c echo.Context) error {
	var params application.UpdateEventDTO

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	accountId := c.Request().Header.Get("account_id")
	if accountId == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	eventId := c.Param("id")

	event, err := handler.usecase.UpdateData(eventId, &params)
	if err != nil {
		logger.Logger.Error("error while updating event", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)

	}

	payload, _ := json.Marshal(event)
	err = handler.MessageQueue.PublishMessage(queue.ExchangeEvents, queue.KeyEventUpdated, payload)
	if err != nil {
		logger.Logger.Error("error while publishing message to the queue", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"event": event})
}
