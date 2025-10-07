package categories

import (
	"net/http"
	"strconv"

	"github.com/amdrx480/angsana-boga/businesses/categories"
	"github.com/amdrx480/angsana-boga/controllers"
	"github.com/amdrx480/angsana-boga/controllers/categories/request"
	"github.com/amdrx480/angsana-boga/controllers/categories/response"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUseCase categories.Usecase
}

func NewCategoryController(categoryUC categories.Usecase) *CategoryController {
	return &CategoryController{
		categoryUseCase: categoryUC,
	}
}

func (cc *CategoryController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	categoriesData, err := cc.categoryUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch data", "")
	}

	categories := response.FromDomainList(categoriesData)

	return controllers.NewResponse(c, http.StatusOK, "success", "all categories", categories)
}

func (cc *CategoryController) GetByID(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	cID, err := strconv.Atoi(categoryID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	category, err := cc.categoryUseCase.GetByID(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "category not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category found", response.FromDomain(category))
}

func (cc *CategoryController) Create(c echo.Context) error {
	categoryReq := request.Category{}
	ctx := c.Request().Context()

	if err := c.Bind(&categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	if err := categoryReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	category, err := cc.categoryUseCase.Create(ctx, categoryReq.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a category", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "category created", response.FromDomain(category))
}

func (cc *CategoryController) Update(c echo.Context) error {
	categoryReq := request.Category{}
	ctx := c.Request().Context()

	if err := c.Bind(&categoryReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	categoryID := c.Param("id")

	cID, err := strconv.Atoi(categoryID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	if err := categoryReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	category, err := cc.categoryUseCase.Update(ctx, categoryReq.ToDomain(), cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to update a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category updated", response.FromDomain(category))
}

func (cc *CategoryController) Delete(c echo.Context) error {
	categoryID := c.Param("id")
	ctx := c.Request().Context()

	cID, err := strconv.Atoi(categoryID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = cc.categoryUseCase.Delete(ctx, cID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to delete a category", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "category deleted", "")
}
