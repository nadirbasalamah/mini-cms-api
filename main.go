package main

import (
	"mini-cms-api/database"
	"mini-cms-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
