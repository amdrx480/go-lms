package enrollments

import (
	"context"
)

type enrollmentUseCase struct {
	enrollmentRepository Repository
}

func NewEnrollmentUseCase(repository Repository) Usecase {
	return &enrollmentUseCase{
		enrollmentRepository: repository,
	}
}

func (usecase *enrollmentUseCase) CreateEnrollmentCourse(ctx context.Context, enrollmentReq *Domain) (Domain, error) {
	return usecase.enrollmentRepository.CreateEnrollmentCourse(ctx, enrollmentReq)
}

func (usecase *enrollmentUseCase) GetEnrollmentByUserCourse(ctx context.Context, userID int, courseID int) (Domain, error) {
	return usecase.enrollmentRepository.GetEnrollmentByUserCourse(ctx, userID, courseID)
}

func (usecase *enrollmentUseCase) GetAllEnrollmentCourseByUserID(ctx context.Context, userID int) ([]Domain, error) {
	return usecase.enrollmentRepository.GetAllEnrollmentCourseByUserID(ctx, userID)
}

// func (usecase *enrollmentUseCase) GetAllUserEnrollment(ctx context.Context) ([]Domain, error) {
// 	return usecase.enrollmentRepository.GetAllUserEnrollment(ctx)
// }

// func (usecase *enrollmentUseCase) CheckEnrollmentStatus(ctx context.Context, id int) (Domain, error) {
// 	return usecase.enrollmentRepository.GetByID(ctx, id)
// }

// func (usecase *enrollmentUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
// 	return usecase.enrollmentRepository.GetByID(ctx, id)
// }

// func (usecase *enrollmentUseCase) Update(ctx context.Context, enrollmentReq *Domain, id int) (Domain, error) {
// 	return usecase.enrollmentRepository.Update(ctx, enrollmentReq, id)
// }

// func (usecase *enrollmentUseCase) Delete(ctx context.Context, id int) error {
// 	return usecase.enrollmentRepository.Delete(ctx, id)
// }
