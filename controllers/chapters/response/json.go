package response

import (
	"time"

	"github.com/amdrx480/angsana-boga/businesses/chapters"
	_lessonsResponse "github.com/amdrx480/angsana-boga/controllers/lessons/response"
	"gorm.io/gorm"
)

type Chapter struct {
	ID        int                       `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
	DeletedAt *gorm.DeletedAt           `json:"deleted_at,omitempty"`
	ModuleID  int                       `json:"module_id"`
	Title     string                    `json:"title"`
	Lessons   []_lessonsResponse.Lesson `json:"lessons,omitempty"`
}

func FromDomain(domain chapters.Domain) *Chapter {
	return &Chapter{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		ModuleID:  domain.ModuleID,
		Title:     domain.Title,
		Lessons:   _lessonsResponse.FromDomainList(domain.Lessons),
	}
}

func FromDomainList(chaptersData []chapters.Domain) []Chapter {
	var chapter []Chapter
	for _, course := range chaptersData {
		chapter = append(chapter, *FromDomain(course))
	}
	return chapter
}
