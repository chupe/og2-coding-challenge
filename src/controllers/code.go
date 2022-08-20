package controllers

import (
	"net/http"

	"github.com/chupe/og2-coding-challenge/data"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type CodeHandler struct {
	repository *data.UserRepository
}

// GetCode godoc
// @Summary Return User by ID from the DB
// @ID get-code
// @Tags Code
// @Param	code	path	int	true	"User Code"
// @Success 200 {object} response.UserResponse
// @Failure 404 {object} response.ErrorResponse
// @Router /{code} [get]
func (handler *CodeHandler) GetCode(c *fiber.Ctx) error {
	code := c.Params("code")
	User, err := handler.repository.FindByCode(code)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"status": http.StatusNotFound,
			"error":  err.Error(),
		})
	}

	return c.Redirect(User.Url)
}

func NewCodeHandler(repository *data.UserRepository) *CodeHandler {
	return &CodeHandler{
		repository: repository,
	}
}

func RegisterCodeHandler(router fiber.Router, database *mongo.Client) {
	repository := data.NewUserRepository(database)
	codeHandler := NewCodeHandler(repository)

	router.Get("/:code", codeHandler.GetCode)
}
