package chapters

import (
	"context"
	"time"

	"github.com/amdrx480/go-lms/businesses/lessons"
	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	ModuleID  int
	Title     string
	Lessons   []lessons.Domain
}

type UseCase interface {
	Create(ctx context.Context, chapterReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, chapterReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, chapterReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, chapterReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}
