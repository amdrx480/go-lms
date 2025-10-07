package chapters

import (
	"context"

	"github.com/amdrx480/angsana-boga/businesses/chapters"

	"gorm.io/gorm"
)

type chapterRepository struct {
	conn *gorm.DB
}

func NewMySQLChapterRepository(conn *gorm.DB) chapters.Repository {
	return &chapterRepository{
		conn: conn,
	}
}

func (cr *chapterRepository) Create(ctx context.Context, chapterReq *chapters.Domain) (chapters.Domain, error) {
	record := FromDomain(chapterReq)
	result := cr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return chapters.Domain{}, err
	}

	if err := result.WithContext(ctx).Last(&record).Error; err != nil {
		return chapters.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (cr *chapterRepository) GetAll(ctx context.Context) ([]chapters.Domain, error) {
	var chapterRecords []Chapter

	if err := cr.conn.WithContext(ctx).Preload("Lesson").Find(&chapterRecords).Error; err != nil {
		return nil, err
	}

	return ToDomainList(chapterRecords), nil
}

func (cr *chapterRepository) GetByID(ctx context.Context, id int) (chapters.Domain, error) {
	var chapter Chapter

	if err := cr.conn.WithContext(ctx).First(&chapter, "id = ?", id).Error; err != nil {
		return chapters.Domain{}, err
	}

	return chapter.ToDomain(), nil
}

func (cr *chapterRepository) Update(ctx context.Context, chapterReq *chapters.Domain, id int) (chapters.Domain, error) {
	chapter, err := cr.GetByID(ctx, id)

	if err != nil {
		return chapters.Domain{}, err
	}

	updatedChapter := FromDomain(&chapter)
	updatedChapter.Title = chapterReq.Title

	if err := cr.conn.WithContext(ctx).Save(&updatedChapter).Error; err != nil {
		return chapters.Domain{}, err
	}

	return updatedChapter.ToDomain(), nil
}

func (cr *chapterRepository) Delete(ctx context.Context, id int) error {
	chapter, err := cr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedChapter := FromDomain(&chapter)

	if err := cr.conn.WithContext(ctx).Unscoped().Delete(&deletedChapter).Error; err != nil {
		return err
	}

	return nil
}
