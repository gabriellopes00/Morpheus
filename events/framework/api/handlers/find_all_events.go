package handlers

import (
	"events/domain/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type findAllEventsHandler struct {
	Usecase usecases.FindEvents
}

func NewFindAllEventsHandler(usecase usecases.FindEvents) *findAllEventsHandler {
	return &findAllEventsHandler{
		Usecase: usecase,
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
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{"error": ErrUnauthorized.Error()})
	}

	events, err := handler.Usecase.FindAll(state, month, ageGroup)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()})
	}

	if len(events) == 0 {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{"error": "No events found with given id"})
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{"events": events})
}
