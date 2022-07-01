package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
)

type User struct {
	Id       string
	Name     string
	LastName string
	Age      int
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Name:     "John",
		LastName: "Doe",
		Age:      30,
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.Id = uuid.New().String()

	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userGroup := app.Group("/users")
	app.Use(requestid.New())
	userGroup.Get("", handleUser)
	userGroup.Post("", handleCreateUser)

	app.Listen(":3000")
}
