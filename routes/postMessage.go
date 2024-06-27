package routes

import (
	"needlechat/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PostMessage(app *fiber.App) {
	app.Post("/message", func(c *fiber.Ctx) error {
		var messageData models.MessageForm
		if err := c.BodyParser(&messageData); err != nil {
			return err
		}
		uuid := uuid.NewString()
		res, _ := c.Locals("db").(*gorm.DB).Create(models.Message{MessageId: uuid, Content: messageData.Content, Author: messageData.Username, Nonce: messageData.Nonce}).Rows()
		defer res.Close()
		return c.JSON(res)
	})
}
