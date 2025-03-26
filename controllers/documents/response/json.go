package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/documents"

	"gorm.io/gorm"
)

type Document struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
	LessonID  int             `json:"chapter_id"`
	Title     string          `json:"title"`
	FileName  string          `json:"file_name"`
	FilePath  string          `json:"file_path"`
}

func FromDomain(domain documents.Domain) *Document {
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

func FromDomainList(documentsData []documents.Domain) []Document {
	var document []Document
	for _, course := range documentsData {
		document = append(document, *FromDomain(course))
	}
	return document
}
