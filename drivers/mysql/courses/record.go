package courses

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/courses"
	_categoriesDB "github.com/amdrx480/go-lms/drivers/mysql/categories"
	_modulesDB "github.com/amdrx480/go-lms/drivers/mysql/modules"

	"gorm.io/gorm"
)

// record layer
type Course struct {
	ID          int                    `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt        `json:"deleted_at" gorm:"index"`
	Slug        string                 `json:"slug"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	CategoryID  int                    `json:"category_id"`
	Category    _categoriesDB.Category `json:"category"`
	Cover       string                 `json:"cover"`
	Instructor  string                 `json:"instructor"`
	Modules     []_modulesDB.Module    `json:"modules"`
}

func (rec *Course) ToDomain() courses.Domain {
	return courses.Domain{
		ID:          rec.ID,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
		DeletedAt:   rec.DeletedAt,
		Slug:        rec.Slug,
		Title:       rec.Title,
		Description: rec.Description,
		CategoryID:  rec.Category.ID,
		Category:    rec.Category.ToDomain(),
		Cover:       rec.Cover,
		Instructor:  rec.Instructor,
		Modules:     _modulesDB.ToDomainList(rec.Modules),
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
		Slug:        domain.Slug,
		Title:       domain.Title,
		Description: domain.Description,
		CategoryID:  domain.CategoryID,
		Category:    *_categoriesDB.FromDomain(&domain.Category),
		Cover:       domain.Cover,
		Instructor:  domain.Instructor,
	}
}
