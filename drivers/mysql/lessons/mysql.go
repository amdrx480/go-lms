package lessons

import (
	"context"

	"github.com/amdrx480/go-lms/businesses/lessons"

	"gorm.io/gorm"
)

type lessonRepository struct {
	conn *gorm.DB
}

func NewMySQLLessonRepository(conn *gorm.DB) lessons.Repository {
	return &lessonRepository{
		conn: conn,
	}
}

func (mr *lessonRepository) Create(ctx context.Context, lessonReq *lessons.Domain) (lessons.Domain, error) {
	record := FromDomain(lessonReq)
	result := mr.conn.WithContext(ctx).Create(&record)

	if err := result.Error; err != nil {
		return lessons.Domain{}, err
	}

	if err := result.WithContext(ctx).Last(&record).Error; err != nil {
		return lessons.Domain{}, err
	}

	return record.ToDomain(), nil
}

func (mr *lessonRepository) GetAll(ctx context.Context) ([]lessons.Domain, error) {
	var lessonRecords []Lesson

	if err := mr.conn.WithContext(ctx).Find(&lessonRecords).Error; err != nil {
		return nil, err
	}

	return ToDomainList(lessonRecords), nil
}

func (mr *lessonRepository) GetByID(ctx context.Context, id int) (lessons.Domain, error) {
	var lesson Lesson

	if err := mr.conn.WithContext(ctx).First(&lesson, "id = ?", id).Error; err != nil {
		return lessons.Domain{}, err
	}

	return lesson.ToDomain(), nil
}

func (mr *lessonRepository) Update(ctx context.Context, lessonReq *lessons.Domain, id int) (lessons.Domain, error) {
	lesson, err := mr.GetByID(ctx, id)

	if err != nil {
		return lessons.Domain{}, err
	}

	updatedLesson := FromDomain(&lesson)
	updatedLesson.Title = lessonReq.Title
	updatedLesson.Content = lesson.Content
	updatedLesson.VideoURL = lesson.VideoURL

	if err := mr.conn.WithContext(ctx).Save(&updatedLesson).Error; err != nil {
		return lessons.Domain{}, err
	}

	return updatedLesson.ToDomain(), nil
}

func (mr *lessonRepository) Delete(ctx context.Context, id int) error {
	lesson, err := mr.GetByID(ctx, id)

	if err != nil {
		return err
	}

	deletedLesson := FromDomain(&lesson)

	if err := mr.conn.WithContext(ctx).Unscoped().Delete(&deletedLesson).Error; err != nil {
		return err
	}

	return nil
}
