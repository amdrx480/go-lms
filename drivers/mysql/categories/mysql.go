package categories

import (
	"context"

	"github.com/amdrx480/go-lms/businesses/categories"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewMySQLCategoryRepository(conn *gorm.DB) categories.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	var records []Category

	if err := cr.conn.WithContext(ctx).Find(&records).Error; err != nil {
		return nil, err
	}

	categories := []categories.Domain{}

	for _, category := range records {
		categories = append(categories, category.ToDomain())
	}

	return categories, nil
}

func (cr *categoryRepository) GetByID(ctx context.Context, id int) (categories.Domain, error) {
	var category Category

	if err := cr.conn.WithContext(ctx).First(&category, "id = ?", id).Error; err != nil {
		return categories.Domain{}, err
	}

	return category.ToDomain(), nil
}

func (cr *categoryRepository) Create(ctx context.Context, categoryDomain *categories.Domain) (categories.Domain, error) {
	record := FromDomain(categoryDomain)

	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return categories.Domain{}, err
	}

	if err := result.Last(&record).Error; err != nil {
		return categories.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *categoryRepository) Update(ctx context.Context, categoryDomain *categories.Domain, id int) (categories.Domain, error) {
	category, err := cr.GetByID(ctx, id)

	if err != nil {
		return categories.Domain{}, err
	}

	updatedCategory := FromDomain(&category)

	updatedCategory.Name = categoryDomain.Name

	if err := cr.conn.WithContext(ctx).Save(&updatedCategory).Error; err != nil {
		return categories.Domain{}, err
	}

	return updatedCategory.ToDomain(), nil
}

func (cr *categoryRepository) Delete(ctx context.Context, id int) error {
	category, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedCategory := FromDomain(&category)

	if err := cr.conn.WithContext(ctx).Unscoped().Delete(&deletedCategory).Error; err != nil {
		return err
	}

	return nil
}
