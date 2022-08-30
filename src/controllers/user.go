package controllers

import (
	"errors"
	"net/http"

	"github.com/chupe/og2-coding-challenge/services"

	"github.com/chupe/og2-coding-challenge/data"
	"github.com/chupe/og2-coding-challenge/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	repo           *data.UserRepository
	factoryService *services.FactoryService
}

func NewUserHandler(repository *data.UserRepository, factoryService *services.FactoryService) *UserHandler {
	return &UserHandler{
		repo:           repository,
		factoryService: factoryService,
	}
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusNotFound).JSON(
			response.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Id must be provided",
				Error:   errors.New("id must be provided").Error(),
			})
	}

	user, err := h.repo.Find(id)
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
		Created:  user.Created,
	})
}

type createUser struct {
	// Full url
	Username string `json:"username" validate:"required,alphanum" example:"exampleUsername"`
} // @name CreateUserBody

func (h *UserHandler) Create(c *fiber.Ctx) error {
	v := validator.New()
	d := new(createUser)
	err := c.BodyParser(d)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid JSON",
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

	item, err := h.repo.Create(d.Username)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to save the User",
				Error:   err.Error(),
			})
	}

	return c.JSON(response.UserResponse{
		ID:       item.ID.Hex(),
		Username: item.Username,
		Created:  item.Created,
	})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusNotFound).JSON(
			response.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Id must be provided",
				Error:   errors.New("id must be provided").Error(),
			})
	}

	id, err := h.repo.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to save the User",
				Error:   err.Error(),
			})
	}

	return c.JSON(response.UserResponse{
		ID: id,
	})
}

func RegisterUserHandler(router fiber.Router, database *mongo.Client) {
	repo := data.NewUserRepository(database)
	factoryService := services.NewFactoryService()
	h := NewUserHandler(repo, factoryService)

	r := router.Group("/user")
	r.Get("/:id", h.Get)
	r.Post("/", h.Create)
	r.Delete("/:id", h.Delete)
}
