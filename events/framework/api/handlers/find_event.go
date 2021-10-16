package handlers

import (
	"events/domain/usecases"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type findAccountEventsHandler struct {
	Usecase usecases.FindEvents
}

func NewFindAccountEventsHandler(usecase usecases.FindEvents) *findAccountEventsHandler {
	return &findAccountEventsHandler{
		Usecase: usecase,
	}
}

func (handler *findAccountEventsHandler) Handle(c echo.Context) error {

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

	events, err := handler.Usecase.FindAccountEvents(accountId)
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
