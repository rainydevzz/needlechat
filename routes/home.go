package routes

import (
	"needlechat/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Home(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		var messages []models.Message
		res, _ := c.Locals("db").(*gorm.DB).Find(&messages).Rows()
		defer res.Close()
		return c.JSON(messages)
	})
}
