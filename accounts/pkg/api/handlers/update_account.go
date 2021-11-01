package handlers

import (
	usecases "accounts/application"
	"accounts/domain"
	"accounts/pkg/queue"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type updateAccountHandler struct {
	Usecase      domain.UpdateAccount
	MessageQueue queue.MessageQueue
}

func NewUpdateAccountHandler(usecase domain.UpdateAccount, messageQueue queue.MessageQueue) *updateAccountHandler {
	return &updateAccountHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
	}
}

func (h *updateAccountHandler) Handle(c echo.Context) error {
	accountId := c.Request().Header.Get("account_id")
	paramId := c.Param("id")
	if accountId != paramId {
		return c.JSON(
			http.StatusForbidden,
			map[string]string{"error": "forbidden accout update"})
	}

	var params *domain.UpdateAccountDTO

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"error": "invalid request data"})
	}

	account, err := h.Usecase.Update(accountId, params)
	if err != nil {
		if errors.Is(err, usecases.ErrIdNotFound) {
			return c.JSON(http.StatusConflict,
				map[string]string{"error": err.Error()})
		} else {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": ErrInternalServer.Error()})
		}
	}

	account.Password = ""

	payload, err := json.Marshal(account)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()})
	}

	err = h.MessageQueue.PublishMessage(queue.ExchangeAccounts, queue.KeyAccountUpdated, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()})
	}

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{"account": account})
}
