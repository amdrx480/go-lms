package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/modules"
	_chaptersResponse "github.com/amdrx480/go-lms/controllers/chapters/response"

	"gorm.io/gorm"
)

type Module struct {
	ID        int                         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
	DeletedAt *gorm.DeletedAt             `json:"deleted_at,omitempty"`
	CourseID  int                         `json:"course_id"`
	Title     string                      `json:"title"`
	Chapters  []_chaptersResponse.Chapter `json:"chapters,omitempty"`
}

func FromDomain(domain modules.Domain) *Module {
	return &Module{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		CourseID:  domain.CourseID,
		Title:     domain.Title,
		Chapters:  _chaptersResponse.FromDomainList(domain.Chapters),
	}
}

func FromDomainList(modulesData []modules.Domain) []Module {
	var module []Module
	for _, course := range modulesData {
		module = append(module, *FromDomain(course))
	}
	return module
}
