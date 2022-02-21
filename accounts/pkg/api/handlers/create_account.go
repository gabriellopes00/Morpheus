package handlers

import (
	"accounts/application"
	"accounts/pkg/encrypter"
	"accounts/pkg/queue"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createAccountHandler struct {
	Usecase      application.CreateAccount
	Encrypter    encrypter.Encrypter
	MessageQueue queue.MessageQueue
}

func NewCreateAccountHandler(
	usecase application.CreateAccount,
	messageQueue queue.MessageQueue,
	encrypter encrypter.Encrypter,
) *createAccountHandler {
	return &createAccountHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
		Encrypter:    encrypter,
	}
}

func (h *createAccountHandler) Handle(c echo.Context) error {
	var params *application.CreateAccountDTO

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"error": "invalid request params"})
	}

	account, err := h.Usecase.Create(params)
	if err != nil {
		if errors.Is(err, application.ErrEmailAlreadyInUse) {
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

	err = h.MessageQueue.PublishMessage(queue.ExchangeAccounts, queue.KeyAccountCreated, payload)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()})
	}

	token, err := h.Encrypter.EncryptAuthToken(account.Id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()})
	}

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{
			"account":       account,
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
		},
	)
}
