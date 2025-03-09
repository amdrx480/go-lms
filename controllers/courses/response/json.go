package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/courses"
	"github.com/amdrx480/go-lms/controllers/modules/response"
	"gorm.io/gorm"
)

type Course struct {
	ID           int               `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time         `json:"created_at"`
	UpdatedAt    time.Time         `json:"updated_at"`
	DeletedAt    *gorm.DeletedAt   `json:"deleted_at,omitempty"` // gunakan pointer dan omitempty
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	CategoryID   int               `json:"category_id"`
	CategoryName string            `json:"category_name"`
	Cover        string            `json:"cover"`
	Instructor   string            `json:"instructor"`
	Modules      []response.Module `json:"modules,omitempty"`
}

func FromDomain(domain courses.Domain) Course {
	return Course{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		Title:        domain.Title,
		Description:  domain.Description,
		CategoryID:   domain.CategoryID,
		CategoryName: domain.CategoryName,
		Cover:        domain.Cover,
		Instructor:   domain.Instructor,
	}
}

func FromDomainList(coursesData []courses.Domain) []Course {
	var courses []Course
	for _, course := range coursesData {
		courses = append(courses, FromDomain(course))
	}
	return courses
}

func FromDomainWithModules(domain courses.Domain) Course {
	// Menggunakan FromDomainList untuk mengonversi modules
	modules := response.FromDomainList(domain.Modules)
	return Course{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		Title:        domain.Title,
		Description:  domain.Description,
		CategoryID:   domain.CategoryID,
		CategoryName: domain.CategoryName,
		Cover:        domain.Cover,
		Instructor:   domain.Instructor,
		Modules:      modules, // Hasil dari FromDomainList
	}
}

func FromDomainWithModulesList(domains []courses.Domain) []Course {
	var courses []Course
	for _, domain := range domains {
		courses = append(courses, FromDomainWithModules(domain))
	}
	return courses
}
