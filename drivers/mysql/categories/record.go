package categories

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/categories"

	"gorm.io/gorm"
)

type Category struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Title     string         `json:"title"`
	Slug      string         `json:"Slug"`
}

func (rec *Category) ToDomain() categories.Domain {
	return categories.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Title:     rec.Title,
		Slug:      rec.Slug,
	}
}

// func ToDomainList(records []Category) []categories.Domain {
// 	var domains []categories.Domain
// 	for _, rec := range records {
// 		domains = append(domains, rec.ToDomain())
// 	}
// 	return domains
// }

func FromDomain(domain *categories.Domain) *Category {
	return &Category{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Title:     domain.Title,
		Slug:      domain.Slug,
	}
}
