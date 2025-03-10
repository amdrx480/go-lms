package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/modules"
	"github.com/amdrx480/go-lms/controllers/chapters/response"

	"gorm.io/gorm"
)

type Module struct {
	ID        int                `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt *gorm.DeletedAt    `json:"deleted_at,omitempty"`
	CourseID  int                `json:"course_id"`
	Title     string             `json:"title"`
	Chapters  []response.Chapter `json:"chapters,omitempty"`
}

func FromDomain(domain modules.Domain) Module {
	return Module{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		CourseID:  domain.CourseID,
		Title:     domain.Title,
	}
}

func FromDomainList(modulesData []modules.Domain) []Module {
	var module []Module
	for _, course := range modulesData {
		module = append(module, FromDomain(course))
	}
	return module
}

func FromDomainWithChapter(domain modules.Domain) Module {
	// Menggunakan FromDomainList untuk mengonversi modules
	chapters := response.FromDomainList(domain.Chapters)
	return Module{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		CourseID:  domain.CourseID,
		Title:     domain.Title,
		Chapters:  chapters, // Hasil dari FromDomainList
	}
}

func FromDomainWithChapterList(domains []modules.Domain) []Module {
	var modules []Module
	for _, domain := range domains {
		modules = append(modules, FromDomainWithChapter(domain))
	}
	return modules
}
