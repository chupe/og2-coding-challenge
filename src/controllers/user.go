package controllers

import (
	"net/http"

	"github.com/chupe/og2-coding-challenge/data"
	"github.com/chupe/og2-coding-challenge/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	repository *data.UserRepository
}

// GetAll godoc
// @Summary Return all Users from the DB
// @ID get-Users
// @Tags User
// @Success 200 {array} []response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /api/User [get]
func (handler *UserHandler) GetAll(c *fiber.Ctx) error {
	result, err := handler.repository.FindAll()
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(
			response.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "No Users found in the DB",
				Error:   err.Error(),
			})
	}

	var res = make([]response.UserResponse, len(result))
	for i, v := range result {
		res[i] = response.UserResponse{
			ID:       v.ID.String(),
			Url:      v.Url,
			Code:     v.Code,
			HitCount: v.HitCount,
			Created:  v.Created,
		}
	}

	return c.JSON(res)
}

// Get godoc
// @Summary Return User by ID from the DB
// @ID get-User
// @Tags User
// @Param	id	path	int	true	"User ID"
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /api/User/{id} [get]
func (handler *UserHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	User, err := handler.repository.Find(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(
			response.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "User not found",
				Error:   err.Error(),
			})
	}

	return c.JSON(response.UserResponse{
		ID:       User.ID.String(),
		Url:      User.Url,
		Code:     User.Code,
		HitCount: User.HitCount,
		Created:  User.Created,
	})
}

type createUser struct {
	// Full url
	Username string `json:"username" validate:"required,alphanum" example:"exampleUsername"`
} // @name CreateUserBody

// Create godoc
// @Summary Send User URL to create a new shortened User
// @ID create-User
// @Tags User
// @Param Body body createUser true "JSON with a 'url' field that contains full URL"
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/User [post]
func (handler *UserHandler) Create(c *fiber.Ctx) error {
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

	item, err := handler.repository.Create(d.Url)
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
		Code:     item.Code,
		Url:      item.Url,
		HitCount: item.HitCount,
		Created:  item.Created,
	})
}

// Delete godoc
// @Summary Delete User object from the DB
// @ID delete-User
// @Tags User
// @Param	id	path	int	true	"User ID"
// @Success 204
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/User/{id} [delete]
func (handler *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid id param",
				Error:   err.Error(),
			})
	}
	_, err = handler.repository.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to delete",
				Error:   err.Error(),
			})
	}

	return c.SendStatus(http.StatusNoContent)
}

func NewUserHandler(repository *data.UserRepository) *UserHandler {
	return &UserHandler{
		repository: repository,
	}
}

func RegisterUserHandler(router fiber.Router, database *mongo.Client) {
	repository := data.NewUserRepository(database)
	UserHandler := NewUserHandler(repository)

	UserRouter := router.Group("/User")
	UserRouter.Get("/", UserHandler.GetAll)
	UserRouter.Get("/:id", UserHandler.Get)
	UserRouter.Post("/", UserHandler.Create)
	UserRouter.Delete("/:id", UserHandler.Delete)
}
