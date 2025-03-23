package enrollments

import (
	"context"
	"time"

	"github.com/amdrx480/go-lms/businesses/courses"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	UserID    int
	CourseID  int
	Progress  int
	Course    courses.Domain

	// LastAccessedAt time.Time
}

type Usecase interface {
	CreateEnrollmentCourse(ctx context.Context, enrollmentReq *Domain) (Domain, error)
	GetEnrollmentByUserCourse(ctx context.Context, userID int, courseID int) (Domain, error)

	GetAllEnrollmentCourseByUserID(ctx context.Context, userID int) ([]Domain, error)

	// GetAllUserEnrollment(ctx context.Context) ([]Domain, error)

	// CheckEnrollmentStatus(ctx context.Context, id int) (Domain, error)

	// GetByID(ctx context.Context, id int) (Domain, error)

	// UpdateLessonProgress(ctx context.Context, enrollmentReq *Domain, id int) (Domain, error)

	// Delete(ctx context.Context, id int) error
}

type Repository interface {
	CreateEnrollmentCourse(ctx context.Context, enrollmentReq *Domain) (Domain, error)
	GetEnrollmentByUserCourse(ctx context.Context, userID int, courseID int) (Domain, error)

	GetAllEnrollmentCourseByUserID(ctx context.Context, userID int) ([]Domain, error)

	// GetAllUserEnrollment(ctx context.Context) ([]Domain, error)

	// GetByID(ctx context.Context, id int) (Domain, error)
	// Update(ctx context.Context, enrollmentReq *Domain, id int) (Domain, error)
	// Delete(ctx context.Context, id int) error
}
