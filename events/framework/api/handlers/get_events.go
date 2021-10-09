package handlers

import (
	"errors"
	"events/domain/usecases"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getEventsHandler struct {
	Usecase usecases.GetEvents
}

func NewGetEventsHandler(usecase usecases.GetEvents) *getEventsHandler {
	return &getEventsHandler{
		Usecase: usecase,
	}
}

var (
	ErrBadRequest = errors.New("invalid request")
)

func (handler *getEventsHandler) Handle(c echo.Context) error {

	accountId := c.Param("account_id")
	if accountId == "" {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"error": ErrBadRequest.Error()},
		)
	}

	if c.Request().Header.Get("account_id") != accountId {
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{"error": ErrUnauthorized.Error()},
		)
	}

	events, err := handler.Usecase.GetAccountEvents(accountId)
	if err != nil {
		log.Println(err)
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{"events": events},
	)
}
