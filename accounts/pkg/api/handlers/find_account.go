package handlers

import (
	"accounts/domain/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type findAccountHandler struct {
	Usecase usecases.FindAccount
}

func NewFindAccountHandler(usecase usecases.FindAccount) *findAccountHandler {
	return &findAccountHandler{
		Usecase: usecase,
	}
}

func (h *findAccountHandler) Handle(c echo.Context) error {

	accountId := c.Param("id")

	account, err := h.Usecase.FindById(accountId)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{"error": "unexpected internal server error"},
		)
	}

	if account == nil {
		return c.JSON(
			http.StatusConflict,
			map[string]string{"error": "account not found"},
		)
	}

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{"account": account},
	)
}
