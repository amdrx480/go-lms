package response

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
	Name      string         `json:"name"`
}

func FromDomain(domain categories.Domain) Category {
	return Category{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Name:      domain.Name,
	}
}

func FromDomainList(categoriesData []categories.Domain) []Category {
	var categories []Category
	for _, category := range categoriesData {
		categories = append(categories, FromDomain(category))
	}
	return categories
}
