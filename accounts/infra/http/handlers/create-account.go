package handlers

import (
	"accounts/application"
	"accounts/domain/entities"
	"accounts/domain/usecases"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountHandler struct {
	CreateAccount usecases.CreateAccount
}

func NewCreateAccountHandler(createAccount usecases.CreateAccount) *createAccountHandler {
	return &createAccountHandler{
		CreateAccount: createAccount,
	}
}

func (h *createAccountHandler) Create(c *gin.Context) {
	var params entities.Account

	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	account, err := h.CreateAccount.Create(params)
	if err != nil {
		if errors.Is(err, application.ErrEmailAlerayInUse) {
			c.String(http.StatusConflict, err.Error())
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		return
	}

	c.JSON(http.StatusCreated, account)
}
