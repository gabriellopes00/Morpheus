package handlers

import (
	"encoding/json"
	"events/application"
	"events/domain/entities"
)

type soldOutEventHandler struct {
	Channel <-chan []byte
	Usecase application.UpdateEvent
}

func NewsoldOutEventHandler(channel <-chan []byte, usecase application.UpdateEvent) *soldOutEventHandler {
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

		err = h.Usecase.SetStatus(data.Id, entities.StatusSoldOut)
		if err != nil {
			return err
		}

	}

	return nil
}
