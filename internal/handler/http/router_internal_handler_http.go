package http

import (
	"github.com/labstack/echo/v4"
)

// NewGinEngine creates an instance of echo.Engine.
// gin.Engine already implements net/http.Handler interface.
func NewGinEngine(supplierHandler *SupplierHandler, transaksiHandler *TransaksiHandler, pelangganHandler *PelangganHandler, produkHandler *ProdukHandler, pembayaranHandler *PembayaranHandler, detailTransaksiHandler *DetailTransaksiHandler, internalUsername, internalPassword string) *echo.Echo {
	engine := echo.New()

	engine.GET("/", Status)
	// engine.GET("/healthz", Health)
	// engine.GET("/version", Version)

	engine.POST("/create-supplier", supplierHandler.CreateSupplier)
	engine.POST("/create-pelanggan", pelangganHandler.CreatePelanggan)
	engine.POST("/create-transaksi", transaksiHandler.CreateTransaksi)
	engine.POST("/create-produk", produkHandler.CreateProduk)
	engine.POST("/create-pembayaran", pembayaranHandler.CreatePembayaran)
	engine.POST("/create-transaksi_detail", detailTransaksiHandler.CreateDetailTransaksi)

	engine.GET("/list-supplier", supplierHandler.GetListSupplier)
	engine.GET("/list-pelanggan", pelangganHandler.GetListPelanggan)
	engine.GET("/list-transaksi", transaksiHandler.GetListTransaksi)
	engine.GET("/list-produk", produkHandler.GetListProduk)
	engine.GET("/list-pembayaran", pembayaranHandler.GetListPembayaran)
	engine.GET("/list-transaksi_detail", detailTransaksiHandler.GetListTransaksi_Detail)
	

	engine.GET("/get-supplier/:id", supplierHandler.GetDetailSupplier)
	engine.GET("/get-pelanggan/:id", pelangganHandler.GetDetailPelanggan)
	engine.GET("/get-transaksi/:id", transaksiHandler.GetDetailTransaksi)
	engine.GET("/get-produk/:id", produkHandler.GetDetailProduk)
	engine.GET("/get-pembayaran/:id", pembayaranHandler.GetDetailPembayaran)
	engine.GET("/get-transaksi_detail/:id", detailTransaksiHandler.GetDetailTransaksi_Detail)

	engine.PUT("/update-supplier/:id", supplierHandler.UpdateSupplier)
	engine.PUT("/update-pelanggan/:id", pelangganHandler.UpdatePelanggan)
	engine.PUT("/update-transaksi/:id", transaksiHandler.UpdateTransaksi)
	engine.PUT("/update-produk/:id", produkHandler.UpdateProduk)
	engine.PUT("/update-pembayaran/:id", pembayaranHandler.UpdatePembayaran)
	engine.PUT("/update-transaksi_detail/:id", detailTransaksiHandler.UpdateDetailTransaksi)

	engine.DELETE("/delete-supplier/:id", supplierHandler.DeleteSupplier)
	engine.DELETE("/delete-pelanggan/:id", pelangganHandler.DeletePelanggan)
	engine.DELETE("/delete-transaksi/:id", transaksiHandler.DeleteTransaksi)
	engine.DELETE("/delete-produk/:id", produkHandler.DeleteProduk)
	engine.DELETE("/delete-pembayaran/:id", pembayaranHandler.DeletePembayaran)
	engine.DELETE("/delete-transaksi_detail/:id", detailTransaksiHandler.DeleteDetailTransaksi)

	// engine.POST("/uploud-file", fileHandler.CreateFile)

	return engine
}
