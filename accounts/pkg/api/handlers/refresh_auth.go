package handlers

import (
	"accounts/pkg/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type refreshAuthHandler struct {
	authProvider auth.AuthProvider
}

func NewRefreshAuthHandler(authProvider auth.AuthProvider) *refreshAuthHandler {
	return &refreshAuthHandler{
		authProvider: authProvider,
	}
}

func (h *refreshAuthHandler) Handle(c echo.Context) error {
	var params struct {
		RefreshToken string `json:"refresh_token,omitempty"`
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.NoContent(http.StatusUnprocessableEntity)
	}

	tokens, err := h.authProvider.RefreshAuth(params.RefreshToken)
	if err != nil { // internal error?
		return c.JSON(
			http.StatusUnauthorized,
			map[string]string{"error": err.Error()},
		)

	}

	return c.JSON(
		http.StatusOK,
		map[string]string{
			"access_token":  tokens.AccessToken,
			"refresh_token": tokens.RefreshToken,
		},
	)
}
