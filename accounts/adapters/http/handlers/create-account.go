package handlers

import (
	"accounts/application/ports"
	"accounts/domain"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountHandler struct {
	CreateAccount domain.CreateAccount
	MessageQueue  ports.MessageQueue
}

func NewCreateAccountHandler(
	createAccount domain.CreateAccount,
	messageQueue ports.MessageQueue,
) *createAccountHandler {

	return &createAccountHandler{
		CreateAccount: createAccount,
		MessageQueue:  messageQueue,
	}

}

func (h *createAccountHandler) Create(c *gin.Context) {
	var params domain.Account

	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	account, err := h.CreateAccount.Create(params)
	if err != nil {
		if errors.Is(err, domain.ErrEmailAlreadyInUse) {
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
