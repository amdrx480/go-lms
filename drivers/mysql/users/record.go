package users

import (
	"time"

	"github.com/amdrx480/angsana-boga/businesses/users"
	"github.com/amdrx480/angsana-boga/utils"
	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	FullName  string         `json:"fullname"`
	Username  string         `json:"username"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	Role      utils.Role     `json:"role" gorm:"type:enum('admin','instructor','user')"`
}

// buat agar otomatis menyimpan user ke role jika tidak ada pilihan role
func (rec *User) BeforeCreate(tx *gorm.DB) (err error) {
	switch rec.Role {
	case utils.ROLE_ADMIN:
		rec.Role = utils.ROLE_ADMIN
	case utils.ROLE_INSTRUCTOR:
		rec.Role = utils.ROLE_INSTRUCTOR
	default:
		rec.Role = utils.ROLE_USER
	}

	return nil
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		FullName:  rec.FullName,
		Username:  rec.Username,
		Email:     rec.Email,
		Password:  rec.Password,
		Role:      rec.Role,
	}
}

func FromDomain(domain *users.Domain) *User {
	return &User{
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
