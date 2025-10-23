package models

import (
	"time"
)

type Email struct {
	ID          uint      `json:"id"`
	Email       string    `json:"email"`
	IsConfirmed uint8     `json:"is_confirmed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`
}

type SubscribeRequest struct {
	Email       string `json:"email"`
	ClientToken string `json:"clientToken"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
