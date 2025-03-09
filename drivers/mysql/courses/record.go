package courses

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/courses"
	"github.com/amdrx480/go-lms/drivers/mysql/categories"

	"gorm.io/gorm"
)

// record layer
type Course struct {
	ID          int                 `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt     `json:"deleted_at" gorm:"index"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Category    categories.Category `json:"category"`
	CategoryID  int                 `json:"category_id"`
	Cover       string              `json:"cover"`
	Instructor  string              `json:"instructor"`
}

func (rec *Course) ToDomain() courses.Domain {
	return courses.Domain{
		ID:           rec.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
		Title:        rec.Title,
		Description:  rec.Description,
		CategoryName: rec.Category.Name,
		CategoryID:   rec.Category.ID,
		Cover:        rec.Cover,
		Instructor:   rec.Instructor,
	}
}

func ToDomainList(records []Course) []courses.Domain {
	var domains []courses.Domain
	for _, rec := range records {
		domains = append(domains, rec.ToDomain())
	}
	return domains
}

func FromDomain(domain *courses.Domain) *Course {
	return &Course{
		ID:          domain.ID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
		Title:       domain.Title,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
		Cover:       domain.Cover,
		Instructor:  domain.Instructor,
	}
}
