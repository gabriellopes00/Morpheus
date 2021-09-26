package handlers

import (
	usecases "accounts/application"
	"accounts/domain"
	"errors"
	"fmt"
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
			map[string]string{"error": err.Error()},
		)
	}

	tokens, err := h.Usecase.Refresh(params.RefreshToken)
	if err != nil {
		fmt.Println(err)
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
			"access_token":  tokens.AccessToken,
			"refresh_token": tokens.RefreshToken,
		},
	)
}
