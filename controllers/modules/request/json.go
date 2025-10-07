package request

import (
	"github.com/amdrx480/angsana-boga/businesses/modules"
	"github.com/go-playground/validator/v10"
)

type Module struct {
	CourseID int    `json:"course_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
}

func (req *Module) ToDomain() *modules.Domain {
	return &modules.Domain{
		CourseID: req.CourseID,
		Title:    req.Title,
	}
}

func (req *Module) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
