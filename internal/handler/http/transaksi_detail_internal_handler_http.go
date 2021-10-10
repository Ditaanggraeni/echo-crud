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
type CreateDetailTransaksiBodyRequest struct {
	Produk    string `json:"produk" binding:"required"`
	Kuantitas int64 `json:"kuantitas" binding:"required"`
	Total      int64  `json:"total" binding:"required"`
}

// TransaksiRowResponse defines all attributes needed to fulfill for transaksi row entity.
type DetailTransaksiRowResponse struct {
	ID         uuid.UUID `json:"id"`
	Produk    string    `json:"produk"`
	Kuantitas int64    `json:"kuantitas"`
	Total      int64     `json:"total"`
}

// TransaksiResponse defines all attributes needed to fulfill for pic transaksi entity.
type DetailTransaksiDetailResponse struct {
	ID         uuid.UUID `json:"id"`
	Produk    string    `json:"produk"`
	Kuantitas int64    `json:"kuantitas"`
	Total      int64     `json:"total"`
}

func buildDetailTransaksiRowResponse(transaksi_detail *entity.TransaksiDetail) DetailTransaksiRowResponse {
	form := DetailTransaksiRowResponse{
		ID:         transaksi_detail.ID,
		Produk:    transaksi_detail.Produk,
		Kuantitas: transaksi_detail.Kuantitas,
		Total:      transaksi_detail.Total,
	}

	return form
}

func buildDetailTransaksiDetailResponse(transaksi_detail *entity.TransaksiDetail) DetailTransaksiDetailResponse {
	form := DetailTransaksiDetailResponse{
		ID:         transaksi_detail.ID,
		Produk:    transaksi_detail.Produk,
		Kuantitas: transaksi_detail.Kuantitas,
		Total:      transaksi_detail.Total,
		//PelangganID: transaksi.PelangganID,
	}

	return form
}

// QueryParamsTransaksi defines all attributes for input query params
type QueryParamsTransaksiDetail struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaTransaksi define attributes needed for Meta
// type MetaTransaksi struct {
// 	Limit  int   `json:"limit"`
// 	Offset int   `json:"offset"`
// 	Total  int64 `json:"total"`
// }

// // NewMetaTransaksi creates an instance of Meta response.
// func NewMetaTransaksi(limit, offset int, total int64) *MetaTransaksi {
// 	return &MetaTransaksi{
// 		Limit:  limit,
// 		Offset: offset,
// 		Total:  total,
// 	}
// }

// TransaksiHandler handles HTTP request related to user flow.
type DetailTransaksiHandler struct {
	service service.DetailTransaksiUseCase
}

// NewTransaksiHandler creates an instance of TransaksiHandler.
func NewDetailTransaksiHandler(service service.DetailTransaksiUseCase) *DetailTransaksiHandler {
	return &DetailTransaksiHandler{
		service: service,
	}
}

// Create handles transaksi creation.
// It will reject the request if the request doesn't have required data,
func (handler *DetailTransaksiHandler) CreateDetailTransaksi(echoCtx echo.Context) error {
	var form CreateDetailTransaksiBodyRequest

	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	transaksiDetailEntity := entity.NewTransaksiDetail(
		uuid.Nil,
		form.Produk,
		int64(form.Kuantitas),
		int64(form.Total),
	)

	if err := handler.service.Create(echoCtx.Request().Context(), transaksiDetailEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", transaksiDetailEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *DetailTransaksiHandler) GetListTransaksi_Detail(echoCtx echo.Context) error {
	var form QueryParamsTransaksi
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	transaksiDetail, err := handler.service.GetListTransaksi_Detail(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", transaksiDetail)
	return echoCtx.JSON(res.Status, res)

}

func (handler *DetailTransaksiHandler) GetDetailTransaksi_Detail(echoCtx echo.Context) error {
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

	transaksiDetail, err := handler.service.GetDetailTransaksi_Detail(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", transaksiDetail)
	return echoCtx.JSON(res.Status, res)
}

func (handler *DetailTransaksiHandler) UpdateDetailTransaksi(echoCtx echo.Context) error {
	var form CreateDetailTransaksiBodyRequest
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

	_, err = handler.service.GetDetailTransaksi_Detail(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	transaksiDetailEntity := &entity.TransaksiDetail{
		id,
		form.Produk,
		int64(form.Kuantitas),
		form.Total,
	}

	if err := handler.service.UpdateDetailTransaksi(echoCtx.Request().Context(), transaksiDetailEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *DetailTransaksiHandler) DeleteDetailTransaksi(echoCtx echo.Context) error {
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

	err = handler.service.DeleteDetailTransaksi(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
