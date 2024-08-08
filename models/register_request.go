package models

import "github.com/go-playground/validator/v10"

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *RegisterRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(r)

	return err
}
