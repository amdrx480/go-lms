package courses

import (
	"context"
	"time"

	"github.com/amdrx480/go-lms/businesses/categories"
	"github.com/amdrx480/go-lms/businesses/modules"
	"gorm.io/gorm"
)

type Domain struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *gorm.DeletedAt
	Slug        string
	Title       string
	Description string
	Category    categories.Domain
	CategoryID  int
	Cover       string
	Instructor  string
	Modules     []modules.Domain
}

type UseCase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, courseReq *Domain) (Domain, error)
	Update(ctx context.Context, courseReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
	Restore(ctx context.Context, id int) (Domain, error)
	ForceDelete(ctx context.Context, id int) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, courseReq *Domain) (Domain, error)
	Update(ctx context.Context, courseReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
	Restore(ctx context.Context, id int) (Domain, error)
	ForceDelete(ctx context.Context, id int) error
}
