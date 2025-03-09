package modules

import (
	"context"
	"log"
)

type moduleUseCase struct {
	moduleRepository Repository
}

func NewModuleUseCase(repository Repository) UseCase {
	return &moduleUseCase{
		moduleRepository: repository,
	}
}

func (usecase moduleUseCase) Create(ctx context.Context, moduleReq *Domain) (Domain, error) {
	return usecase.moduleRepository.Create(ctx, moduleReq)
}

func (usecase moduleUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	log.Println("[INFO] Starting GetAll use case")

	modules, err := usecase.moduleRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[ERROR] Failed to get modules: %v\n", err)
		return nil, err
	}

	log.Printf("[INFO] Successfully retrieved modules: %v\n", modules)
	return modules, nil
}

func (usecase moduleUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	return usecase.moduleRepository.GetByID(ctx, id)
}

func (usecase moduleUseCase) Update(ctx context.Context, moduleReq *Domain, id int) (Domain, error) {
	return usecase.moduleRepository.Update(ctx, moduleReq, id)
}

func (usecase moduleUseCase) Delete(ctx context.Context, id int) error {
	return usecase.moduleRepository.Delete(ctx, id)
}
