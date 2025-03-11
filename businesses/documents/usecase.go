package documents

import (
	"context"
	"log"
)

type documentUseCase struct {
	documentRepository Repository
}

func NewDocumentUseCase(repository Repository) UseCase {
	return &documentUseCase{
		documentRepository: repository,
	}
}

func (usecase documentUseCase) Create(ctx context.Context, documentReq *Domain) (Domain, error) {
	return usecase.documentRepository.Create(ctx, documentReq)
}

func (usecase documentUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	log.Println("[INFO] Starting GetAll use case")

	documents, err := usecase.documentRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[ERROR] Failed to get documents: %v\n", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully retrieved documents: %v\n", documents)
	return documents, nil
}

func (usecase documentUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	return usecase.documentRepository.GetByID(ctx, id)
}

func (usecase documentUseCase) Update(ctx context.Context, documentReq *Domain, id int) (Domain, error) {
	return usecase.documentRepository.Update(ctx, documentReq, id)
}

func (usecase documentUseCase) Delete(ctx context.Context, id int) error {
	return usecase.documentRepository.Delete(ctx, id)
}
