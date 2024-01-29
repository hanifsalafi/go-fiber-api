package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/database"
	"go-fiber-api/route"
)

func main() {

	// INITIAL DATABASE
	database.DatabaseInit()
	//migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.MainRouteInit(app)
	route.UserRouteInit(app)

	err := app.Listen(":8800")
	if err != nil {
		return
	}
}
