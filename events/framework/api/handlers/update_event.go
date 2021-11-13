package handlers

import (
	"encoding/json"
	"events/application"
	"events/domain/usecases"
	"events/framework/queue"
	"net/http"

	"github.com/labstack/echo/v4"
)

type updateEventHandler struct {
	Usecase      usecases.UpdateEvents
	MessageQueue queue.MessageQueue
}

func NewUpdateEventHandler(usecase usecases.UpdateEvents, messageQueue queue.MessageQueue) *updateEventHandler {
	return &updateEventHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
	}
}

func (handler *updateEventHandler) Handle(c echo.Context) error {
	var params application.UpdateEventDTO

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

	eventId := c.Param("id")

	event, err := handler.Usecase.UpdateData(eventId, &params)
	if err != nil {
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

	err = handler.MessageQueue.PublishMessage(queue.ExchangeEvents, queue.KeyEventUpdated, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{"event": event},
	)
}
