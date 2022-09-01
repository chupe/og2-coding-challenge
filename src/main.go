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

// @title OG2 Coding Challenge
// @version 0.1
// @description A coding challange with the task of finishing it in under 4 hours
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
	cfg := &config.Config{}
	config.LoadFromFile(cfg)

	db := database.Connect(&cfg.DB)
	defer db.Disconnect(context.TODO())

	env := &config.Env{
		DB:  db,
		Cfg: cfg,
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	controllers.RegisterHealthCheckHandler(app)
	controllers.RegisterUserHandler(app, env)
	controllers.RegisterDashboardHandler(app, env)
	controllers.RegisterUpgradeHandler(app, env)

	log.Fatal(app.Listen(":5000"))
}
