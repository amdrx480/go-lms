package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/lessons"
	"github.com/amdrx480/go-lms/controllers/documents/response"

	"gorm.io/gorm"
)

type Lesson struct {
	ID        int                 `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt *gorm.DeletedAt     `json:"deleted_at,omitempty"`
	ChapterID int                 `json:"chapter_id"`
	Title     string              `json:"title"`
	Content   string              `json:"content"`
	VideoURL  string              `json:"video_url"`
	Documents []response.Document `json:"documents,omitempty"`
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

func FromDomainWithDocument(domain lessons.Domain) Lesson {
	// Menggunakan FromDomainList untuk mengonversi modules
	documents := response.FromDomainList(domain.Documents)
	return Lesson{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ChapterID: domain.ChapterID,
		Title:     domain.Title,
		Content:   domain.Content,
		VideoURL:  domain.VideoURL,
		Documents: documents, // Hasil dari FromDomainList
	}
}

func FromDomainWithDocumentList(domains []lessons.Domain) []Lesson {
	var lessons []Lesson
	for _, domain := range domains {
		lessons = append(lessons, FromDomainWithDocument(domain))
	}
	return lessons
}
