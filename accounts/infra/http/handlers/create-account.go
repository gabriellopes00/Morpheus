package handlers

import (
	"accounts/entities"
	"accounts/usecases"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountHandler struct {
	CreateAccount entities.CreateAccount
}

func NewCreateAccountHandler(createAccount entities.CreateAccount) *createAccountHandler {
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
		if errors.Is(err, usecases.ErrEmailAlerayInUse) {
			c.String(http.StatusConflict, err.Error())
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
		}

		return
	}

	c.JSON(http.StatusCreated, account)
}
