package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/lessons"

	"gorm.io/gorm"
)

type Lesson struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
	ChapterID int             `json:"chapter_id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	VideoURL  string          `json:"video_url"`
}

func FromDomain(domain lessons.Domain) Lesson {
	return Lesson{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ChapterID: domain.ChapterID,
		Title:     domain.Title,
		Content:   domain.Content,
		VideoURL:  domain.VideoURL,
	}
}

func FromDomainList(chaptersData []lessons.Domain) []Lesson {
	var chapter []Lesson
	for _, course := range chaptersData {
		chapter = append(chapter, FromDomain(course))
	}
	return chapter
}
