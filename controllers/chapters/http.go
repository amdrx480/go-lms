package chapters

import (
	"net/http"
	"strconv"

	"github.com/amdrx480/angsana-boga/businesses/chapters"
	"github.com/amdrx480/angsana-boga/controllers"
	"github.com/amdrx480/angsana-boga/controllers/chapters/request"
	"github.com/amdrx480/angsana-boga/controllers/chapters/response"

	"github.com/labstack/echo/v4"
)

type ChapterController struct {
	chapterUseCase chapters.UseCase
}

func NewChapterController(chapterUC chapters.UseCase) *ChapterController {
	return &ChapterController{
		chapterUseCase: chapterUC,
	}
}

func (cc *ChapterController) Create(c echo.Context) error {
	chapterReq := request.Chapter{}
	ctx := c.Request().Context()

	if err := c.Bind(&chapterReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := chapterReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	chapter, err := cc.chapterUseCase.Create(ctx, chapterReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a chapter", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "chapter created", response.FromDomain(chapter))
}

func (cc *ChapterController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	chaptersData, err := cc.chapterUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch chapters", "")
	}

	chapters := response.FromDomainList(chaptersData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all chapters", chapters)
}

func (cc *ChapterController) GetByID(c echo.Context) error {
	chapterID := c.Param("id")
	ctx := c.Request().Context()

	cID, err := strconv.Atoi(chapterID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	chapter, err := cc.chapterUseCase.GetByID(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "chapter not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "chapter found", response.FromDomain(chapter))
}

func (cc *ChapterController) Update(c echo.Context) error {
	chapterReq := request.Chapter{}
	ctx := c.Request().Context()

	if err := c.Bind(&chapterReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	chapterID := c.Param("id")

	cID, err := strconv.Atoi(chapterID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	if err := chapterReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	chapter, err := cc.chapterUseCase.Update(ctx, chapterReq.ToDomain(), cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update chapter failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "chapter updated", response.FromDomain(chapter))
}

func (cc *ChapterController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	chapterID := c.Param("id")

	cID, err := strconv.Atoi(chapterID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = cc.chapterUseCase.Delete(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete chapter failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "chapter deleted", "")
}
