package chapters

import (
	"context"
	"log"
)

type chapterUseCase struct {
	chapterRepository Repository
}

func NewChapterUseCase(repository Repository) UseCase {
	return &chapterUseCase{
		chapterRepository: repository,
	}
}

func (usecase chapterUseCase) Create(ctx context.Context, chapterReq *Domain) (Domain, error) {
	return usecase.chapterRepository.Create(ctx, chapterReq)
}

func (usecase chapterUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	log.Println("[INFO] Starting GetAll use case")

	chapters, err := usecase.chapterRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[ERROR] Failed to get chapters: %v\n", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully retrieved chapters: %v\n", chapters)
	return chapters, nil
}

func (usecase chapterUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	return usecase.chapterRepository.GetByID(ctx, id)
}

func (usecase chapterUseCase) Update(ctx context.Context, chapterReq *Domain, id int) (Domain, error) {
	return usecase.chapterRepository.Update(ctx, chapterReq, id)
}

func (usecase chapterUseCase) Delete(ctx context.Context, id int) error {
	return usecase.chapterRepository.Delete(ctx, id)
}
