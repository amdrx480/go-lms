package request

import (
	"github.com/amdrx480/go-lms/businesses/documents"
	"github.com/go-playground/validator/v10"
)

type Document struct {
	LessonID int    `json:"lesson_id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	FileName string `json:"file_name" validate:"required"`
	FilePath string `json:"file_path" validate:"required"`
}

func (req *Document) ToDomain() *documents.Domain {
	return &documents.Domain{
		LessonID: req.LessonID,
		Title:    req.Title,
		FileName: req.FileName,
		FilePath: req.FilePath,
	}
}

func (req *Document) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
