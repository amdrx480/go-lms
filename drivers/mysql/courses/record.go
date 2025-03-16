package courses

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/courses"
	"github.com/amdrx480/go-lms/businesses/modules"
	"github.com/amdrx480/go-lms/drivers/mysql/categories"
	_modulesDB "github.com/amdrx480/go-lms/drivers/mysql/modules"

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
	Modules     []_modulesDB.Module `json:"modules" gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE"`
}

func (rec *Course) ToDomain() courses.Domain {
	modulesDomain := []modules.Domain{}

	// Iterasi rec.Module untuk memanggil metode ToDomain pada setiap elemen
	for _, module := range rec.Modules {
		modulesDomain = append(modulesDomain, module.ToDomain())
	}

	return courses.Domain{
		ID:            rec.ID,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
		DeletedAt:     rec.DeletedAt,
		Title:         rec.Title,
		Description:   rec.Description,
		CategoryTitle: rec.Category.Title,
		CategoryID:    rec.Category.ID,
		Cover:         rec.Cover,
		Instructor:    rec.Instructor,
		Modules:       modulesDomain,
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
