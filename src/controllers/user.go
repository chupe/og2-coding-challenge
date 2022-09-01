package controllers

import (
	"errors"
	"net/http"

	"github.com/chupe/og2-coding-challenge/config"
	"github.com/chupe/og2-coding-challenge/models"

	"github.com/chupe/og2-coding-challenge/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	users          *models.Users
	factoryService *models.Factories
	factoryConfig  *config.Factories
}

func NewUserHandler(
	userssitory *models.Users,
	factoryService *models.Factories,
	factoryConfig *config.Factories,
) *UserHandler {
	return &UserHandler{
		users:          userssitory,
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

	user, err := h.users.Find(id)
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

	item, err := h.users.Create(d.Username)
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

	_, err := h.users.Delete(id)
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

func RegisterUserHandler(r fiber.Router, env *config.Env) {
	users := models.NewUsers(env)
	fs := models.NewFactories(&env.Cfg.Factories)
	h := NewUserHandler(users, fs, &env.Cfg.Factories)

	rg := r.Group("/user")
	rg.Get("/:id", h.Get)
	rg.Post("/", h.Create)
	rg.Delete("/:id", h.Delete)
}
