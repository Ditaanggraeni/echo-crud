package main

import (
	"context"
	"echo-crud/internal/config"
	"echo-crud/internal/handler/http"
	"echo-crud/internal/repository"
	"echo-crud/internal/service"

	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Info().Msg("echo-crud-setting starting")
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// tool.ErrorClient = setupErrorReporting(context.Background(), cfg)

	var db *gorm.DB
	db = openDatabase(cfg)

	defer func() {
		if sqlDB, err := db.DB(); err != nil {
			log.Fatal().Err(err)
			panic(err)
		} else {
			_ = sqlDB.Close()
		}
	}()

	supplierHandler := buildSupplierHandler(db)
	transaksiHandler := buildTransaksiHandler(db)
	pelangganHandler := buildPelangganHandler(db)
	produkHandler := buildProdukHandler(db)

	engine := http.NewGinEngine(supplierHandler, transaksiHandler, pelangganHandler, produkHandler, cfg.InternalConfig.Username, cfg.InternalConfig.Password)

	server := &nethttp.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: engine,
	}

	// setGinMode(cfg.Env)
	runServer(server)
	waitForShutdown(server)
}

func runServer(srv *nethttp.Server) {
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
			log.Fatal().Err(err)
		}
	}()
}

func waitForShutdown(server *nethttp.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("shutting down echo-crud")

	// The context is used to inform the server it has 2 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("codelabs-service forced to shutdown")
	}

	log.Info().Msg("codelabs-service exiting")
}

func openDatabase(config *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checkError(err)
	return db
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func buildSupplierHandler(db *gorm.DB) *http.SupplierHandler {
	repo := repository.NewSupplierRepository(db)
	supplierService := service.NewSupplierService(repo)
	return http.NewSupplierHandler(supplierService)
}

func buildTransaksiHandler(db *gorm.DB) *http.TransaksiHandler {
	repo := repository.NewTransaksiRepository(db)
	transaksiService := service.NewTransaksiService(repo)
	return http.NewTransaksiHandler(transaksiService)
}

func buildPelangganHandler(db *gorm.DB) *http.PelangganHandler {
	repo := repository.NewPelangganRepository(db)
	pelangganService := service.NewPelangganService(repo)
	return http.NewPelangganHandler(pelangganService)
}

func buildProsukHandler(db *gorm.DB) *http.ProsukHandler {
	repo := repository.NewProsukRepository(db)
	prosukService := service.NewProsukService(repo)
	return http.NewProsukHandler(prosukService)
}
