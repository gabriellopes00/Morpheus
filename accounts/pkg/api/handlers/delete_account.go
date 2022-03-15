package handlers

import (
	"accounts/application"
	"accounts/pkg/logger"
	"accounts/pkg/queue"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type deleteAccountHandler struct {
	usecase      application.DeleteAccount
	messageQueue queue.MessageQueue
}

func NewDeleteAccountHandler(
	usecase application.DeleteAccount,
	messageQueue queue.MessageQueue,
) *deleteAccountHandler {
	return &deleteAccountHandler{
		usecase:      usecase,
		messageQueue: messageQueue,
	}
}

func (h *deleteAccountHandler) Handle(c echo.Context) error {
	accountId := c.Request().Header.Get("account_id")
	paramId := c.Param("id")

	if accountId != paramId {
		return c.JSON(
			http.StatusForbidden,
			map[string]string{"error": "forbidden account deletion"})
	}

	err := h.usecase.Delete(accountId)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, application.ErrIdNotFound) {
			return c.JSON(http.StatusConflict, err.Error())
		} else {
			logger.Logger.Error("error while deleting a user", zap.String("error_message", err.Error()))
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	payload, _ := json.Marshal(map[string]string{"id": accountId})
	err = h.messageQueue.PublishMessage(queue.ExchangeAccounts, queue.KeyAccountDeleted, payload)
	if err != nil {
		logger.Logger.Error("error while publishing message to the queue", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusNoContent, nil)
}
