package http

import (
	"echo-crud/entity"
	"echo-crud/internal/service"
	"net/http"
	nethttp "net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// CreateProdukBodyRequest defines all body attributes needed to add Produk.
type CreateProdukBodyRequest struct {
	KodeProduk string `json:"kode_produk" binding:"required"`
	NamaProduk       string `json:"nama_produk" binding:"required"`
	Harga        int `json:"harga" binding:"required"`
	Stok        int64 `json:"stok" binding:"required"`
}

// ProdukRowResponse defines all attributes needed to fulfill for Produk row entity.
type ProdukRowResponse struct {
	Id            uuid.UUID `json:"id"`
	KodeProduk string    `json:"kode_produk"`
	NamaProduk       string    `json:"nama_Produk"`
	Harga        int    `json:"harga"`
	Stok        int64    `json:"stok"`
}

// ProdukResponse defines all attributes needed to fulfill for pic Produk entity.
type ProdukDetailResponse struct {
	Id            uuid.UUID `json:"id_Produk"`
	KodeProduk string    `json:"kode_produk"`
	NamaProduk       string    `json:"nama_Produk"`
	Harga        int    `json:"harga"`
	Stok        int64    `json:"stok"`
}

func buildProdukRowResponse(produk *entity.Produk) ProdukRowResponse {
	form := ProdukRowResponse{
		Id:            produk.Id,
		KodeProduk: produk.KodeProduk,
		NamaProduk:        produk.NamaProduk,
		Harga:       produk.Harga,
		Stok:       produk.Stok,
	}

	return form
}

func buildProdukDetailResponse(produk *entity.Produk) ProdukDetailResponse {
	form := ProdukDetailResponse{
		Id:            produk.Id,
		KodeProduk: produk.KodeProduk,
		NamaProduk:        produk.NamaProduk,
		Harga:       produk.Harga,
		Stok:       produk.Stok,
	}

	return form
}

// QueryParamsProduk defines all attributes for input query params
type QueryParamsProduk struct {
	Limit  string `form:"limit"`
	Offset string `form:"offset"`
	SortBy string `form:"sort_by"`
	Order  string `form:"order"`
	Status string `form:"status"`
}

// MetaProduk define attributes needed for Meta
type MetaProduk struct {
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

// NewMetaProduk creates an instance of Meta response.
func NewMetaProduk(limit, offset int, total int64) *MetaProduk {
	return &MetaProduk{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}

// ProdukHandler handles HTTP request related to user flow.
type ProdukHandler struct {
	service service.ProdukUseCase
}

// NewProdukHandler creates an instance of ProdukHandler.
func NewProdukHandler(service service.ProdukUseCase) *ProdukHandler {
	return &ProdukHandler{
		service: service,
	}
}

// Create handles Produk creation.
// It will reject the request if the request doesn't have required data,
func (handler *ProdukHandler) CreateProduk(echoCtx echo.Context) error {
	var form CreateProdukBodyRequest
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	produkEntity := entity.NewProduk(
		uuid.Nil,
		form.KodeProduk,
		form.NamaProduk,
		int(form.Harga),
		int64(form.Stok),
	)

	if err := handler.service.Create(echoCtx.Request().Context(), produkEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusCreated, "Request processed successfully.", produkEntity)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProdukHandler) GetListProduk(echoCtx echo.Context) error {
	var form QueryParamsProduk
	if err := echoCtx.Bind(&form); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInvalidInput)
		return echoCtx.JSON(nethttp.StatusBadRequest, errorResponse)
	}

	produk, err := handler.service.GetListProduk(echoCtx.Request().Context(), form.Limit, form.Offset)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}
	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", produk)
	return echoCtx.JSON(res.Status, res)

}

func (handler *ProdukHandler) GetDetailProduk(echoCtx echo.Context) error {
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

	produk, err := handler.service.GetDetailProduk(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", produk)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProdukHandler) UpdateProduk(echoCtx echo.Context) error {
	var form CreateProdukBodyRequest
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

	_, err = handler.service.GetDetailProduk(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	produkEntity := &entity.Produk{
		id,
		form.KodeProduk,
		form.NamaProduk,
		form.Harga,
		form.Stok,
		
	}

	if err := handler.service.UpdateProduk(echoCtx.Request().Context(), produkEntity); err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(nethttp.StatusInternalServerError, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}

func (handler *ProdukHandler) DeleteProduk(echoCtx echo.Context) error {
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

	err = handler.service.DeleteProduk(echoCtx.Request().Context(), id)
	if err != nil {
		errorResponse := buildErrorResponse(err, entity.ErrInternalServerError)
		return echoCtx.JSON(http.StatusBadRequest, errorResponse)
	}

	var res = entity.NewResponse(nethttp.StatusOK, "Request processed successfully.", nil)
	return echoCtx.JSON(res.Status, res)
}
