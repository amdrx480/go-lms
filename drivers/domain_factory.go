package drivers

import (
	categoryDomain "github.com/amdrx480/go-lms/businesses/categories"
	categoryDB "github.com/amdrx480/go-lms/drivers/mysql/categories"

	chapterDomain "github.com/amdrx480/go-lms/businesses/chapters"
	chapterDB "github.com/amdrx480/go-lms/drivers/mysql/chapters"

	courseDomain "github.com/amdrx480/go-lms/businesses/courses"
	courseDB "github.com/amdrx480/go-lms/drivers/mysql/courses"

	documentDomain "github.com/amdrx480/go-lms/businesses/documents"
	documentDB "github.com/amdrx480/go-lms/drivers/mysql/documents"

	lessonDomain "github.com/amdrx480/go-lms/businesses/lessons"
	lessonDB "github.com/amdrx480/go-lms/drivers/mysql/lessons"

	moduleDomain "github.com/amdrx480/go-lms/businesses/modules"
	moduleDB "github.com/amdrx480/go-lms/drivers/mysql/modules"

	userDomain "github.com/amdrx480/go-lms/businesses/users"
	userDB "github.com/amdrx480/go-lms/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewCategoryRepository(conn *gorm.DB) categoryDomain.Repository {
	return categoryDB.NewMySQLCategoryRepository(conn)
}

func NewChapterRepository(conn *gorm.DB) chapterDomain.Repository {
	return chapterDB.NewMySQLChapterRepository(conn)
}

func NewCourseRepository(conn *gorm.DB) courseDomain.Repository {
	return courseDB.NewMySQLRepository(conn)
}

func NewDocumentRepository(conn *gorm.DB) documentDomain.Repository {
	return documentDB.NewMySQLDocumentRepository(conn)
}

func NewLessonRepository(conn *gorm.DB) lessonDomain.Repository {
	return lessonDB.NewMySQLLessonRepository(conn)
}

func NewModuleRepository(conn *gorm.DB) moduleDomain.Repository {
	return moduleDB.NewMySQLModuleRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
