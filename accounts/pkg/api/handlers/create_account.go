package handlers

import (
	"accounts/application"
	app_error "accounts/domain/errors"
	"accounts/pkg/auth"
	"accounts/pkg/queue"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createAccountHandler struct {
	Usecase      application.CreateAccount
	AuthProvider auth.AuthProvider
	MessageQueue queue.MessageQueue
}

func NewCreateAccountHandler(
	usecase application.CreateAccount,
	messageQueue queue.MessageQueue,
	authProvider auth.AuthProvider,
) *createAccountHandler {
	return &createAccountHandler{
		Usecase:      usecase,
		MessageQueue: messageQueue,
		AuthProvider: authProvider,
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
		fmt.Println(err)
		if errors.Is(err, application.ErrEmailAlreadyInUse) {
			return c.JSON(http.StatusConflict,
				map[string]string{"error": err.Error()})
		} else if app_error.IsAppError(err) {
			return c.JSON(http.StatusBadRequest,
				map[string]string{"error": err.Error()})
		} else {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": ErrInternalServer.Error()})
		}
	}

	token, err := h.AuthProvider.SignInUser(
		auth.AuthUserCredentials{
			Email:    account.Email,
			Password: params.Password,
		},
	)
	if err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": ErrInternalServer.Error()})
	}

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

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{
			"account":       account,
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
		},
	)
}
