package http

import (
	"echo-crud/entity"
	"echo-crud/internal/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateTransaksiBodyRequest defines all body attributes needed to add transaksi.
type CreateTransaksiBodyRequest struct {
	Tanggal    string `json:"tanggal" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	Total      int64  `json:"total" binding:"required"`
}

// TransaksiRowResponse defines all attributes needed to fulfill for transaksi row entity.
type TransaksiRowResponse struct {
	ID         uuid.UUID `json:"id_transaksi"`
	Tanggal    string    `json:"tanggal"`
	Keterangan string    `json:"keterangan"`
	Total      int64     `json:"total"`
}

// TransaksiResponse defines all attributes needed to fulfill for pic transaksi entity.
type TransaksiDetailResponse struct {
	ID         uuid.UUID `json:"id_transaksi"`
	Tanggal    string    `json:"tanggal"`
	Keterangan string    `json:"keterangan"`
	Total      int64     `json:"total"`
}

func buildTransaksiRowResponse(transaksi *entity.Transaksi) TransaksiRowResponse {
	form := TransaksiRowResponse{
		ID:         transaksi.ID,
		Tanggal:    transaksi.Tanggal,
		Keterangan: transaksi.Keterangan,
		Total:      transaksi.Total,
	}

	return form
}

func buildTransaksiDetailResponse(transaksi *entity.Transaksi) TransaksiDetailResponse {
	form := TransaksiDetailResponse{
		ID:         transaksi.ID,
		Tanggal:    transaksi.Tanggal,
		Keterangan: transaksi.Keterangan,
		Total:      transaksi.Total,
	}

	return form
}

// QueryParamsTransaksi defines all attributes for input query params
type QueryParamsTransaksi struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaTransaksi define attributes needed for Meta
type MetaTransaksi struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaTransaksi creates an instance of Meta response.
func NewMetaTransaksi(limit, offset int, total int64) *MetaTransaksi {
	return &MetaTransaksi{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// TransaksiHandler handles HTTP request related to user flow.
type TransaksiHandler struct {
	service service.TransaksiUseCase
}

// NewTransaksiHandler creates an instance of TransaksiHandler.
func NewTransaksiHandler(service service.TransaksiUseCase) *TransaksiHandler {
	return &TransaksiHandler{
		service: service,
	}
}

// Create handles transaksi creation.
// It will reject the request if the request doesn't have required data,
func (handler *TransaksiHandler) CreateTransaksi(echoCtx echo.Context) error {
	var form CreateTransaksiBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	transaksiEntity := entity.NewTransaksi(
		uuid.Nil,
		form.Tanggal,
		form.Keterangan,
		form.Total,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), transaksiEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", transaksiEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TransaksiHandler) GetListTransaksi(echoCtx echo.Context) error {
	var form QueryParamsTransaksi
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	transaksi, err := handler.service.GetListTransaksi(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", transaksi)
	return echoCtx.JSON(res.Status, res)

}

func (handler *TransaksiHandler) GetDetailTransaksi(echoCtx echo.Context) error {
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

	transaksi, err := handler.service.GetDetailTransaksi(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", transaksi)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TransaksiHandler) UpdateTransaksi(echoCtx echo.Context) error {
	var form CreateTransaksiBodyRequest
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

	_, err = handler.service.GetDetailTransaksi(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	transaksiEntity := &entity.Transaksi{
		id,
		form.Tanggal,
		form.Keterangan,
		form.Total,
	}

	if err := handler.service.UpdateTransaksi(echoCtx.Request().Context(), transaksiEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *TransaksiHandler) DeleteTransaksi(echoCtx echo.Context) error {
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

	err = handler.service.DeleteTransaksi(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
