package models

import "github.com/go-playground/validator/v10"

type CategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

func (c *CategoryRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(c)

	return err
}
