package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/veron-baranige/fire-bucket/internal/handlers"
)

func SetupFileRoutes(e *echo.Echo) {
	r := e.Group("/api/files")
	
	r.POST("", handlers.SaveFile)
	r.GET("/:id", handlers.GetFile)
	r.DELETE("/:id", handlers.DeleteFile)
}