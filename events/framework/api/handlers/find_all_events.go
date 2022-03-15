package handlers

import (
	"events/domain/usecases"
	"events/framework/logger"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type findAllEventsHandler struct {
	usecase usecases.FindEvents
}

func NewFindAllEventsHandler(usecase usecases.FindEvents) *findAllEventsHandler {
	return &findAllEventsHandler{
		usecase: usecase,
	}
}

func (handler *findAllEventsHandler) Handle(c echo.Context) error {

	state := c.QueryParam("state")
	month, err := strconv.Atoi(c.QueryParam("month"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	ageGroup, err := strconv.Atoi(c.QueryParam("age_group"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if c.Request().Header.Get("account_id") == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	events, err := handler.usecase.FindAll(state, month, ageGroup)
	if err != nil {
		logger.Logger.Error("error while finding events", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	if len(events) == 0 {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{"error": "No events found with given id"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"events": events})
}
