package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/modules"

	"gorm.io/gorm"
)

type Module struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
	CourseID  int             `json:"course_id"`
	Title     string          `json:"title"`
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
