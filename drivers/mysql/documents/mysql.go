package documents

import (
	"context"

	"github.com/amdrx480/angsana-boga/businesses/documents"

	"gorm.io/gorm"
)

type documentRepository struct {
	conn *gorm.DB
}

func NewMySQLDocumentRepository(conn *gorm.DB) documents.Repository {
	return &documentRepository{
		conn: conn,
	}
}

func (dr *documentRepository) Create(ctx context.Context, documentReq *documents.Domain) (documents.Domain, error) {
	record := FromDomain(documentReq)
	result := dr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return documents.Domain{}, err
	}

	if err := result.WithContext(ctx).Last(&record).Error; err != nil {
		return documents.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (dr *documentRepository) GetAll(ctx context.Context) ([]documents.Domain, error) {
	var lessonRecords []Document

	if err := dr.conn.WithContext(ctx).Find(&lessonRecords).Error; err != nil {
		return nil, err
	}

	return ToDomainList(lessonRecords), nil
}

func (dr *documentRepository) GetByID(ctx context.Context, id int) (documents.Domain, error) {
	var lesson Document

	if err := dr.conn.WithContext(ctx).First(&lesson, "id = ?", id).Error; err != nil {
		return documents.Domain{}, err
	}

	return lesson.ToDomain(), nil
}

func (dr *documentRepository) Update(ctx context.Context, documentReq *documents.Domain, id int) (documents.Domain, error) {
	lesson, err := dr.GetByID(ctx, id)

	if err != nil {
		return documents.Domain{}, err
	}

	updatedLesson := FromDomain(&lesson)
	updatedLesson.Title = documentReq.Title
	updatedLesson.FileName = lesson.FileName
	updatedLesson.FilePath = lesson.FilePath

	if err := dr.conn.WithContext(ctx).Save(&updatedLesson).Error; err != nil {
		return documents.Domain{}, err
	}

	return updatedLesson.ToDomain(), nil
}

func (dr *documentRepository) Delete(ctx context.Context, id int) error {
	lesson, err := dr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedLesson := FromDomain(&lesson)

	if err := dr.conn.WithContext(ctx).Unscoped().Delete(&deletedLesson).Error; err != nil {
		return err
	}

	return nil
}
