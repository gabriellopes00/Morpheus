package handlers

import (
	"accounts/application"
	"accounts/pkg/auth"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	AuthProvider auth.AuthProvider
}

func NewAuthHandler(AuthProvider auth.AuthProvider) *authHandler {
	return &authHandler{
		AuthProvider: AuthProvider,
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

	token, err := h.AuthProvider.SignInUser(params.Email, params.Password)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, application.ErrUnregisteredEmail) {
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
