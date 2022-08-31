package controllers

import (
	"errors"
	"net/http"

	"github.com/chupe/og2-coding-challenge/config"
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
	factoryConfig  *config.FactoryConfig
}

func NewUserHandler(
	repository *data.UserRepository,
	factoryService *services.FactoryService,
	factoryConfig *config.FactoryConfig,
) *UserHandler {
	return &UserHandler{
		repo:           repository,
		factoryService: factoryService,
		factoryConfig:  factoryConfig,
	}
}

// GetUser godoc
// @Summary Get user by id
// @ID get-user
// @Tags user
// @Param	id	path	string	true	"user id"
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /user/{id} [get]
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
		Created:  user.Created,
	})
}

type createUser struct {
	// Full url
	Username string `json:"username" validate:"required,alphanum" example:"exampleUsername"`
} // @name CreateUserBody

// CreateUser godoc
// @Summary Create new user
// @ID create-user
// @Tags user
// @Param	createUser	body	createUser	true	"json containing username"
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /user [post]
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

// DeleteUser godoc
// @Summary Delete user by id
// @ID delete-user
// @Tags user
// @Param	id	path	string	true	"user id"
// @Success 204
// @Failure 404 {object} response.ErrorResponse
// @Router /user/{id} [delete]
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

	_, err := h.repo.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to save the User",
				Error:   err.Error(),
			})
	}

	return c.SendStatus(http.StatusNoContent)
}

func RegisterUserHandler(router fiber.Router, database *mongo.Client, factoryConfig *config.FactoryConfig) {
	repo := data.NewUserRepository(database)
	factoryService := services.NewFactoryService(factoryConfig)
	h := NewUserHandler(repo, factoryService, factoryConfig)

	r := router.Group("/user")
	r.Get("/:id", h.Get)
	r.Post("/", h.Create)
	r.Delete("/:id", h.Delete)
}
