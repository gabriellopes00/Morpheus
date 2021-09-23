package handlers

import (
	usecases "accounts/application"
	"accounts/domain"
	"accounts/interfaces"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type deleteAccountHandler struct {
	Usecase      domain.DeleteAccount
	Encrypter    interfaces.Encrypter
	MessageQueue interfaces.MessageQueue
}

func NewDeleteAccountHandler(
	usecase domain.DeleteAccount,
	messageQueue interfaces.MessageQueue,
) *deleteAccountHandler {
	return &deleteAccountHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
	}
}

func (h *deleteAccountHandler) Delete(c echo.Context) error {

	accountId := c.Request().Header.Get("account_id")

	err := h.Usecase.Delete(accountId)
	if err != nil {
		if errors.Is(err, usecases.ErrIdNotFound) {
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

	err = h.MessageQueue.SendMessage(interfaces.QueueAccountDeleted, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	return c.JSON(http.StatusNoContent, nil)
}
