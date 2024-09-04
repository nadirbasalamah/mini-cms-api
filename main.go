package main

import (
	"fmt"
	"os"

	_driverFactory "mini-cms-api/drivers"

	_contentUseCase "mini-cms-api/businesses/contents"
	_contentController "mini-cms-api/controllers/contents"

	_categoryUseCase "mini-cms-api/businesses/categories"
	_categoryController "mini-cms-api/controllers/categories"

	_userUseCase "mini-cms-api/businesses/users"
	_userController "mini-cms-api/controllers/users"

	_dbDriver "mini-cms-api/drivers/mysql"

	_middlewares "mini-cms-api/app/middlewares"
	_routes "mini-cms-api/app/routes"

	echo "github.com/labstack/echo/v4"
)

func main() {
	configDB := _dbDriver.DBConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.MigrateDB(db)

	configJWT := _middlewares.JWTConfig{
		SecretKey:       os.Getenv("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	categoryRepo := _driverFactory.NewCategoryRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoryUseCase(categoryRepo)
	categoryCtrl := _categoryController.NewCategoryController(categoryUseCase)

	contentRepo := _driverFactory.NewContentRepository(db)
	contentUseCase := _contentUseCase.NewContentUseCase(contentRepo)
	contentCtrl := _contentController.NewContentController(contentUseCase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUseCase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware:   configLogger.Init(),
		JWTMiddleware:      configJWT.Init(),
		CategoryController: *categoryCtrl,
		ContentController:  *contentCtrl,
		UserController:     *userCtrl,
	}

	routesInit.RegisterRoute(e)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	e.Logger.Fatal(e.Start(port))
}
