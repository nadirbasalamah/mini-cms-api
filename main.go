package main

import (
	"mini-cms-api/utils"

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
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.MigrateDB(db)

	configJWT := _middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
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

	e.Logger.Fatal(e.Start(":1323"))
}
