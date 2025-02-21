package handlers

import "github.com/gofiber/fiber/v2"

func ErrorHandler(err string, c *fiber.Ctx) error {
	var statusCode int
	var errorMessage string

	switch err {
	case "INVALID_BODY":
		statusCode = fiber.StatusBadRequest
		errorMessage = "All data is required"
	case "DATABASE_ERROR":
		statusCode = fiber.StatusBadRequest
		errorMessage = "Database error"
	default:
		statusCode = fiber.StatusInternalServerError
		errorMessage = "an error occured"
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message": errorMessage,
	})
}
