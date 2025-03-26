package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/courses"
	_modulesResponse "github.com/amdrx480/go-lms/controllers/modules/response"
	"gorm.io/gorm"
)

type Course struct {
	ID          int                       `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt           `json:"deleted_at,omitempty"` // gunakan pointer dan omitempty
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	CategoryID  int                       `json:"category_id"`
	Cover       string                    `json:"cover"`
	Instructor  string                    `json:"instructor"`
	Modules     []_modulesResponse.Module `json:"modules,omitempty"` // Gunakan pointer & omitempty
}

func FromDomain(domain courses.Domain) Course {
	return Course{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		Title:       domain.Title,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
		Cover:       domain.Cover,
		Instructor:  domain.Instructor,
		Modules:     _modulesResponse.FromDomainList(domain.Modules),
	}
}

func FromDomainList(coursesData []courses.Domain) []Course {
	var courses []Course
	for _, course := range coursesData {
		courses = append(courses, FromDomain(course))
	}
	return courses
}
