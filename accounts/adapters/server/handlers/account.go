package handlers

import (
	"accounts/application/ports"
	"accounts/domain"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type accountHandler struct {
	Usecase      domain.AccountUsecase
	MessageQueue ports.MessageQueue
}

func NewAccountHandler(AccountUsecase domain.AccountUsecase, messageQueue ports.MessageQueue) *accountHandler {
	return &accountHandler{
		Usecase:      AccountUsecase,
		MessageQueue: messageQueue,
	}
}

func (h *accountHandler) Create(c echo.Context) error {
	var params domain.Account

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	account, err := h.Usecase.Create(params)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyInUse) {
			return c.JSON(http.StatusConflict, err.Error())
		} else {
			return c.JSON(http.StatusInternalServerError, map[string]string{"asdf": "asdf"})
		}
	}

	payload, err := json.Marshal(account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"asdf": "asdf"})
	}

	err = h.MessageQueue.SendMessage([]byte(payload))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"asdf": "asdf"})
	}

	return c.JSON(http.StatusCreated, account)
}

func (h *accountHandler) Auth(c echo.Context) error {
	var params domain.AuthCredentials

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	authToken, err := h.Usecase.Auth(params)
	if err != nil {
		if errors.Is(err, domain.ErrUnregisteredEmail) {
			return c.JSON(http.StatusConflict, err.Error())
		} else {
			log.Println(err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"asdf": "asdf"})
		}
	}

	return c.JSON(http.StatusCreated, map[string]string{"auth_token": authToken})
}
