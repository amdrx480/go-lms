package lessons

import (
	"context"
	"log"
)

type lessonUseCase struct {
	lessonRepository Repository
}

func NewLessonUseCase(repository Repository) UseCase {
	return &lessonUseCase{
		lessonRepository: repository,
	}
}

func (usecase lessonUseCase) Create(ctx context.Context, lessonReq *Domain) (Domain, error) {
	return usecase.lessonRepository.Create(ctx, lessonReq)
}

func (usecase lessonUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	log.Println("[INFO] Starting GetAll use case")

	lessons, err := usecase.lessonRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[ERROR] Failed to get lessons: %v\n", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully retrieved lessons: %v\n", lessons)
	return lessons, nil
}

func (usecase lessonUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	return usecase.lessonRepository.GetByID(ctx, id)
}

func (usecase lessonUseCase) Update(ctx context.Context, lessonReq *Domain, id int) (Domain, error) {
	return usecase.lessonRepository.Update(ctx, lessonReq, id)
}

func (usecase lessonUseCase) Delete(ctx context.Context, id int) error {
	return usecase.lessonRepository.Delete(ctx, id)
}
