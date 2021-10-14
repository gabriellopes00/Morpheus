package handlers

import (
	"encoding/json"
	"tickets/domain/usecases"
	"time"
)

type createEventHandler struct {
	Channel <-chan []byte
	Usecase usecases.CreateEvent
}

func NewCreateEventHandler(channel <-chan []byte, usecase usecases.CreateEvent) *createEventHandler {
	return &createEventHandler{
		Channel: channel,
		Usecase: usecase,
	}
}

func (h *createEventHandler) Create() error {
	var data struct {
		Id   string `json:"id,omitempty"`
		Date string `json:"date,omitempty"`
	}

	for payload := range h.Channel {
		if err := json.Unmarshal(payload, &data); err != nil {
			return err
		}

		dateParsed, err := time.Parse(time.RFC3339, data.Date)
		if err != nil {
			return err
		}

		if err = h.Usecase.Create(data.Id, dateParsed); err != nil {
			return err
		}

	}

	return nil
}
