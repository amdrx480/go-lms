package enrollments

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/amdrx480/go-lms/businesses/enrollments"
	"github.com/amdrx480/go-lms/controllers"
	"github.com/amdrx480/go-lms/controllers/enrollments/request"
	"github.com/amdrx480/go-lms/controllers/enrollments/response"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type EnrollmentController struct {
	enrollmentUseCase enrollments.Usecase
}

func NewEnrollmentController(enrollmentUC enrollments.Usecase) *EnrollmentController {
	return &EnrollmentController{
		enrollmentUseCase: enrollmentUC,
	}
}

func (ec *EnrollmentController) CreateEnrollmentCourse(c echo.Context) error {
	enrollmentReq := request.Enrollment{}
	ctx := c.Request().Context()

	if err := c.Bind(&enrollmentReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	if err := enrollmentReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	// Cek apakah user sudah enroll di course yang sama
	if _, err := ec.enrollmentUseCase.GetEnrollmentByUserCourse(ctx, enrollmentReq.UserID, enrollmentReq.CourseID); err == nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "User already enrolled in this course", "")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to check enrollment", "")
	}

	// Jika belum ada, buat enrollment baru
	enrollment, err := ec.enrollmentUseCase.CreateEnrollmentCourse(ctx, enrollmentReq.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create enrollment", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "enrollment created", response.FromDomain(enrollment))
}

func (ec *EnrollmentController) GetEnrollmentByUserCourse(c echo.Context) error {
	userID := c.Param("user_id")     // Ambil user_id dari URL
	courseID := c.Param("course_id") // Ambil course_id dari URL
	ctx := c.Request().Context()     // Ambil context request

	// Konversi string ke int
	uID, err := strconv.Atoi(userID)
	cID, err2 := strconv.Atoi(courseID)
	if err != nil || err2 != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "user_id and course_id must be valid integers", "")
	}

	// Panggil UseCase untuk mencari Enrollment
	enrollment, err := ec.enrollmentUseCase.GetEnrollmentByUserCourse(ctx, uID, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "enrollment not found", "")
	}

	// Jika ditemukan, kembalikan response
	return controllers.NewResponse(c, http.StatusOK, "success", "enrollment found", response.FromDomain(enrollment))
}

func (ec *EnrollmentController) GetAllEnrollmentCourseByUserID(c echo.Context) error {
	ctx := c.Request().Context()

	enrollmentsData, err := ec.enrollmentUseCase.GetAllEnrollmentCourseByUserID(ctx)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	enrollments := response.FromDomainList(enrollmentsData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all enrollments", enrollments)
}
