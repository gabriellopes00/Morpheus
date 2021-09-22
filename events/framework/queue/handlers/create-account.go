package handlers

import (
	"encoding/json"
	"events/domain/entities"
	"events/domain/usecases"
)

type createAccounthandler struct {
	Channel <-chan []byte
	Usecase usecases.CreateAccount
}

func NewCreateAccountHandler(channel <-chan []byte, usecase usecases.CreateAccount) *createAccounthandler {
	return &createAccounthandler{
		Channel: channel,
		Usecase: usecase,
	}
}

func (h *createAccounthandler) Create() error {
	var data struct {
		Id string `json:"id,omitempty"`
	}

	var account entities.Account

	for payload := range h.Channel {
		err := json.Unmarshal(payload, &data)
		if err != nil {
			return err
		}

		account.Id = data.Id

		err = h.Usecase.Create(&account)
		if err != nil {
			return err
		}

	}

	return nil
}
