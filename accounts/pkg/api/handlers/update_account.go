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

type updateAccountHandler struct {
	usecase      application.UpdateAccount
	messageQueue queue.MessageQueue
}

func NewUpdateAccountHandler(usecase application.UpdateAccount, messageQueue queue.MessageQueue) *updateAccountHandler {
	return &updateAccountHandler{
		usecase:      usecase,
		messageQueue: messageQueue,
	}
}

func (h *updateAccountHandler) Handle(c echo.Context) error {
	accountId := c.Request().Header.Get("account_id")
	paramId := c.Param("id")

	if accountId != paramId {
		return c.NoContent(http.StatusForbidden)
	}

	var params *application.UpdateAccountDTO

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	account, err := h.usecase.Update(accountId, params)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, application.ErrIdNotFound) {
			return c.JSON(http.StatusConflict,
				map[string]string{"error": err.Error()})
		} else {
			logger.Logger.Error("error while updating account data", zap.String("error_message", err.Error()))
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	payload, _ := json.Marshal(account)
	err = h.messageQueue.PublishMessage(queue.ExchangeAccounts, queue.KeyAccountUpdated, payload)
	if err != nil {
		logger.Logger.Error("error while publishing message to the queue", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"account": account})
}
