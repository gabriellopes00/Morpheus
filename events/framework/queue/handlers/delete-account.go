package handlers

import (
	"encoding/json"
	"events/domain/usecases"
)

type deleteAccountHandler struct {
	Channel <-chan []byte
	Usecase usecases.DeleteAccount
}

func NewDeleteAccountHandler(channel <-chan []byte, usecase usecases.DeleteAccount) *deleteAccountHandler {
	return &deleteAccountHandler{
		Channel: channel,
		Usecase: usecase,
	}
}

func (h *deleteAccountHandler) Delete() error {
	var data struct {
		Id string `json:"id,omitempty"`
	}

	for payload := range h.Channel {
		err := json.Unmarshal(payload, &data)
		if err != nil {
			return err
		}

		err = h.Usecase.Delete(data.Id)
		if err != nil {
			return err
		}

	}

	return nil
}
