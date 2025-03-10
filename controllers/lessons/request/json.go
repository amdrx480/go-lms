package request

import (
	"github.com/amdrx480/go-lms/businesses/lessons"
	"github.com/go-playground/validator/v10"
)

type Lesson struct {
	ChapterID int    `json:"chapter_id" validate:"required"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	VideoURL  string `json:"video_url" validate:"required"`
}

func (req *Lesson) ToDomain() *lessons.Domain {
	return &lessons.Domain{
		ChapterID: req.ChapterID,
		Title:     req.Title,
		Content:   req.Content,
		VideoURL:  req.VideoURL,
	}
}

func (req *Lesson) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
