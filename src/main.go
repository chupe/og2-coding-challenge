package main

import (
	"context"
	"log"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/controllers"
	"github.com/chupe/og2-coding-challenge/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	database.DbClient()
	defer database.DbClient().Disconnect(context.TODO())

	controllers.RegisterHealthCheckHandler(app)

	api := app.Group("/")
	controllers.RegisterUserHandler(api, database.DbClient())
	controllers.RegisterDashboardHandler(api, database.DbClient())
	controllers.RegisterUpgradeHandler(api, database.DbClient())

	log.Fatal(app.Listen(":5000"))
}
