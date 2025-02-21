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
	case "EMAIL_EXIST":
		statusCode = fiber.StatusBadRequest
		errorMessage = "Email already registerd"
	case "EMAIL_AND_PASSWORD_REQ":
		statusCode = fiber.StatusBadRequest
		errorMessage = "Email and password required"
	case "NOT_FOUND":
		statusCode = fiber.StatusNotFound
		errorMessage = "Email is not registered"
		
	case "UNAUTHORIZED":
		statusCode = fiber.StatusForbidden
		errorMessage = "Please login"
	default:
		statusCode = fiber.StatusInternalServerError
		errorMessage = "an error occured"
	}

	return c.Status(statusCode).JSON(fiber.Map{
		"message": errorMessage,
	})
}
