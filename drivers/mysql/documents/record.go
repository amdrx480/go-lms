package documents

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/documents"
	"gorm.io/gorm"
)

type Document struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	LessonID  int             `json:"lesson_id" gorm:"index;not null"`
	Title     string          `json:"title"`
	FileName  string          `json:"file_name"`
	FilePath  string          `json:"file_path"`
}

func (rec *Document) ToDomain() documents.Domain {
	return documents.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		LessonID:  rec.LessonID,
		Title:     rec.Title,
		FileName:  rec.FileName,
		FilePath:  rec.FilePath,
	}
}

func ToDomainList(records []Document) []documents.Domain {
	var domains []documents.Domain
	for _, rec := range records {
		domains = append(domains, rec.ToDomain())
	}
	return domains
}

func FromDomain(domain *documents.Domain) *Document {
	return &Document{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		LessonID:  domain.LessonID,
		Title:     domain.Title,
		FileName:  domain.FileName,
		FilePath:  domain.FilePath,
	}
}
