package request

import (
	"github.com/amdrx480/angsana-boga/businesses/enrollments"
	"github.com/go-playground/validator/v10"
)

type Enrollment struct {
	UserID   int `json:"user_id" validate:"required"`
	CourseID int `json:"course_id" validate:"required"`
}

func (req *Enrollment) ToDomain() *enrollments.Domain {
	return &enrollments.Domain{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}
}

func (req *Enrollment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
