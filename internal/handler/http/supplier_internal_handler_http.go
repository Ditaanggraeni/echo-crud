package http

import (
	"echo-crud/entity"
	"echo-crud/internal/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateSupplierBodyRequest defines all body attributes needed to add supplier.
type CreateSupplierBodyRequest struct {
	NamaSupplier string `json:"nama_supplier" binding:"required"`
	Telepon      string `json:"telepon" binding:"required"`
	Alamat       string `json:"alamat" binding:"required"`
}

// SupplierRowResponse defines all attributes needed to fulfill for supplier row entity.
type SupplierRowResponse struct {
	ID           uuid.UUID `json:"id_supplier"`
	NamaSupplier string    `json:"nama_supplier"`
	Telepon      string    `json:"telepon"`
	Alamat       string    `json:"alamat"`
}

// SupplierResponse defines all attributes needed to fulfill for pic supplier entity.
type SupplierDetailResponse struct {
	ID           uuid.UUID `json:"id_supplier"`
	NamaSupplier string    `json:"nama_supplier"`
	Telepon      string    `json:"telepon"`
	Alamat       string    `json:"alamat"`
}

func buildSupplierRowResponse(supplier *entity.Supplier) SupplierRowResponse {
	form := SupplierRowResponse{
		ID:           supplier.ID,
		NamaSupplier: supplier.NamaSupplier,
		Alamat:       supplier.Alamat,
		Telepon:      supplier.Telepon,
	}

	return form
}

func buildSupplierDetailResponse(supplier *entity.Supplier) SupplierDetailResponse {
	form := SupplierDetailResponse{
		ID:           supplier.ID,
		NamaSupplier: supplier.NamaSupplier,
		Alamat:       supplier.Alamat,
		Telepon:      supplier.Telepon,
	}

	return form
}

// QueryParamsSupplier defines all attributes for input query params
type QueryParamsSupplier struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaSupplier define attributes needed for Meta
type MetaSupplier struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaSupplier creates an instance of Meta response.
func NewMetaSupplier(limit, offset int, total int64) *MetaSupplier {
	return &MetaSupplier{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// SupplierHandler handles HTTP request related to user flow.
type SupplierHandler struct {
	service service.SupplierUseCase
}

// NewSupplierHandler creates an instance of SupplierHandler.
func NewSupplierHandler(service service.SupplierUseCase) *SupplierHandler {
	return &SupplierHandler{
		service: service,
	}
}

// Create handles supplier creation.
// It will reject the request if the request doesn't have required data,
func (handler *SupplierHandler) CreateSupplier(echoCtx echo.Context) error {
	var form CreateSupplierBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	supplierEntity := entity.NewSupplier(
		uuid.Nil,
		form.NamaSupplier,
		form.Telepon,
		form.Alamat,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), supplierEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", supplierEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *SupplierHandler) GetListSupplier(echoCtx echo.Context) error {
	var form QueryParamsSupplier
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	supplier, err := handler.service.GetListSupplier(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", supplier)
	return echoCtx.JSON(res.Status, res)

}

func (handler *SupplierHandler) GetDetailSupplier(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	supplier, err := handler.service.GetDetailSupplier(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", supplier)
	return echoCtx.JSON(res.Status, res)
}

func (handler *SupplierHandler) UpdateSupplier(echoCtx echo.Context) error {
	var form CreateSupplierBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)

	}

	idParam := echoCtx.Param("id")

	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	_, err = handler.service.GetDetailSupplier(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	supplierEntity := &entity.Supplier{
		id,
		form.NamaSupplier,
		form.Telepon,
		form.Alamat,
	}

	if err := handler.service.UpdateSupplier(echoCtx.Request().Context(), supplierEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *SupplierHandler) DeleteSupplier(echoCtx echo.Context) error {
	idParam := echoCtx.Param("id")
	if len(idParam) == 0 {
		errorResponse := buildErrorResponse(nil, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	err = handler.service.DeleteSupplier(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
