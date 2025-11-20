package route

import (
	"github.com/gofiber/fiber/v2"
    "go-proj/controllers"
)

func InitRoutes(app *fiber.App){

	api := app.Group("api/v1")
    api.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World! KKK")
    })

    api.Get("/:email", controllers.Register)
    api.Get("/send-mails", controllers.SendManyEmails)
}