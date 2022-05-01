package handlers

import (
	"events/application"
	"events/framework/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type findNearbyEventsHandler struct {
	findEvent *application.FindEvents
}

func NewFindNearbyEventsHandler(findEvent *application.FindEvents) *findNearbyEventsHandler {
	return &findNearbyEventsHandler{
		findEvent: findEvent,
	}
}

func (handler *findNearbyEventsHandler) Handle(c echo.Context) error {

	lat := c.QueryParam("latitude")
	lng := c.QueryParam("longitude")

	_, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "Invalid query params. Coordinates must be float numbers"})
	}

	_, err = strconv.ParseFloat(lng, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			map[string]string{"error": "Invalid query params. Coordinates must be float numbers"})
	}

	events, err := handler.findEvent.FindNearbyEvents(lat, lng)
	if err != nil {
		logger.Logger.Error("error while finding event", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	if events == nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{"error": "No events found with given id"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"events": events})
}
