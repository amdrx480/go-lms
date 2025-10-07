package modules

import (
	"context"

	"github.com/amdrx480/angsana-boga/businesses/modules"

	"gorm.io/gorm"
)

type moduleRepository struct {
	conn *gorm.DB
}

func NewMySQLModuleRepository(conn *gorm.DB) modules.Repository {
	return &moduleRepository{
		conn: conn,
	}
}

func (mr *moduleRepository) Create(ctx context.Context, moduleReq *modules.Domain) (modules.Domain, error) {
	record := FromDomain(moduleReq)
	result := mr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return modules.Domain{}, err
	}

	if err := result.WithContext(ctx).Last(&record).Error; err != nil {
		return modules.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (mr *moduleRepository) GetAll(ctx context.Context) ([]modules.Domain, error) {
	var moduleRecords []Module

	if err := mr.conn.WithContext(ctx).Preload("Chapter").Find(&moduleRecords).Error; err != nil {
		return nil, err
	}

	return ToDomainList(moduleRecords), nil
}

func (mr *moduleRepository) GetByID(ctx context.Context, id int) (modules.Domain, error) {
	var module Module

	if err := mr.conn.WithContext(ctx).First(&module, "id = ?", id).Error; err != nil {
		return modules.Domain{}, err
	}

	return module.ToDomain(), nil
}

func (mr *moduleRepository) Update(ctx context.Context, moduleReq *modules.Domain, id int) (modules.Domain, error) {
	module, err := mr.GetByID(ctx, id)

	if err != nil {
		return modules.Domain{}, err
	}

	updatedModule := FromDomain(&module)
	updatedModule.Title = moduleReq.Title

	if err := mr.conn.WithContext(ctx).Save(&updatedModule).Error; err != nil {
		return modules.Domain{}, err
	}

	return updatedModule.ToDomain(), nil
}

func (mr *moduleRepository) Delete(ctx context.Context, id int) error {
	module, err := mr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedModule := FromDomain(&module)

	if err := mr.conn.WithContext(ctx).Unscoped().Delete(&deletedModule).Error; err != nil {
		return err
	}

	return nil
}
