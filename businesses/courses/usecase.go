package courses

import (
	"context"
	"log"
)

type courseUsecase struct {
	courseRepository Repository
}

func NewCourseUsecase(repository Repository) UseCase {
	return &courseUsecase{
		courseRepository: repository,
	}
}

func (usecase *courseUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	log.Println("[INFO] Starting GetAll use case")
	courses, err := usecase.courseRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[ERROR] Failed to get courses: %v\n", err)
		return nil, err
	}
	log.Printf("[INFO] Successfully retrieved courses: %v\n", courses)

	return courses, nil
}

func (usecase *courseUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	return usecase.courseRepository.GetByID(ctx, id)
}

func (usecase *courseUsecase) Create(ctx context.Context, courseDomain *Domain) (Domain, error) {
	return usecase.courseRepository.Create(ctx, courseDomain)
}

func (usecase *courseUsecase) Update(ctx context.Context, courseDomain *Domain, id int) (Domain, error) {
	return usecase.courseRepository.Update(ctx, courseDomain, id)
}

func (usecase *courseUsecase) Delete(ctx context.Context, id int) error {
	return usecase.courseRepository.Delete(ctx, id)
}

func (usecase *courseUsecase) Restore(ctx context.Context, id int) (Domain, error) {
	return usecase.courseRepository.Restore(ctx, id)
}

func (usecase *courseUsecase) ForceDelete(ctx context.Context, id int) error {
	return usecase.courseRepository.ForceDelete(ctx, id)
}
