package main

import (
	"needlechat/models"
	"needlechat/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{})

	db, _ := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

	db.AutoMigrate(&models.Message{}, &models.User{})

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Use(cors.New())

	routes.CreateUser(app)
	routes.Home(app)
	routes.PostMessage(app)
	routes.User(app)

	app.Listen(":3000")
}
