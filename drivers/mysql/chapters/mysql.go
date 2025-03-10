package chapters

import (
	"context"

	"github.com/amdrx480/go-lms/businesses/chapters"

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

func (mr *chapterRepository) Create(ctx context.Context, chapterReq *chapters.Domain) (chapters.Domain, error) {
	record := FromDomain(chapterReq)
	result := mr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return chapters.Domain{}, err
	}

	if err := result.WithContext(ctx).Last(&record).Error; err != nil {
		return chapters.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (mr *chapterRepository) GetAll(ctx context.Context) ([]chapters.Domain, error) {
	var chapterRecords []Chapter

	if err := mr.conn.WithContext(ctx).Find(&chapterRecords).Error; err != nil {
		return nil, err
	}

	return ToDomainList(chapterRecords), nil
}

func (mr *chapterRepository) GetByID(ctx context.Context, id int) (chapters.Domain, error) {
	var chapter Chapter

	if err := mr.conn.WithContext(ctx).First(&chapter, "id = ?", id).Error; err != nil {
		return chapters.Domain{}, err
	}

	return chapter.ToDomain(), nil
}

func (mr *chapterRepository) Update(ctx context.Context, chapterReq *chapters.Domain, id int) (chapters.Domain, error) {
	chapter, err := mr.GetByID(ctx, id)

	if err != nil {
		return chapters.Domain{}, err
	}

	updatedChapter := FromDomain(&chapter)
	updatedChapter.Title = chapterReq.Title

	if err := mr.conn.WithContext(ctx).Save(&updatedChapter).Error; err != nil {
		return chapters.Domain{}, err
	}

	return updatedChapter.ToDomain(), nil
}

func (mr *chapterRepository) Delete(ctx context.Context, id int) error {
	chapter, err := mr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedChapter := FromDomain(&chapter)

	if err := mr.conn.WithContext(ctx).Unscoped().Delete(&deletedChapter).Error; err != nil {
		return err
	}

	return nil
}
