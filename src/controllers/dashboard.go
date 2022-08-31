package controllers

import (
	"errors"
	"net/http"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/data"
	"github.com/chupe/og2-coding-challenge/response"
	"github.com/chupe/og2-coding-challenge/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type DashboardHandler struct {
	repo *data.UserRepository
	fs   *services.FactoryService
}

func NewDashboardHandler(
	repository *data.UserRepository,
	factoryService *services.FactoryService,
) *DashboardHandler {
	return &DashboardHandler{
		repo: repository,
		fs:   factoryService,
	}
}

// GetDashboard godoc
// @Summary Return dashboard for username
// @ID get-dashboard
// @Tags dashboard
// @Param	username	query	string	true	"username"
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /dashboard [get]
func (h *DashboardHandler) GetDashboard(c *fiber.Ctx) error {
	username := c.Query("username")
	if username == "" {
		return c.Status(http.StatusNotFound).JSON(
			response.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Id must be provided",
				Error:   errors.New("id must be provided").Error(),
			})
	}

	user, err := h.repo.FindByUsername(username)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(
			response.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "User not found",
				Error:   err.Error(),
			})
	}

	ir, err := h.fs.GetRate(&user.IronFactory)
	cr, err := h.fs.GetRate(&user.CopperFactory)
	gr, err := h.fs.GetRate(&user.GoldFactory)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "failed to get factory rates",
				Error:   err.Error(),
			})
	}

	iron, err := h.fs.OreProduced(&user.IronFactory)
	copper, err := h.fs.OreProduced(&user.CopperFactory)
	gold, err := h.fs.OreProduced(&user.GoldFactory)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "failed to get amount of ores",
				Error:   err.Error(),
			})
	}

	return c.JSON(response.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Iron:     iron,
		Copper:   copper,
		Gold:     gold,
		Factories: []response.Factory{
			{
				Type:              string(user.IronFactory.Type),
				Level:             user.IronFactory.GetLevel(),
				RatePerMinute:     ir,
				UnderConstruction: user.IronFactory.UnderConstruction(),
				TimeToFinish:      user.IronFactory.TimeToUpgrade(),
			},
			{
				Type:              string(user.CopperFactory.Type),
				Level:             user.CopperFactory.GetLevel(),
				RatePerMinute:     cr,
				UnderConstruction: user.CopperFactory.UnderConstruction(),
				TimeToFinish:      user.CopperFactory.TimeToUpgrade(),
			},
			{
				Type:              string(user.GoldFactory.Type),
				Level:             user.GoldFactory.GetLevel(),
				RatePerMinute:     gr,
				UnderConstruction: user.GoldFactory.UnderConstruction(),
				TimeToFinish:      user.GoldFactory.TimeToUpgrade(),
			},
		},
		Created: user.Created,
	})
}

func RegisterDashboardHandler(r fiber.Router, database *mongo.Client, factoryConfig *config.Factories) {
	repo := data.NewUserRepository(database)
	fs := services.NewFactoryService(factoryConfig)
	h := NewDashboardHandler(repo, fs)

	r.Get("/dashboard", h.GetDashboard)
}
