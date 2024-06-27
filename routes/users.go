package routes

import (
	"needlechat/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func User(app *fiber.App) {
	app.Get("/users", func(c *fiber.Ctx) error {
		var user []models.User
		c.Locals("db").(*gorm.DB).Find(&user)
		return c.JSON(user)
	})
}
