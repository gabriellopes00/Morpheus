package handlers

import (
	usecases "accounts/application"
	"accounts/domain"
	"errors"
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

func (h *refreshAuthHandler) Auth(c echo.Context) error {

	params := map[string]string{}

	if err := (&echo.DefaultBinder{}).BindBody(c, &params); err != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"error": err.Error()},
		)
	}

	refreshToken := params["refresh_token"]

	accessToken, err := h.Usecase.Refresh(refreshToken)
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

	return c.JSON(http.StatusOK, map[string]string{"access_token": accessToken})
}
