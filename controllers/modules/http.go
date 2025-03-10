package modules

import (
	"net/http"
	"strconv"

	"github.com/amdrx480/go-lms/businesses/modules"
	"github.com/amdrx480/go-lms/controllers"
	"github.com/amdrx480/go-lms/controllers/modules/request"
	"github.com/amdrx480/go-lms/controllers/modules/response"

	"github.com/labstack/echo/v4"
)

type ModuleController struct {
	moduleUseCase modules.UseCase
}

func NewModuleController(moduleUC modules.UseCase) *ModuleController {
	return &ModuleController{
		moduleUseCase: moduleUC,
	}
}

func (mc *ModuleController) Create(c echo.Context) error {
	moduleReq := request.Module{}
	ctx := c.Request().Context()

	if err := c.Bind(&moduleReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := moduleReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	module, err := mc.moduleUseCase.Create(ctx, moduleReq.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a module", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "module created", response.FromDomain(module))
}

func (mc *ModuleController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	modulesData, err := mc.moduleUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch modules", "")
	}

	modules := []response.Module{}

	for _, module := range modulesData {
		modules = append(modules, response.FromDomain(module))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all modules", modules)
}

func (mc *ModuleController) GetAllWithChapter(c echo.Context) error {
	ctx := c.Request().Context()
	coursesData, err := mc.moduleUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch modules", "")
	}

	// Mengonversi seluruh data domain ke model respons
	courses := response.FromDomainWithChapterList(coursesData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all modules", courses)
}

func (mc *ModuleController) GetByID(c echo.Context) error {
	moduleID := c.Param("id")
	ctx := c.Request().Context()

	mID, err := strconv.Atoi(moduleID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	module, err := mc.moduleUseCase.GetByID(ctx, mID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "module not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "module found", response.FromDomain(module))
}

func (mc *ModuleController) Update(c echo.Context) error {
	moduleReq := request.Module{}
	ctx := c.Request().Context()

	if err := c.Bind(&moduleReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	moduleID := c.Param("id")

	mID, err := strconv.Atoi(moduleID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	if err := moduleReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	module, err := mc.moduleUseCase.Update(ctx, moduleReq.ToDomain(), mID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update module failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "module updated", response.FromDomain(module))
}

func (mc *ModuleController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	moduleID := c.Param("id")

	mID, err := strconv.Atoi(moduleID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = mc.moduleUseCase.Delete(ctx, mID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete module failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "module deleted", "")
}
