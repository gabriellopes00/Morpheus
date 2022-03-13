package handlers

import (
	"accounts/application"
	"accounts/pkg/encrypter"
	"accounts/pkg/queue"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deleteAccountHandler struct {
	Usecase      application.DeleteAccount
	Encrypter    encrypter.Encrypter
	MessageQueue queue.MessageQueue
}

func NewDeleteAccountHandler(
	usecase application.DeleteAccount,
	messageQueue queue.MessageQueue,
) *deleteAccountHandler {
	return &deleteAccountHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
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

	err := h.Usecase.Delete(accountId)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, application.ErrIdNotFound) {
			return c.JSON(http.StatusConflict, err.Error())
		} else {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": ErrInternalServer.Error()},
			)
		}
	}

	payload, err := json.Marshal(map[string]string{"id": accountId})
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = h.MessageQueue.PublishMessage(queue.ExchangeAccounts, queue.KeyAccountDeleted, payload)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusNoContent, nil)
}
