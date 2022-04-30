package handlers

import (
	"events/application"
	"events/framework/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type findNearEventsHandler struct {
	findEvent *application.FindEvents
}

func NewFindNearEventsHandler(findEvent *application.FindEvents) *findNearEventsHandler {
	return &findNearEventsHandler{
		findEvent: findEvent,
	}
}

func (handler *findNearEventsHandler) Handle(c echo.Context) error {

	lat := c.QueryParam("latitude")
	lng := c.QueryParam("longitude")

	latitude, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "Invalid query params. Coordinates must be float numbers"})
	}

	longitude, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "Invalid query params. Coordinates must be float numbers"})
	}

	event, err := handler.findEvent.FindNearEvents(latitude, longitude)
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
