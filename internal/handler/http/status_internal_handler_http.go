package http

import (
	"echo-crud/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Status returns health check for the service.
func Status(echoCtx echo.Context) error {
	var res = entity.NewResponse(http.StatusOK, "It is work!", nil)
	return echoCtx.JSON(res.Status, res)
}
