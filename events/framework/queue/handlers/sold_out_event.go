package handlers

import (
	"encoding/json"
	"events/domain/entities"
	"events/domain/usecases"
)

type soldOutEventHandler struct {
	Channel <-chan []byte
	Usecase usecases.UpdateEvents
}

func NewsoldOutEventHandler(channel <-chan []byte, usecase usecases.UpdateEvents) *soldOutEventHandler {
	return &soldOutEventHandler{
		Channel: channel,
		Usecase: usecase,
	}
}

func (h *soldOutEventHandler) Handle() error {
	var data struct {
		Id string `json:"id,omitempty"`
	}

	for payload := range h.Channel {
		err := json.Unmarshal(payload, &data)
		if err != nil {
			return err
		}

		err = h.Usecase.UpdateStatus(data.Id, entities.StatusSoldOut)
		if err != nil {
			return err
		}

	}

	return nil
}
