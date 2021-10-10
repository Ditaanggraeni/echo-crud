package http

import (
	"echo-crud/entity"
	"echo-crud/internal/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreatePembayaranBodyRequest defines all body attributes needed to add Pembayaran.
type CreatePembayaranBodyRequest struct {
	TglBayar    string `json:"tgl_bayar" binding:"required"`	
	Total      int64  `json:"total" binding:"required"`
}

// PembayaranRowResponse defines all attributes needed to fulfill for Pembayaran row entity.
type PembayaranRowResponse struct {
	Id        uuid.UUID `json:"id`
	TglBayar    string    `json:"tanggal"`
	Total      int64     `json:"total"`
	//PelangganID uuid.UUID `json:"pelanggan_id"`
}

// PembayaranResponse defines all attributes needed to fulfill for pic Pembayaran entity.
type PembayaranDetailResponse struct {
	Id         uuid.UUID `json:"id_Pembayaran"`
	TglBayar    string    `json:"tanggal"`
	Total      int64     `json:"total"`
	//PelangganID uuid.UUID `json:"pelanggan_id"`
}

func buildPembayaranRowResponse(pembayaran *entity.Pembayaran) PembayaranRowResponse {
	form := PembayaranRowResponse{
		Id:         pembayaran.Id,
		TglBayar:    pembayaran.TglBayar,
		Total:      pembayaran.Total,
		//PelangganID: Pembayaran.PelangganID,
	}

	return form
}

func buildPembayaranDetailResponse(pembayaran *entity.Pembayaran) PembayaranDetailResponse {
	form := PembayaranDetailResponse{
		Id:         pembayaran.Id,
		TglBayar:    pembayaran.TglBayar,
		Total:      pembayaran.Total,
		//PelangganID: Pembayaran.PelangganID,
	}

	return form
}

// QueryParamsPembayaran defines all attributes for input query params
type QueryParamsPembayaran struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaPembayaran define attributes needed for Meta
// type MetaPembayaran struct {
// 	Limit  int   `json:"limit"`
// 	Offset int   `json:"offset"`
// 	Total  int64 `json:"total"`
// }

// // NewMetaPembayaran creates an instance of Meta response.
// func NewMetaPembayaran(limit, offset int, total int64) *MetaPembayaran {
// 	return &MetaPembayaran{
// 		Limit:  limit,
// 		Offset: offset,
// 		Total:  total,
// 	}
// }

// PembayaranHandler handles HTTP request related to user flow.
type PembayaranHandler struct {
	service service.PembayaranUseCase
}

// NewPembayaranHandler creates an instance of PembayaranHandler.
func NewPembayaranHandler(service service.PembayaranUseCase) *PembayaranHandler {
	return &PembayaranHandler{
		service: service,
	}
}

// Create handles Pembayaran creation.
// It will reject the request if the request doesn't have required data,
func (handler *PembayaranHandler) CreatePembayaran(echoCtx echo.Context) error {
	var form CreatePembayaranBodyRequest

	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	pembayaranEntity := entity.NewPembayaran(
		uuid.Nil,
		form.TglBayar,
		int64(form.Total),
	)

	if err := handler.service.Create(echoCtx.Request().Context(), pembayaranEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", pembayaranEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *PembayaranHandler) GetListPembayaran(echoCtx echo.Context) error {
	var form QueryParamsPembayaran
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	pembayaran, err := handler.service.GetListPembayaran(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", pembayaran)
	return echoCtx.JSON(res.Status, res)

}

func (handler *PembayaranHandler) GetDetailPembayaran(echoCtx echo.Context) error {
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

	pembayaran, err := handler.service.GetDetailPembayaran(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", pembayaran)
	return echoCtx.JSON(res.Status, res)
}

func (handler *PembayaranHandler) UpdatePembayaran(echoCtx echo.Context) error {
	var form CreatePembayaranBodyRequest
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

	_, err = handler.service.GetDetailPembayaran(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	pembayaranEntity := &entity.Pembayaran{
		id,
		form.TglBayar,
		form.Total,
	}

	if err := handler.service.UpdatePembayaran(echoCtx.Request().Context(), pembayaranEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *PembayaranHandler) DeletePembayaran(echoCtx echo.Context) error {
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

	err = handler.service.DeletePembayaran(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
