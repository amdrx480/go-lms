package request

import (
	"github.com/amdrx480/angsana-boga/businesses/categories"
	"github.com/go-playground/validator/v10"
)

type Category struct {
	Title string `json:"title" validate:"required"`
}

func (req *Category) ToDomain() *categories.Domain {
	return &categories.Domain{
		Title: req.Title,
	}
}

func (req *Category) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
