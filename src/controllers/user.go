package controllers

import (
	"net/http"

	"github.com/chupe/og2-coding-challenge/data"
	"github.com/chupe/og2-coding-challenge/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	repository *data.UserRepository
}

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
		Username: User.Username,
		Iron:     User.GetIronOre(),
		Copper:   User.GetCopperOre(),
		Gold:     User.GetGoldOre(),
		Created:  User.Created,
	})
}

func (handler *UserHandler) GetDashboard(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := handler.repository.Find(id)

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
				Rate:              user.IronFactory.GetRate(),
				UnderConstruction: user.IronFactory.UnderConstruction(),
				TimeToFinish:      user.IronFactory.TimeToUpgrade(),
			},
			{
				Type:              string(user.CopperFactory.Type),
				Level:             user.CopperFactory.GetLevel(),
				Rate:              user.CopperFactory.GetRate(),
				UnderConstruction: user.CopperFactory.UnderConstruction(),
				TimeToFinish:      user.CopperFactory.TimeToUpgrade(),
			},
			{
				Type:              string(user.GoldFactory.Type),
				Level:             user.GoldFactory.GetLevel(),
				Rate:              user.GoldFactory.GetRate(),
				UnderConstruction: user.GoldFactory.UnderConstruction(),
				TimeToFinish:      user.GoldFactory.TimeToUpgrade(),
			},
		},
		Created: user.Created,
	})
}

type createUser struct {
	// Full url
	Username string `json:"username" validate:"required,alphanum" example:"exampleUsername"`
} // @name CreateUserBody

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

	item, err := handler.repository.Create(d.Username)
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

func (handler *UserHandler) UpgradeFactory(c *fiber.Ctx) error {
	fac := c.Params("factory")
	username := c.Params("username")
	user, err := handler.repository.UpgradeFactory(username, fac)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			response.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to upgrade",
				Error:   err.Error(),
			})
	}

	return c.JSON(response.UserResponse{
		ID:       user.ID.Hex(),
		Username: user.Username,
	})
}

func NewUserHandler(repository *data.UserRepository) *UserHandler {
	return &UserHandler{
		repository: repository,
	}
}

func RegisterUserHandler(router fiber.Router, database *mongo.Client) {
	repository := data.NewUserRepository(database)
	userHandler := NewUserHandler(repository)

	UserRouter := router.Group("/user")
	UserRouter.Get("/:id", userHandler.Get)
	UserRouter.Post("/", userHandler.Create)
	UserRouter.Post("/dashboard/:username", userHandler.GetDashboard)
}
