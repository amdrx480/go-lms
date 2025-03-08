package drivers

import (
	userDomain "github.com/amdrx480/go-clean-architecture-hexagonal/businesses/users"
	userDB "github.com/amdrx480/go-clean-architecture-hexagonal/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
