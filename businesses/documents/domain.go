package documents

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	LessonID  int
	Title     string
	FileName  string
	FilePath  string
}

type UseCase interface {
	Create(ctx context.Context, documentReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, documentReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, documentReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, documentReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}
