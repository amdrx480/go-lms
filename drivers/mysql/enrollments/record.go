package enrollments

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/enrollments"
	"github.com/amdrx480/go-lms/drivers/mysql/courses"
	"github.com/amdrx480/go-lms/drivers/mysql/users"

	"gorm.io/gorm"
)

type Enrollment struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User      users.User     `json:"user"`
	UserID    int            `json:"user_id"`
	Course    courses.Course `json:"course"`
	CourseID  int            `json:"course_id"`
	Progress  int            `gorm:"default:0" json:"progress"` // 0 - 100
}

func (rec *Enrollment) ToDomain() enrollments.Domain {
	return enrollments.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		UserID:    rec.UserID,
		CourseID:  rec.CourseID,
		Course:    rec.Course.ToDomain(),
		Progress:  rec.Progress,
	}
}

func ToDomainList(records []Enrollment) []enrollments.Domain {
	var domains []enrollments.Domain
	for _, rec := range records {
		domains = append(domains, rec.ToDomain())
	}
	return domains
}

func FromDomain(domain *enrollments.Domain) *Enrollment {
	return &Enrollment{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		UserID:    domain.UserID,
		// Course:    courses.FromDomain(domain.Course),
		Course:   *courses.FromDomain(&domain.Course),
		CourseID: domain.CourseID,
		Progress: domain.Progress,
	}
}
