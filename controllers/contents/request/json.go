package request

import (
	"mini-cms-api/businesses/contents"

	"github.com/go-playground/validator/v10"
)

type Content struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

func (req *Content) ToDomain() *contents.Domain {
	return &contents.Domain{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
	}
}

func (req *Content) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
