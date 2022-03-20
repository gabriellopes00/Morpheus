package handlers

import (
	"events/application"
	"events/framework/logger"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type findEventsHandler struct {
	findEvent *application.FindEvents
}

func NewFindEventsHandler(findEvent *application.FindEvents) *findEventsHandler {
	return &findEventsHandler{
		findEvent: findEvent,
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
		return c.NoContent(http.StatusUnauthorized)
	}

	event, err := handler.findEvent.FindEventById(eventId)
	if err != nil {
		logger.Logger.Error("error while finding event", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	if event == nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{"error": "No events found with given id"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"event": event})
}
