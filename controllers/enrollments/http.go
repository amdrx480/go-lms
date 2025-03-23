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

// func (ec *EnrollmentController) CreateEnrollmentCourse(c echo.Context) error {
// 	enrollmentReq := request.Enrollment{}
// 	ctx := c.Request().Context()

// 	if err := c.Bind(&enrollmentReq); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
// 	}

// 	if err := enrollmentReq.Validate(); err != nil {
// 		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
// 	}

// 	enrollment, err := ec.enrollmentUseCase.CreateEnrollmentCourse(ctx, enrollmentReq.ToDomain())
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a enrollment", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusCreated, "success", "enrollment created", response.FromDomain(enrollment))
// }

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
	userID := c.Param("user_id")
	ctx := c.Request().Context()

	uID, err := strconv.Atoi(userID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	enrollmentsData, err := ec.enrollmentUseCase.GetAllEnrollmentCourseByUserID(ctx, uID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	// ✅ Konversi ke response
	enrollments := response.FromDomainList(enrollmentsData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all enrollments", enrollments)
}

// func (ec *EnrollmentController) GetAll(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	enrollmentsData, err := ec.enrollmentUseCase.GetAll(ctx)

// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
// 	}

// 	enrollments := response.FromDomainList(enrollmentsData)

// 	return controllers.NewResponse(c, http.StatusOK, "success", "all enrollments", enrollments)
// }

// func (ec *EnrollmentController) GetByID(c echo.Context) error {
// 	enrollmentID := c.Param("id")
// 	ctx := c.Request().Context()

// 	cID, err := strconv.Atoi(enrollmentID)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
// 	}

// 	enrollment, err := ec.enrollmentUseCase.GetByID(ctx, cID)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusNotFound, "failed", "enrollment not found", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, "success", "enrollment found", response.FromDomain(enrollment))
// }

// func (ec *EnrollmentController) Update(c echo.Context) error {
// 	enrollmentReq := request.Enrollment{}
// 	ctx := c.Request().Context()

// 	if err := c.Bind(&enrollmentReq); err != nil {
// 		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
// 	}

// 	enrollmentID := c.Param("id")

// 	cID, err := strconv.Atoi(enrollmentID)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
// 	}

// 	if err := enrollmentReq.Validate(); err != nil {
// 		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
// 	}

// 	enrollment, err := ec.enrollmentUseCase.Update(ctx, enrollmentReq.ToDomain(), cID)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update a enrollment", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, "success", "enrollment updated", response.FromDomain(enrollment))
// }

// func (ec *EnrollmentController) Delete(c echo.Context) error {
// 	enrollmentID := c.Param("id")
// 	ctx := c.Request().Context()

// 	cID, err := strconv.Atoi(enrollmentID)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
// 	}

// 	err = ec.enrollmentUseCase.Delete(ctx, cID)
// 	if err != nil {
// 		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete a enrollment", "")
// 	}

// 	return controllers.NewResponse(c, http.StatusOK, "success", "enrollment deleted", "")
// }
