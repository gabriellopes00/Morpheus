package handlers

import (
	"accounts/domain"
	"accounts/usecases"
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

func (h *authHandler) Auth(c echo.Context) error {

	var params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	authToken, err := h.Usecase.Auth(params.Email, params.Password)
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

	return c.JSON(http.StatusOK, map[string]string{"auth_token": authToken})
}
