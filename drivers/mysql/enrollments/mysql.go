package enrollments

import (
	"context"
	"errors"
	"time"

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

// func (er *enrollmentRepository) CreateEnrollmentCourse(ctx context.Context, enrollmentDomain *enrollments.Domain) (enrollments.Domain, error) {
// 	record := FromDomain(enrollmentDomain)

// 	if err := er.conn.WithContext(ctx).Create(&record).Error; err != nil {
// 		return enrollments.Domain{}, err
// 	}

// 	return record.ToDomain(), nil
// }

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

func (er *enrollmentRepository) GetAllEnrollmentCourseByUserID(ctx context.Context, userID int) ([]enrollments.Domain, error) {
	var enrollmentRecords []Enrollment

	if err := er.conn.WithContext(ctx).
		Preload("Course"). // Menampilkan data course dalam enrollment
		Find(&enrollmentRecords, "user_id = ?", userID).Error; err != nil {
		return []enrollments.Domain{}, err
	}

	return ToDomainList(enrollmentRecords), nil
}

// func (er *enrollmentRepository) GetAll(ctx context.Context) ([]enrollments.Domain, error) {
// 	var records []Enrollment

// 	if err := er.conn.WithContext(ctx).Find(&records).Error; err != nil {
// 		return nil, err
// 	}

// 	enrollments := []enrollments.Domain{}

// 	for _, enrollment := range records {
// 		enrollments = append(enrollments, enrollment.ToDomain())
// 	}

// 	return enrollments, nil
// }

// func (er *enrollmentRepository) GetByID(ctx context.Context, id int) (enrollments.Domain, error) {
// 	var enrollment Enrollment

// 	if err := er.conn.WithContext(ctx).First(&enrollment, "user_id = ?", id).Error; err != nil {
// 		return enrollments.Domain{}, err
// 	}

// 	return enrollment.ToDomain(), nil
// }

// func (er *enrollmentRepository) Update(ctx context.Context, enrollmentDomain *enrollments.Domain, id int) (enrollments.Domain, error) {
// 	enrollment, err := er.GetByID(ctx, id)
// 	if err != nil {
// 		return enrollments.Domain{}, err
// 	}

// 	updatedCategory := FromDomain(&enrollment)

// 	updatedCategory.Title = enrollmentDomain.Title
// 	updatedCategory.Slug = enrollmentDomain.Slug // Pastikan slug ikut diperbarui

// 	if err := er.conn.WithContext(ctx).Save(&updatedCategory).Error; err != nil {
// 		return enrollments.Domain{}, err
// 	}

// 	return updatedCategory.ToDomain(), nil
// }

// func (er *enrollmentRepository) Delete(ctx context.Context, id int) error {
// 	enrollment, err := er.GetByID(ctx, id)

// 	if err != nil {
// 		return err
// 	}

// 	deletedCategory := FromDomain(&enrollment)

// 	if err := er.conn.WithContext(ctx).Unscoped().Delete(&deletedCategory).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }
