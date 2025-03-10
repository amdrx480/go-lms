package chapters

import (
	"github.com/amdrx480/go-lms/businesses/chapters"

	"time"

	"gorm.io/gorm"
)

type Chapter struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	ModuleID  int             `json:"course_id" gorm:"index;not null"`
	Title     string          `json:"title"`
}

func (rec *Chapter) ToDomain() chapters.Domain {
	return chapters.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		ModuleID:  rec.ModuleID,
		Title:     rec.Title,
	}
}

func ToDomainList(records []Chapter) []chapters.Domain {
	var domains []chapters.Domain
	for _, rec := range records {
		domains = append(domains, rec.ToDomain())
	}
	return domains
}

func FromDomain(domain *chapters.Domain) *Chapter {
	return &Chapter{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ModuleID:  domain.ModuleID,
		Title:     domain.Title,
	}
}
