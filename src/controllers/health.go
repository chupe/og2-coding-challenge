package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct{}

// HealthCheck godoc
// @Summary Check the status of the service
// @ID check-health
// @Tags Health
// @Success 200 {string} string "OK"
// @Router /health [get]
func (handler *HealthCheckHandler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func RegisterHealthCheckHandler(router fiber.Router) {
	healthCheckHandler := NewHealthCheckHandler()

	router.Get("/health", healthCheckHandler.HealthCheck)
}
