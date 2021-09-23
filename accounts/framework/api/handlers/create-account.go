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

type createAccountHandler struct {
	Usecase      domain.CreateAccount
	Encrypter    interfaces.Encrypter
	MessageQueue interfaces.MessageQueue
}

func NewCreateAccountHandler(
	usecase domain.CreateAccount,
	messageQueue interfaces.MessageQueue,
	encrypter interfaces.Encrypter,
) *createAccountHandler {
	return &createAccountHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
		Encrypter:    encrypter,
	}
}

func (h *createAccountHandler) Create(c echo.Context) error {
	var params domain.Account

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.String(http.StatusBadRequest, ErrInvalidRequest.Error())
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

	err = h.MessageQueue.SendMessage(interfaces.QueueAccountCreated, []byte(payload))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	authToken, err := h.Encrypter.Encrypt(account)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{"account": account, "auth_token": authToken},
	)
}
