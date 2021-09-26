package handlers

import (
	"accounts/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getAccountHandler struct {
	Usecase domain.GetAccount
}

func NewGetAccountHandler(usecase domain.GetAccount) *getAccountHandler {
	return &getAccountHandler{
		Usecase: usecase,
	}
}

func (h *getAccountHandler) Handle(c echo.Context) error {

	accountId := c.Param("id")

	account, err := h.Usecase.GetById(accountId)
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

	account.Password = ""

	return c.JSON(
		http.StatusCreated,
		map[string]interface{}{"account": account},
	)
}
