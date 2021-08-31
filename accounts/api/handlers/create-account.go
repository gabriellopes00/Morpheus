package handlers

import (
	"accounts/domain"
	"accounts/interfaces"
	"accounts/usecases"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createAccountHandler struct {
	Usecase      domain.CreateAccount
	MessageQueue interfaces.MessageQueue
}

func NewCreateAccountHandler(usecase domain.CreateAccount, messageQueue interfaces.MessageQueue) *createAccountHandler {
	return &createAccountHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
	}
}

var (
	ErrInternalServer = errors.New("unexpected internal server error")
)

func (h *createAccountHandler) Create(c echo.Context) error {
	var params domain.Account

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	account, err := h.Usecase.Create(params)
	if err != nil {
		if errors.Is(err, usecases.ErrEmailAlreadyInUse) {
			return c.JSON(http.StatusConflict, err.Error())
		} else {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": ErrInternalServer.Error()},
			)
		}
	}

	payload, err := json.Marshal(account)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	err = h.MessageQueue.SendMessage([]byte(payload))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	return c.JSON(http.StatusCreated, account)
}
