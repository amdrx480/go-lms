package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/courses"

	"gorm.io/gorm"
)

type Course struct {
	ID           int             `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    *gorm.DeletedAt `json:"deleted_at,omitempty"` // gunakan pointer dan omitempty
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	CategoryName string          `json:"category_name"`
	CategoryID   int             `json:"category_id"`
	Cover        string          `json:"cover"`
	Instructor   string          `json:"instructor"`
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
