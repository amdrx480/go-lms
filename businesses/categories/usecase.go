package categories

import (
	"context"

	"github.com/amdrx480/go-lms/utils"
)

type categoryUseCase struct {
	categoryRepository Repository
}

func NewCategoryUseCase(repository Repository) Usecase {
	return &categoryUseCase{
		categoryRepository: repository,
	}
}

func (usecase *categoryUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	return usecase.categoryRepository.GetAll(ctx)
}

func (usecase *categoryUseCase) GetByID(ctx context.Context, id int) (Domain, error) {
	return usecase.categoryRepository.GetByID(ctx, id)
}

func (usecase *categoryUseCase) Create(ctx context.Context, categoryReq *Domain) (Domain, error) {
	categoryReq.Slug = utils.GenerateSlug(categoryReq.Name)
	return usecase.categoryRepository.Create(ctx, categoryReq)
}

func (usecase *categoryUseCase) Update(ctx context.Context, categoryReq *Domain, id int) (Domain, error) {
	existingCategory, err := usecase.categoryRepository.GetByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	// Jika Name berubah, regenerasi slug
	if categoryReq.Name != existingCategory.Name {
		categoryReq.Slug = utils.GenerateSlug(categoryReq.Name)
	} else {
		categoryReq.Slug = existingCategory.Slug // Gunakan slug lama jika nama sama
	}

	// Kirim data ke repository
	return usecase.categoryRepository.Update(ctx, categoryReq, id)
}

func (usecase *categoryUseCase) Delete(ctx context.Context, id int) error {
	return usecase.categoryRepository.Delete(ctx, id)
}
