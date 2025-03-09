package request

import (
	"github.com/amdrx480/go-lms/businesses/courses"

	"github.com/go-playground/validator/v10"
)

type Course struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Cover       string `json:"cover" validate:"required"`
	Instructor  string `json:"instructor" validate:"required"`
}

func (req *Course) ToDomain() *courses.Domain {
	return &courses.Domain{
		Title:       req.Title,
		Description: req.Description,
		CategoryID:  req.CategoryID,
		Cover:       req.Cover,
		Instructor:  req.Instructor,
	}
}

func (req *Course) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
