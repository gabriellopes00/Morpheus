package handlers

import (
	"accounts/application"
	"accounts/domain"
	"accounts/pkg/encrypter"
	"accounts/pkg/queue"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deleteAccountHandler struct {
	Usecase      domain.DeleteAccount
	Encrypter    encrypter.Encrypter
	MessageQueue queue.MessageQueue
}

func NewDeleteAccountHandler(
	usecase domain.DeleteAccount,
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
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	err = h.MessageQueue.PublishMessage(queue.ExchangeAccounts, queue.KeyAccountDeleted, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	return c.JSON(http.StatusNoContent, nil)
}
