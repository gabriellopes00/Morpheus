package handlers

import (
	"events/domain/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type findEventsHandler struct {
	Usecase usecases.FindEvents
}

func NewFindEventsHandler(usecase usecases.FindEvents) *findEventsHandler {
	return &findEventsHandler{
		Usecase: usecase,
	}
}

func (handler *findEventsHandler) Handle(c echo.Context) error {

	eventId := c.Param("id")
	if eventId == "" {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": ErrBadRequest.Error()})
	}

	if c.Request().Header.Get("account_id") == "" {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{"error": ErrUnauthorized.Error()})
	}

	event, err := handler.Usecase.FindEventById(eventId)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()})
	}

	if event == nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{"error": "No events found with given id"})
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{"event": event})
}
