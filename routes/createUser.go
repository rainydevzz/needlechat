package routes

import (
	"needlechat/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(app *fiber.App) {
	app.Post("/user", func(c *fiber.Ctx) error {
		var userData models.UserForm
		if err := c.BodyParser(&userData); err != nil {
			return err
		}
		uuid := uuid.NewString()
		res, _ := c.Locals("db").(*gorm.DB).Create(models.User{UserId: uuid, PublicKey: userData.PublicKey, Name: userData.Name}).Rows()
		defer res.Close()
		return c.JSON(res)
	})
}
