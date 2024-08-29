package routes

import (
	"mini-cms-api/app/middlewares"
	"mini-cms-api/controllers/categories"
	"mini-cms-api/controllers/contents"
	"mini-cms-api/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	UserController     users.UserController
	CategoryController categories.CategoryController
	ContentController  contents.ContentController
}

func (cl *ControllerList) RegisterRoute(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	// user routes
	userRoutes := e.Group("/api/v1/users")

	userRoutes.POST("/register", cl.UserController.Register)
	userRoutes.POST("/login", cl.UserController.Login)
	userRoutes.GET(
		"/info",
		cl.UserController.GetUserInfo,
		echojwt.WithConfig(cl.JWTMiddleware),
		middlewares.VerifyToken,
	)

	// category routes
	categoryRoutes := e.Group("/api/v1", echojwt.WithConfig(cl.JWTMiddleware), middlewares.VerifyToken)

	categoryRoutes.GET("/categories", cl.CategoryController.GetAll)
	categoryRoutes.GET("/categories/:id", cl.CategoryController.GetByID)
	categoryRoutes.POST("/categories", cl.CategoryController.Create)
	categoryRoutes.PUT("/categories/:id", cl.CategoryController.Update)
	categoryRoutes.DELETE("/categories/:id", cl.CategoryController.Delete)

	// content routes
	contentRoutes := e.Group("/api/v1", echojwt.WithConfig(cl.JWTMiddleware), middlewares.VerifyToken)

	contentRoutes.GET("/contents", cl.ContentController.GetAll)
	contentRoutes.GET("/contents/:id", cl.ContentController.GetByID)
	contentRoutes.POST("/contents", cl.ContentController.Create)
	contentRoutes.PUT("/contents/:id", cl.ContentController.Update)
	contentRoutes.DELETE("/contents/:id", cl.ContentController.Delete)

}
