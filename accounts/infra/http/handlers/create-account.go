package handlers

import (
	"accounts/application"
	"accounts/domain/entities"
	"accounts/domain/usecases"
	"accounts/ports"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountHandler struct {
	CreateAccount usecases.CreateAccount
	MessageQueue  ports.MessageQueue
}

func NewCreateAccountHandler(
	createAccount usecases.CreateAccount,
	messageQueue ports.MessageQueue,
) *createAccountHandler {

	return &createAccountHandler{
		CreateAccount: createAccount,
		MessageQueue:  messageQueue,
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

	payload, err := json.Marshal(account)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	err = h.MessageQueue.SendMessage([]byte(payload))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, account)
}
