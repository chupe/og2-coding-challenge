package controllers

import (
	"errors"
	"net/http"

	"github.com/chupe/og2-coding-challenge/data"
	"github.com/chupe/og2-coding-challenge/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type DashboardHandler struct {
	repo *data.UserRepository
}

func NewDashboardHandler(repository *data.UserRepository) *DashboardHandler {
	return &DashboardHandler{
		repo: repository,
	}
}

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

	return c.JSON(response.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Iron:     user.GetIronOre(),
		Copper:   user.GetCopperOre(),
		Gold:     user.GetGoldOre(),
		Factories: []response.Factories{
			{
				Type:              string(user.IronFactory.Type),
				Level:             user.IronFactory.GetLevel(),
				RatePerMinute:     user.IronFactory.GetRate(),
				UnderConstruction: user.IronFactory.UnderConstruction(),
				TimeToFinish:      user.IronFactory.TimeToUpgrade(),
			},
			{
				Type:              string(user.CopperFactory.Type),
				Level:             user.CopperFactory.GetLevel(),
				RatePerMinute:     user.CopperFactory.GetRate(),
				UnderConstruction: user.CopperFactory.UnderConstruction(),
				TimeToFinish:      user.CopperFactory.TimeToUpgrade(),
			},
			{
				Type:              string(user.GoldFactory.Type),
				Level:             user.GoldFactory.GetLevel(),
				RatePerMinute:     user.GoldFactory.GetRate(),
				UnderConstruction: user.GoldFactory.UnderConstruction(),
				TimeToFinish:      user.GoldFactory.TimeToUpgrade(),
			},
		},
		Created: user.Created,
	})
}

func RegisterDashboardHandler(router fiber.Router, database *mongo.Client) {
	repo := data.NewUserRepository(database)
	h := NewDashboardHandler(repo)

	r := router.Group("/")
	r.Get("/dashboard", h.GetDashboard)
}
