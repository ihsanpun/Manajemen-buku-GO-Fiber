package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihsanpun/go-fiber-part2/database"
	"github.com/ihsanpun/go-fiber-part2/database/migration"
	"github.com/ihsanpun/go-fiber-part2/route"
)

func main() {
	//INITIAL DATABASE
	database.DatabaseInit()
	migration.RunMigration()

	app := fiber.New()

	//INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":3000")

}
