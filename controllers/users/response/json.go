package response

import (
	"time"

	"github.com/amdrx480/go-lms/businesses/users"
	"github.com/amdrx480/go-lms/utils"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	FullName  string         `json:"fullname"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      utils.Role     `json:"role"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		FullName:  domain.FullName,
		Username:  domain.Username,
		Email:     domain.Email,
		Password:  domain.Password,
		Role:      domain.Role,
	}
}
