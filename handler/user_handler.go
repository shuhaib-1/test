package handler

import (
	"strconv"
	"test/domain"
	"test/usecase"

	"github.com/gofiber/fiber"
)

type UserHandler struct {
	UseCase usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{UseCase: uc}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	userGroup := app.Group("/user")
	userGroup.Post("/", h.CreateUser)
	userGroup.Get("/get", h.GetUserById)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "failed to parse the input",
		})
	}

	if user.Email == "" || user.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name and email is required",
		})
	}

	if err := h.UseCase.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "could not create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUserById(c *fiber.Ctx) error {

	idStr := c.Query("id")
	if idStr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id format",
		})
	}

	user, err := h.UseCase.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(user)
}
