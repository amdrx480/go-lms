package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/chapters"
	"github.com/amdrx480/go-lms/controllers/lessons/response"

	"gorm.io/gorm"
)

type Chapter struct {
	ID        int               `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt *gorm.DeletedAt   `json:"deleted_at,omitempty"`
	ModuleID  int               `json:"module_id"`
	Title     string            `json:"title"`
	Lessons   []response.Lesson `json:"chapters,omitempty"`
}

func FromDomain(domain chapters.Domain) Chapter {
	return Chapter{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ModuleID:  domain.ModuleID,
		Title:     domain.Title,
	}
}

func FromDomainList(chaptersData []chapters.Domain) []Chapter {
	var chapter []Chapter
	for _, course := range chaptersData {
		chapter = append(chapter, FromDomain(course))
	}
	return chapter
}

func FromDomainWithLessons(domain chapters.Domain) Chapter {
	// Menggunakan FromDomainList untuk mengonversi modules
	lessons := response.FromDomainList(domain.Lessons)
	return Chapter{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ModuleID:  domain.ModuleID,
		Title:     domain.Title,
		Lessons:   lessons, // Hasil dari FromDomainList
	}
}

func FromDomainWithChapterList(domains []chapters.Domain) []Chapter {
	var chapters []Chapter
	for _, domain := range domains {
		chapters = append(chapters, FromDomainWithLessons(domain))
	}
	return chapters
}
