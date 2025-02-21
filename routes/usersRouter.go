package routes

import (
	"BE/controllers"

	"github.com/gofiber/fiber/v2"
)

func UsersRouter(app fiber.Router) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
}
