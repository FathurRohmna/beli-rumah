package web

import "github.com/google/uuid"

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

type RegisterUserRequest struct {
	Name     string `json:"full_name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"example_password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	Password string `json:"password" validate:"required,min=8" example:"example_password"`
}
