package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(supplierHandler *SupplierHandler, transaksiHandler *TransaksiHandler, pelangganHandler *PelangganHandler, internalUsername, internalPassword string) *echo.Echo {
	engine := echo.New()

	engine.GET("/", Status)
	// engine.GET("/healthz", Health)
	// engine.GET("/version", Version)

	engine.POST("/create-supplier", supplierHandler.CreateSupplier)
	engine.POST("/create-pelanggan", pelangganHandler.CreatePelanggan)
	engine.POST("/create-transaksi", transaksiHandler.CreateTransaksi)

	engine.GET("/list-supplier", supplierHandler.GetListSupplier)
	engine.GET("/list-pelanggan", pelangganHandler.GetListPelanggan)
	engine.GET("/list-transaksi", transaksiHandler.GetListTransaksi)

	engine.GET("/get-supplier/:id", supplierHandler.GetDetailSupplier)
	engine.GET("/get-pelanggan/:id", pelangganHandler.GetDetailPelanggan)
	engine.GET("/get-transaksi/:id", transaksiHandler.GetDetailTransaksi)

	engine.PUT("/update-supplier/:id", supplierHandler.UpdateSupplier)
	engine.PUT("/update-pelanggan/:id", pelangganHandler.UpdatePelanggan)
	engine.PUT("/update-transaksi/:id", transaksiHandler.UpdateTransaksi)

	engine.DELETE("/delete-supplier/:id", supplierHandler.DeleteSupplier)
	engine.DELETE("/delete-pelanggan/:id", pelangganHandler.DeletePelanggan)
	engine.DELETE("/delete-transaksi/:id", transaksiHandler.DeleteTransaksi)

	// engine.POST("/uploud-file", fileHandler.CreateFile)

	return engine
}
