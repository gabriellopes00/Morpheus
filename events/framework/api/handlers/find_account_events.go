package handlers

import (
	"events/application"
	"events/framework/logger"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type findAccountEventsHandler struct {
	findAccountEvents *application.FindEvents
}

func NewFindAccountEventsHandler(findAccountEvents *application.FindEvents) *findAccountEventsHandler {
	return &findAccountEventsHandler{
		findAccountEvents: findAccountEvents,
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
		return c.NoContent(http.StatusUnauthorized)
	}

	events, err := handler.findAccountEvents.FindAccountEvents(accountId)
	if err != nil {
		logger.Logger.Error("error while find account events", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"events": events})
}
