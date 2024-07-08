package models

import "github.com/google/uuid"

type ErrorMessage struct {
	Message string `json:"error"`
}

type LoginRequestDTO struct {
	Login             string `json:"login" validate:"required"` // email/phone
	Password          string `json:"password" validate:"required"`
	UniqueInformation map[string]string
}

type RegisterRequestDTO struct {
	Name    string `json:"name" validate:"required"`
	Surname string `json:"surname"`

	Email    string `json:"email"`
	Phone    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`

	UniqueInformation map[string]string
}

type LoginResponseDTO struct {
	UserID    uuid.UUID `json:"user_id"`
	IsAdmin   bool      `json:"is_admin"`
	IsSeller  bool      `json:"is_seller"`
	SessionId uuid.UUID `json:"session_id"`
}
