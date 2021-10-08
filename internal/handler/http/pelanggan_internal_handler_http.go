package http

import (
	"echo-crud/entity"
	"echo-crud/internal/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreatePelangganBodyRequest defines all body attributes needed to add pelanggan.
type CreatePelangganBodyRequest struct {
	NamaPelanggan string `json:"nama_pelanggan" binding:"required"`
	Telepon       string `json:"telepon" binding:"required"`
	Alamat        string `json:"alamat" binding:"required"`
}

// PelangganRowResponse defines all attributes needed to fulfill for pelanggan row entity.
type PelangganRowResponse struct {
	ID            uuid.UUID `json:"id_pelanggan"`
	NamaPelanggan string    `json:"nama_pelanggan"`
	Telepon       string    `json:"telepon"`
	Alamat        string    `json:"alamat"`
}

// PelangganResponse defines all attributes needed to fulfill for pic pelanggan entity.
type PelangganDetailResponse struct {
	ID            uuid.UUID `json:"id_pelanggan"`
	NamaPelanggan string    `json:"nama_pelanggan"`
	Telepon       string    `json:"telepon"`
	Alamat        string    `json:"alamat"`
}

func buildPelangganRowResponse(pelanggan *entity.Pelanggan) PelangganRowResponse {
	form := PelangganRowResponse{
		ID:            pelanggan.ID,
		NamaPelanggan: pelanggan.NamaPelanggan,
		Alamat:        pelanggan.Alamat,
		Telepon:       pelanggan.Telepon,
	}

	return form
}

func buildPelangganDetailResponse(pelanggan *entity.Pelanggan) PelangganDetailResponse {
	form := PelangganDetailResponse{
		ID:            pelanggan.ID,
		NamaPelanggan: pelanggan.NamaPelanggan,
		Alamat:        pelanggan.Alamat,
		Telepon:       pelanggan.Telepon,
	}

	return form
}

// QueryParamsPelanggan defines all attributes for input query params
type QueryParamsPelanggan struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaPelanggan define attributes needed for Meta
type MetaPelanggan struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaPelanggan creates an instance of Meta response.
func NewMetaPelanggan(limit, offset int, total int64) *MetaPelanggan {
	return &MetaPelanggan{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// PelangganHandler handles HTTP request related to user flow.
type PelangganHandler struct {
	service service.PelangganUseCase
}

// NewPelangganHandler creates an instance of PelangganHandler.
func NewPelangganHandler(service service.PelangganUseCase) *PelangganHandler {
	return &PelangganHandler{
		service: service,
	}
}

// Create handles pelanggan creation.
// It will reject the request if the request doesn't have required data,
func (handler *PelangganHandler) CreatePelanggan(echoCtx echo.Context) error {
	var form CreatePelangganBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	pelangganEntity := entity.NewPelanggan(
		uuid.Nil,
		form.NamaPelanggan,
		form.Telepon,
		form.Alamat,
	)

	if err := handler.service.Create(echoCtx.Request().Context(), pelangganEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", pelangganEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *PelangganHandler) GetListPelanggan(echoCtx echo.Context) error {
	var form QueryParamsPelanggan
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	pelanggan, err := handler.service.GetListPelanggan(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", pelanggan)
	return echoCtx.JSON(res.Status, res)

}

func (handler *PelangganHandler) GetDetailPelanggan(echoCtx echo.Context) error {
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

	pelanggan, err := handler.service.GetDetailPelanggan(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", pelanggan)
	return echoCtx.JSON(res.Status, res)
}

func (handler *PelangganHandler) UpdatePelanggan(echoCtx echo.Context) error {
	var form CreatePelangganBodyRequest
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

	_, err = handler.service.GetDetailPelanggan(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	pelangganEntity := &entity.Pelanggan{
		id,
		form.NamaPelanggan,
		form.Alamat,
		form.Telepon,
	}

	if err := handler.service.UpdatePelanggan(echoCtx.Request().Context(), pelangganEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *PelangganHandler) DeletePelanggan(echoCtx echo.Context) error {
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

	err = handler.service.DeletePelanggan(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
