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

	enrollmentDomain "github.com/amdrx480/go-lms/businesses/enrollments"
	enrollmentDB "github.com/amdrx480/go-lms/drivers/mysql/enrollments"

	lessonDomain "github.com/amdrx480/go-lms/businesses/lessons"
	lessonDB "github.com/amdrx480/go-lms/drivers/mysql/lessons"

	moduleDomain "github.com/amdrx480/go-lms/businesses/modules"
	moduleDB "github.com/amdrx480/go-lms/drivers/mysql/modules"

	otpDomain "github.com/amdrx480/go-lms/businesses/otp"
	otpDB "github.com/amdrx480/go-lms/drivers/redis/otp"

	userDomain "github.com/amdrx480/go-lms/businesses/users"
	userDB "github.com/amdrx480/go-lms/drivers/mysql/users"

	"github.com/redis/go-redis/v9"
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

func NewEnrollmentRepository(conn *gorm.DB) enrollmentDomain.Repository {
	return enrollmentDB.NewMySQLEnrollmentRepository(conn)
}

func NewLessonRepository(conn *gorm.DB) lessonDomain.Repository {
	return lessonDB.NewMySQLLessonRepository(conn)
}

func NewModuleRepository(conn *gorm.DB) moduleDomain.Repository {
	return moduleDB.NewMySQLModuleRepository(conn)
}

func NewOTPRepository(conn *redis.Client) otpDomain.Repository {
	return otpDB.NewOTPRepository(conn)
}

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}
