package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lmittmann/tint"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/veron-baranige/fire-bucket/docs/swagger"
	"github.com/veron-baranige/fire-bucket/internal/config"
	db "github.com/veron-baranige/fire-bucket/internal/database"
	"github.com/veron-baranige/fire-bucket/internal/routes"
	"github.com/veron-baranige/fire-bucket/internal/storage"
)

// @title Fire-Bucket
// @version 1.0
// @description API for uploading and retrieving files through Firebase Storage
// @contact.name Veron Baranige
// @contact.email veronsajendra@gmail.com
// @host http://localhost:8000
// @BasePath /api
func main() {
	setupLogger()

	if err := config.LoadEnv("."); err != nil {
		slog.Error("Failed to load environment variables", "err", err)
		os.Exit(1)
	}
	slog.Info("Loaded configurations")

	if err := db.SetupClient(); err != nil {
		slog.Error("Failed establish database connection", "err", err)
		os.Exit(1)
	}
	slog.Info("Established database connection")

	if err := storage.Setup(); err != nil {
		slog.Error("Failed aquire storage bucekt connection", "err", err)
		os.Exit(1)
	}
	slog.Info("Established storage bucket connection")

	e := echo.New()

	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())
	setupRoutes(e)
	
	if err := e.Start(":" + config.Get(config.ServerPort)); err != nil {
		slog.Error("Failed to start the server", "err", err)
	}
}

func setupLogger() {
	logger := slog.New(tint.NewHandler(os.Stderr, nil))
	slog.SetDefault(logger)
}

func setupRoutes(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	routes.SetupFileRoutes(e)
}
