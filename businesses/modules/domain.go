package modules

import (
	"context"
	"time"

	"github.com/amdrx480/angsana-boga/businesses/chapters"
	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *gorm.DeletedAt
	CourseID  int
	Title     string
	Chapters  []chapters.Domain
}

type UseCase interface {
	Create(ctx context.Context, moduleReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, moduleReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Create(ctx context.Context, moduleReq *Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, moduleReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}
