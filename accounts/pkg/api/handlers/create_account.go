package handlers

import (
	"accounts/application"
	app_error "accounts/domain/errors"
	"accounts/pkg/auth"
	"accounts/pkg/logger"
	"accounts/pkg/queue"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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
		return c.NoContent(http.StatusUnprocessableEntity)
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
			return c.NoContent(http.StatusInternalServerError)
		}
	}

	token, err := h.AuthProvider.SignInUser(params.Email, params.Password)
	if err != nil {
		fmt.Println(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	payload, err := json.Marshal(account)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	err = h.MessageQueue.PublishMessage(queue.ExchangeAccounts, queue.KeyAccountCreated, payload)
	if err != nil {
		logger.Logger.Error("error while publishing message to the queue", zap.String("rabbitmq", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
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
