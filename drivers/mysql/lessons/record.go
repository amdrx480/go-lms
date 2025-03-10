package lessons

import (
	"github.com/amdrx480/go-lms/businesses/lessons"

	"time"

	"gorm.io/gorm"
)

type Lesson struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ChapterID int             `json:"chapter_id" gorm:"index;not null"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	VideoURL  string          `json:"video_url"`
}

func (rec *Lesson) ToDomain() lessons.Domain {
	return lessons.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		ChapterID: rec.ChapterID,
		Title:     rec.Title,
		Content:   rec.Content,
		VideoURL:  rec.VideoURL,
	}
}

func ToDomainList(records []Lesson) []lessons.Domain {
	var domains []lessons.Domain
	for _, rec := range records {
		domains = append(domains, rec.ToDomain())
	}
	return domains
}

func FromDomain(domain *lessons.Domain) *Lesson {
	return &Lesson{
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
