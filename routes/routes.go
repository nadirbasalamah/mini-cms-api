package routes

import (
	"mini-cms-api/controllers"
	"mini-cms-api/middlewares"
	"mini-cms-api/models"
	"mini-cms-api/utils"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// initialize logger middleware
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	// initialize authentication middleware
	jwtConfig := middlewares.JWTConfig{
		SecretKey: utils.GetConfig("JWT_SECRET_KEY"),
	}

	authMiddlewareConfig := jwtConfig.Init()

	jwtOptions := models.JWTOptions{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	// user routes
	userController := controllers.InitUserController(jwtOptions)
	userRoutes := e.Group("/api/v1/users")

	userRoutes.POST("/register", userController.Register)
	userRoutes.POST("/login", userController.Login)
	userRoutes.GET(
		"/info",
		userController.GetUserInfo,
		echojwt.WithConfig(authMiddlewareConfig),
		middlewares.VerifyToken,
	)

	// category routes
	categoryController := controllers.InitCategoryController()
	categoryRoutes := e.Group("/api/v1", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	categoryRoutes.GET("/categories", categoryController.GetAll)
	categoryRoutes.GET("/categories/:id", categoryController.GetByID)
	categoryRoutes.POST("/categories", categoryController.Create)
	categoryRoutes.PUT("/categories/:id", categoryController.Update)
	categoryRoutes.DELETE("/categories/:id", categoryController.Delete)

	// content routes
	contentController := controllers.InitContentController()
	contentRoutes := e.Group("/api/v1", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	contentRoutes.GET("/contents", contentController.GetAll)
	contentRoutes.GET("/contents/:id", contentController.GetByID)
	contentRoutes.POST("/contents", contentController.Create)
	contentRoutes.PUT("/contents/:id", contentController.Update)
	contentRoutes.DELETE("/contents/:id", contentController.Delete)
}
