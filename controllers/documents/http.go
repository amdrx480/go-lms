package documents

import (
	"net/http"
	"strconv"

	"github.com/amdrx480/go-lms/businesses/documents"
	"github.com/amdrx480/go-lms/controllers"
	"github.com/amdrx480/go-lms/controllers/documents/request"
	"github.com/amdrx480/go-lms/controllers/documents/response"

	"github.com/labstack/echo/v4"
)

type DocumentController struct {
	documentUseCase documents.UseCase
}

func NewDocumentController(documentUC documents.UseCase) *DocumentController {
	return &DocumentController{
		documentUseCase: documentUC,
	}
}

func (dc *DocumentController) Create(c echo.Context) error {
	documentReq := request.Document{}
	ctx := c.Request().Context()

	if err := c.Bind(&documentReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "invalid request", "")
	}

	if err := documentReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	document, err := dc.documentUseCase.Create(ctx, documentReq.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to create a document", "")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "document created", response.FromDomain(document))
}

func (dc *DocumentController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()
	documentsData, err := dc.documentUseCase.GetAll(ctx)

	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "failed to fetch documents", "")
	}

	documents := response.FromDomainList(documentsData)
	return controllers.NewResponse(c, http.StatusOK, "success", "all documents", documents)
}

func (dc *DocumentController) GetByID(c echo.Context) error {
	documentsID := c.Param("id")
	ctx := c.Request().Context()

	dID, err := strconv.Atoi(documentsID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	document, err := dc.documentUseCase.GetByID(ctx, dID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "document not found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "document found", response.FromDomain(document))
}

func (dc *DocumentController) Update(c echo.Context) error {
	documentReq := request.Document{}
	ctx := c.Request().Context()

	if err := c.Bind(&documentReq); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	documentsID := c.Param("id")

	dID, err := strconv.Atoi(documentsID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	if err := documentReq.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", err.Error(), "")
	}

	document, err := dc.documentUseCase.Update(ctx, documentReq.ToDomain(), dID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "update document failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "document updated", response.FromDomain(document))
}

func (dc *DocumentController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	documentsID := c.Param("id")

	dID, err := strconv.Atoi(documentsID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusUnprocessableEntity, "failed", "id must be valid integer", "")
	}

	err = dc.documentUseCase.Delete(ctx, dID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", "delete document failed", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "document deleted", "")
}
