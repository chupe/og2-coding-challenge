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
	err := config.LoadToEnv()
	if err != nil {
		log.Fatal(err.Error())
	}

	fc := config.NewFactoryConfig()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	database.DbClient()
	defer database.DbClient().Disconnect(context.TODO())

	controllers.RegisterHealthCheckHandler(app)
	controllers.RegisterUserHandler(app, database.DbClient(), fc)
	controllers.RegisterDashboardHandler(app, database.DbClient(), fc)
	controllers.RegisterUpgradeHandler(app, database.DbClient(), fc)

	log.Fatal(app.Listen(":5000"))
}
