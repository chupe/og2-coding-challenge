package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// HealthCheck godoc
// @Summary Check the status of the service
// @ID check-health
// @Tags Health
// @Success 200 {string} string "OK"
// @Router /health [get]
func (handler *HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func RegisterHealthCheckHandler(router fiber.Router) {
	healthCheckHandler := NewHealthCheckHandler()

	router.Get("/health", healthCheckHandler.HealthCheck)
}
