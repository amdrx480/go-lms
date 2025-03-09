package drivers

import (
	categoryDomain "github.com/amdrx480/go-lms/businesses/categories"
	categoryDB "github.com/amdrx480/go-lms/drivers/mysql/categories"

	moduleDomain "github.com/amdrx480/go-lms/businesses/modules"
	moduleDB "github.com/amdrx480/go-lms/drivers/mysql/modules"

	courseDomain "github.com/amdrx480/go-lms/businesses/courses"
	courseDB "github.com/amdrx480/go-lms/drivers/mysql/courses"

	userDomain "github.com/amdrx480/go-lms/businesses/users"
	userDB "github.com/amdrx480/go-lms/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLRepository(conn)
}

func NewModuleRepository(conn *gorm.DB) moduleDomain.Repository {
	return moduleDB.NewMySQLModuleRepository(conn)
}

func NewCourseRepository(conn *gorm.DB) courseDomain.Repository {
	return courseDB.NewMySQLRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
