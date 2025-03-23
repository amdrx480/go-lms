package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/enrollments"
	_courseResponse "github.com/amdrx480/go-lms/controllers/courses/response"

	"gorm.io/gorm"
)

type Enrollment struct {
	ID        int                    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	DeletedAt gorm.DeletedAt         `json:"deleted_at" gorm:"index"`
	UserID    int                    `json:"user_id"`
	Progress  int                    `json:"progress"`
	CourseID  int                    `json:"course_id"`
	Courses   _courseResponse.Course `json:"courses,omitempty"`
}

func FromDomain(domain enrollments.Domain) Enrollment {
	return Enrollment{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UserID:    domain.UserID,
		Progress:  domain.Progress,
		CourseID:  domain.CourseID,
		Courses:   _courseResponse.FromDomain(domain.Course),
	}
}

func FromDomainList(categoriesData []enrollments.Domain) []Enrollment {
	var enrollments []Enrollment
	for _, category := range categoriesData {
		enrollments = append(enrollments, FromDomain(category))
	}
	return enrollments
}

// func FromDomainWithCourse(domain enrollments.Domain) Enrollment {
// 	// Menggunakan FromDomainList untuk mengonversi modules
// 	modules := response.FromDomainList(domain.CourseID)
// 	return Enrollment{
// 		ID:           domain.ID,
// 		CreatedAt:    domain.CreatedAt,
// 		UpdatedAt:    domain.UpdatedAt,
// 		DeletedAt:    domain.DeletedAt,
// 		Title:        domain.Title,
// 		Description:  domain.Description,
// 		CategoryID:   domain.CategoryID,
// 		CategoryName: domain.CategoryTitle,
// 		Cover:        domain.Cover,
// 		Instructor:   domain.Instructor,
// 		Modules:      modules, // Hasil dari FromDomainList
// 	}
// }

// func FromDomainWithCourseList(domains []enrollments.Domain) []Enrollment {
// 	var enrollments []Enrollment
// 	for _, domain := range domains {
// 		enrollments = append(enrollments, FromDomainWithCourse(domain))
// 	}
// 	return enrollments
// }
