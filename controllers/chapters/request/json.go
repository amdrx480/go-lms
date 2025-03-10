package request

import (
	"github.com/amdrx480/go-lms/businesses/chapters"
	"github.com/go-playground/validator/v10"
)

type Chapter struct {
	ModuleID int    `json:"module_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
}

func (req *Chapter) ToDomain() *chapters.Domain {
	return &chapters.Domain{
		ModuleID: req.ModuleID,
		Title:    req.Title,
	}
}

func (req *Chapter) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
