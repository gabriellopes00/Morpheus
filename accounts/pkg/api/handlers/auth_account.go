package handlers

import (
	usecases "accounts/application"
	"accounts/domain"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	Usecase domain.AuthAccount
}

func NewAuthHandler(usecase domain.AuthAccount) *authHandler {
	return &authHandler{
		Usecase: usecase,
	}
}

func (h *authHandler) Handle(c echo.Context) error {

	var params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"error": ErrUnprocessableEntity.Error()})
	}

	token, err := h.Usecase.Auth(params.Email, params.Password)
	if err != nil {
		if errors.Is(err, usecases.ErrUnregisteredEmail) {
			return c.JSON(http.StatusConflict, err.Error())
		} else {
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": ErrInternalServer.Error()},
			)
		}
	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"access_token":  token.AccessToken,
			"refresh_token": token.RefreshToken,
		},
	)
}
