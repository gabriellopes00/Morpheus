package handlers

import (
	"encoding/json"
	"events/domain/entities"
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

func (handler *createEventHandler) Create(c echo.Context) error {
	var params entities.Event

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.String(http.StatusUnprocessableEntity, "")
	}

	params.OrganizerAccountId = c.Request().Header.Get("account_id")

	event, err := handler.Usecase.Create(&params)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "unexpected internal server error"},
		)

	}

	payload, err := json.Marshal(event)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "unexpected internal server error"},
		)
	}

	err = handler.MessageQueue.PublishMessage("", "", payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "unexpected internal server error"},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{"event": event},
	)
}
