package lessons

import (
	"net/http"
	"strconv"

	"github.com/amdrx480/go-lms/businesses/lessons"
	"github.com/amdrx480/go-lms/controllers"
	"github.com/amdrx480/go-lms/controllers/lessons/request"
	"github.com/amdrx480/go-lms/controllers/lessons/response"

	"github.com/labstack/echo/v4"
)

type LessonController struct {
	lessonUseCase lessons.UseCase
}

func NewLessonController(lessonUC lessons.UseCase) *LessonController {
	return &LessonController{
		lessonUseCase: lessonUC,
	}
}

func (lc *LessonController) Create(c echo.Context) error {
	lessonReq := request.Lesson{}
	ctx := c.Request().Context()

	if err := c.Bind(&lessonReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := lessonReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	lesson, err := lc.lessonUseCase.Create(ctx, lessonReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a lesson", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "lesson created", response.FromDomain(lesson))
}

func (lc *LessonController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	lessonsData, err := lc.lessonUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch lessons", "")
	}

	lessons := response.FromDomainList(lessonsData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all lessons", lessons)
}

func (lc *LessonController) GetByID(c echo.Context) error {
	lessonsID := c.Param("id")
	ctx := c.Request().Context()

	lID, err := strconv.Atoi(lessonsID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	lesson, err := lc.lessonUseCase.GetByID(ctx, lID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "lesson not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "lesson found", response.FromDomain(lesson))
}

func (lc *LessonController) Update(c echo.Context) error {
	lessonReq := request.Lesson{}
	ctx := c.Request().Context()

	if err := c.Bind(&lessonReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	lessonsID := c.Param("id")
	lID, err := strconv.Atoi(lessonsID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	if err := lessonReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	lesson, err := lc.lessonUseCase.Update(ctx, lessonReq.ToDomain(), lID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update lesson failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "lesson updated", response.FromDomain(lesson))
}

func (lc *LessonController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	lessonsID := c.Param("id")

	lID, err := strconv.Atoi(lessonsID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = lc.lessonUseCase.Delete(ctx, lID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete lesson failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "lesson deleted", "")
}
