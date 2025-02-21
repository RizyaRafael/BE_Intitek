package routes

import (
	"BE/controllers"

	"github.com/gofiber/fiber/v2"
)

func ProductsRouter(app fiber.Router) {
	app.Get("/", controllers.GetAllProducts)
	app.Post("/", controllers.CreateProduct)
}
