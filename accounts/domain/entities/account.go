package entities

import "time"

type Account struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
