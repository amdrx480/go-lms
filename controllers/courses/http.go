package courses

import (
	"strconv"

	"github.com/amdrx480/go-lms/businesses/courses"
	"github.com/amdrx480/go-lms/controllers"
	"github.com/amdrx480/go-lms/controllers/courses/request"
	"github.com/amdrx480/go-lms/controllers/courses/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseController struct {
	courseUseCase courses.UseCase
}

func NewCourseController(courseUC courses.UseCase) *CourseController {
	return &CourseController{
		courseUseCase: courseUC,
	}
}

func (cc *CourseController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	coursesData, err := cc.courseUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch courses", "")
	}

	// Mengonversi seluruh data domain ke model respons
	courses := response.FromDomainList(coursesData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all courses", courses)
}

func (cc *CourseController) GetAllWithModules(c echo.Context) error {
	ctx := c.Request().Context()
	coursesData, err := cc.courseUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch courses", "")
	}

	// Mengonversi seluruh data domain ke model respons
	courses := response.FromDomainWithModulesList(coursesData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all courses", courses)
}

func (cc *CourseController) GetByID(c echo.Context) error {
	courseID := c.Param("id")
	ctx := c.Request().Context()

	cID, err := strconv.Atoi(courseID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	course, err := cc.courseUseCase.GetByID(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "course not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course found", response.FromDomain(course))
}

func (cc *CourseController) Create(c echo.Context) error {
	courseReq := request.Course{}
	ctx := c.Request().Context()

	if err := c.Bind(&courseReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	if err := courseReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	course, err := cc.courseUseCase.Create(ctx, courseReq.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a course", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "course created", response.FromDomain(course))
}

func (cc *CourseController) Update(c echo.Context) error {
	courseReq := request.Course{}
	ctx := c.Request().Context()

	if err := c.Bind(&courseReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	courseID := c.Param("id")

	cID, err := strconv.Atoi(courseID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	if err := courseReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	course, err := cc.courseUseCase.Update(ctx, courseReq.ToDomain(), cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update course failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course updated", response.FromDomain(course))
}

func (cc *CourseController) Delete(c echo.Context) error {
	courseID := c.Param("id")
	ctx := c.Request().Context()

	cID, err := strconv.Atoi(courseID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = cc.courseUseCase.Delete(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete course failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course deleted", "")
}

func (cc *CourseController) Restore(c echo.Context) error {
	ctx := c.Request().Context()
	courseID := c.Param("id")

	cID, err := strconv.Atoi(courseID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	course, err := cc.courseUseCase.Restore(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "course not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course restored", response.FromDomain(course))
}

func (cc *CourseController) ForceDelete(c echo.Context) error {
	ctx := c.Request().Context()
	courseID := c.Param("id")

	cID, err := strconv.Atoi(courseID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = cc.courseUseCase.ForceDelete(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "force delete course failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "course deleted permanently", "")
}
