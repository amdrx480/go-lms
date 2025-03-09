package modules

import (
	"github.com/amdrx480/go-lms/businesses/modules"

	"time"

	"gorm.io/gorm"
)

type Module struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CourseID  int             `json:"course_id" gorm:"index;not null"`
	Title     string          `json:"title"`
}

func (rec *Module) ToDomain() modules.Domain {
	return modules.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		CourseID:  rec.CourseID,
		Title:     rec.Title,
	}
}

func ToDomainList(records []Module) []modules.Domain {
	var domains []modules.Domain
	for _, rec := range records {
		domains = append(domains, rec.ToDomain())
	}
	return domains
}

func FromDomain(domain *modules.Domain) *Module {
	return &Module{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		CourseID:  domain.CourseID,
		Title:     domain.Title,
	}
}
