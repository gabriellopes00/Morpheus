package handlers

import (
	"accounts/application"
	"accounts/pkg/auth"
	"accounts/pkg/logger"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type authHandler struct {
	authProvider auth.AuthProvider
}

func NewAuthHandler(authProvider auth.AuthProvider) *authHandler {
	return &authHandler{
		authProvider: authProvider,
	}
}

func (h *authHandler) Handle(c echo.Context) error {

	var params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	token, err := h.authProvider.SignInUser(params.Email, params.Password)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, application.ErrUnregisteredEmail) {
			return c.JSON(http.StatusConflict, err.Error())
		} else {
			logger.Logger.Error("error while signin in account", zap.String("error_message", err.Error()))
			return c.NoContent(http.StatusInternalServerError)
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
