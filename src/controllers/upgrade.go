package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/data"
	"github.com/chupe/og2-coding-challenge/response"
	"github.com/chupe/og2-coding-challenge/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UpgradeHandler struct {
	repo           *data.UserRepository
	factoryService *services.FactoryService
}

func NewUpgradeHandler(repository *data.UserRepository, factoryService *services.FactoryService) *UpgradeHandler {
	return &UpgradeHandler{
		repo:           repository,
		factoryService: factoryService,
	}
}

type upgradeFactory struct {
	// Full url
	Username string `json:"username" validate:"required,alphanum" example:"exampleUsername"`
	Factory  string `json:"factory" validate:"required,alpha" example:"exampleFactory"`
} // @name UpgradeFactoryBody

func (h *UpgradeHandler) UpgradeFactory(c *fiber.Ctx) error {
	v := validator.New()
	d := new(upgradeFactory)
	err := c.BodyParser(d)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Username and factory are required",
				Error:   err.Error(),
			})
	}

	err = v.Struct(d)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation failed",
				Error:   err.Error(),
			})
	}

	user, err := h.repo.FindByUsername(d.Username)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Failed to find user by username",
				Error:   err.Error(),
			})
	}

	user, err = h.factoryService.UpgradeFactory(user, d.Factory)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to upgrade",
				Error:   err.Error(),
			})
	}

	err = h.repo.Update(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update user",
				Error:   err.Error(),
			})
	}

	return c.JSON(response.UserResponse{
		ID:       user.ID.Hex(),
		Username: user.Username,
	})
}

func RegisterUpgradeHandler(r fiber.Router, database *mongo.Client, fc *config.FactoryConfig) {
	repo := data.NewUserRepository(database)
	factoryService := services.NewFactoryService(fc)
	h := NewUpgradeHandler(repo, factoryService)

	r.Post("/upgrade", h.UpgradeFactory)
}
