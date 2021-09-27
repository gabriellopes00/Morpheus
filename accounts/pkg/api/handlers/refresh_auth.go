package handlers

import (
	"accounts/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type refreshAuthHandler struct {
	Usecase domain.RefreshAuth
}

func NewRefreshAuthHandler(usecase domain.RefreshAuth) *refreshAuthHandler {
	return &refreshAuthHandler{
		Usecase: usecase,
	}
}

func (h *refreshAuthHandler) Handle(c echo.Context) error {
	var params struct {
		RefreshToken string `json:"refresh_token,omitempty"`
	}

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"error": ErrUnprocessableEntity.Error()},
		)
	}

	tokens, err := h.Usecase.Refresh(params.RefreshToken)
	if err != nil {
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