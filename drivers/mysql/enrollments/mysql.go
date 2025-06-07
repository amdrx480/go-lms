package enrollments

import (
	"context"
	"errors"
	"time"

	"github.com/amdrx480/go-lms/app/middlewares"
	"github.com/amdrx480/go-lms/businesses/enrollments"

	"gorm.io/gorm"
)

type enrollmentRepository struct {
	conn *gorm.DB
}

func NewMySQLEnrollmentRepository(conn *gorm.DB) enrollments.Repository {
	return &enrollmentRepository{
		conn: conn,
	}
}

func (er *enrollmentRepository) CreateEnrollmentCourse(ctx context.Context, enrollmentDomain *enrollments.Domain) (enrollments.Domain, error) {
	var existingEnrollment Enrollment

	// Cek apakah user sudah enroll di course yang sama
	err := er.conn.WithContext(ctx).
		Where("user_id = ? AND course_id = ?", enrollmentDomain.UserID, enrollmentDomain.CourseID).
		First(&existingEnrollment).Error

	// Jika ditemukan, return error
	if err == nil {
		return enrollments.Domain{}, errors.New("user already enrolled in this course")
	}

	// Jika error bukan karena record tidak ditemukan, return error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return enrollments.Domain{}, err
	}

	// Jika tidak ditemukan, buat enrollment baru
	record := FromDomain(enrollmentDomain)
	if err := er.conn.WithContext(ctx).Create(&record).Error; err != nil {
		return enrollments.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (er *enrollmentRepository) GetEnrollmentByUserCourse(ctx context.Context, userID, courseID int) (enrollments.Domain, error) {
	var enrollment Enrollment

	if err := er.conn.WithContext(ctx).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		First(&enrollment).Error; err != nil {
		return enrollments.Domain{}, err
	}

	return enrollment.ToDomain(), nil
}

func (er *enrollmentRepository) UpdateProgress(ctx context.Context, userID, courseID, completedLessons int, totalLessons int) (Enrollment, error) {
	// Hitung progress
	progress := (completedLessons * 100) / totalLessons

	// Update progress di database
	var enrollment Enrollment
	if err := er.conn.WithContext(ctx).
		Where("user_id = ? AND course_id = ?", userID, courseID).
		First(&enrollment).Error; err != nil {
		return Enrollment{}, err
	}

	enrollment.Progress = progress
	enrollment.UpdatedAt = time.Now()

	if err := er.conn.WithContext(ctx).Save(&enrollment).Error; err != nil {
		return Enrollment{}, err
	}

	return enrollment, nil
}

func (er *enrollmentRepository) GetAllEnrollmentCourseByUserID(ctx context.Context) ([]enrollments.Domain, error) {
	id, err := middlewares.GetUserID(ctx)
	if err != nil {
		return []enrollments.Domain{}, errors.New("invalid token")
	}

	var enrollmentRecords []Enrollment

	if err := er.conn.WithContext(ctx).
		Preload("Course.Category").
		Preload("Course.Modules.Chapter.Lesson.Documents").
		Find(&enrollmentRecords, "user_id = ?", id).Error; err != nil {
		return []enrollments.Domain{}, err
	}

	return ToDomainList(enrollmentRecords), nil
}
