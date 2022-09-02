package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/models"
	"github.com/chupe/og2-coding-challenge/response"
	"github.com/gofiber/fiber/v2"
)

type UpgradeHandler struct {
	users          *models.Users
	factoryService *models.Factories
}

func NewUpgradeHandler(users *models.Users, factoryService *models.Factories) *UpgradeHandler {
	return &UpgradeHandler{
		users:          users,
		factoryService: factoryService,
	}
}

type upgradeFactory struct {
	// Full url
	Username string `json:"username" validate:"required,alphanum" example:"exampleUsername"`
	Factory  string `json:"factory" validate:"required,alpha" example:"exampleFactory"`
} // @name UpgradeFactoryBody

// UpgradeFactory godoc
// @Summary Upgrade factory type for a user
// @ID upgrade-factory
// @Tags factory
// @Param	UpgradeFactoryBody	body	upgradeFactory	true	"username and factory type"
// @Success 204
// @Failure 404 {object} response.ErrorResponse
// @Router /upgrade [post]
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

	user, err := h.users.FindByUsername(d.Username)
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
		return c.Status(http.StatusInternalServerError).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to upgrade",
				Error:   err.Error(),
			})
	}

	err = h.users.Update(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to update user",
				Error:   err.Error(),
			})
	}

	return c.SendStatus(http.StatusNoContent)
}

func RegisterUpgradeHandler(r fiber.Router, env *config.Env) {
	users := models.NewUsers(env)
	fs := models.NewFactories(&env.Cfg.Factories)
	h := NewUpgradeHandler(users, fs)

	r.Post("/upgrade", h.UpgradeFactory)
}
