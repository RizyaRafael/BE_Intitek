package routes

import (
	"BE/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	user := app.Group("user")
	UsersRouter(user)

	app.Use(middleware.Authorization)

	product := app.Group("/product")
	ProductsRouter(product)
}
