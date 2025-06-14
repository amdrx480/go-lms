package users

import (
	"context"
	"time"

	"github.com/amdrx480/go-lms/utils"
	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	FullName  string
	Username  string
	Email     string
	Password  string
	Role      utils.Role
}

type UseCase interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	Login(ctx context.Context, userDomain *Domain) (string, string, error)
	RefreshAccessToken(ctx context.Context, refreshToken string) (string, error)
	GetUserProfile(ctx context.Context) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, userDomain *Domain) (Domain, error)
	GetByEmail(ctx context.Context, userDomain *Domain) (Domain, error)
	GetUserProfile(ctx context.Context) (Domain, error)
	FindByEmail(ctx context.Context, email string) (Domain, error)
}
