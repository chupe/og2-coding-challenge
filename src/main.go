package main

import (
	"context"
	"log"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/controllers"
	"github.com/chupe/og2-coding-challenge/database"

	_ "github.com/chupe/og2-coding-challenge/docs"
	_ "github.com/chupe/og2-coding-challenge/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title User Shortener API Demo
// @version 1.0
// @description This is a practice project for Go
// @termsOfService http://chupe.ba/terms/
// @contact.name Adnan
// @contact.email chupe@chupe.ba
// @license.name GPLv3
// @license.url https://www.gnu.org/licenses/gpl-3.0.html
// @host localhost:5000
// @accept json
// @produce json
// @schemes http
// @BasePath /
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

	app.Get("/swagger/*", swagger.HandlerDefault)

	controllers.RegisterHealthCheckHandler(app)
	controllers.RegisterCodeHandler(app, database.DbClient())

	api := app.Group("/api")
	controllers.RegisterUserHandler(api, database.DbClient())

	log.Fatal(app.Listen(":5000"))
}
