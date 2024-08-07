package routes

import (
	"mini-cms-api/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// category routes
	categoryController := controllers.InitCategoryController()

	categoryRoutes := e.Group("/api/v1")

	categoryRoutes.GET("/categories", categoryController.GetAll)
	categoryRoutes.GET("/categories/:id", categoryController.GetByID)
	categoryRoutes.POST("/categories", categoryController.Create)
	categoryRoutes.PUT("/categories/:id", categoryController.Update)
	categoryRoutes.DELETE("/categories/:id", categoryController.Delete)

	// content routes
	contentController := controllers.InitContentController()

	contentRoutes := e.Group("/api/v1")

	contentRoutes.GET("/contents", contentController.GetAll)
	contentRoutes.GET("/contents/:id", contentController.GetByID)
	contentRoutes.POST("/contents", contentController.Create)
	contentRoutes.PUT("/contents/:id", contentController.Update)
	contentRoutes.DELETE("/contents/:id", contentController.Delete)
}
