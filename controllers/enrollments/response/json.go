package response

import (
	"time"

	"github.com/amdrx480/angsana-boga/businesses/enrollments"
	_courseResponse "github.com/amdrx480/angsana-boga/controllers/courses/response"

	"gorm.io/gorm"
)

type Enrollment struct {
	ID        int                     `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
	DeletedAt gorm.DeletedAt          `json:"deleted_at" gorm:"index"`
	UserID    int                     `json:"user_id"`
	Progress  int                     `json:"progress"`
	CourseID  int                     `json:"course_id"`
	Courses   *_courseResponse.Course `json:"courses,omitempty"`
}

func FromDomain(domain enrollments.Domain) Enrollment {
	var courseResponse *_courseResponse.Course
	if domain.Course.ID != 0 { // Jika ID Course tidak kosong
		course := _courseResponse.FromDomain(domain.Course)
		courseResponse = &course // Assign ke pointer
	}

	return Enrollment{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UserID:    domain.UserID,
		Progress:  domain.Progress,
		CourseID:  domain.CourseID,
		Courses:   courseResponse,
	}
}

func FromDomainList(categoriesData []enrollments.Domain) []Enrollment {
	var enrollments []Enrollment
	for _, category := range categoriesData {
		enrollments = append(enrollments, FromDomain(category))
	}
	return enrollments
}
