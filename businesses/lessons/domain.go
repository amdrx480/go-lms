package lessons

import (
	"context"
	"time"

	"github.com/amdrx480/angsana-boga/businesses/documents"
	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	ChapterID int
	Title     string
	Content   string
	VideoURL  string
	Documents []documents.Domain
}

type UseCase interface {
	Create(ctx context.Context, lessonReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, lessonReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, lessonReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, lessonReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}
