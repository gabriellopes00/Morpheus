package handlers

import (
	"accounts/application"
	"accounts/pkg/logger"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type findAccountHandler struct {
	usecase application.FindAccount
}

func NewFindAccountHandler(usecase application.FindAccount) *findAccountHandler {
	return &findAccountHandler{
		usecase: usecase,
	}
}

func (h *findAccountHandler) Handle(c echo.Context) error {

	accountId := c.Param("id")

	account, err := h.usecase.FindById(accountId)
	if err != nil {
		logger.Logger.Error("error whilefinding account by id", zap.String("error_message", err.Error()))
		return c.NoContent(http.StatusInternalServerError)
	}

	if account == nil {
		return c.JSON(
			http.StatusNotFound,
			map[string]string{"error": "account not found"},
		)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"account": account})
}
