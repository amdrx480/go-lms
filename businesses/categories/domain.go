package categories

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Title     string
	Slug      string
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, categoryReq *Domain) (Domain, error)
	Update(ctx context.Context, categoryReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Create(ctx context.Context, categoryReq *Domain) (Domain, error)
	Update(ctx context.Context, categoryReq *Domain, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
}
