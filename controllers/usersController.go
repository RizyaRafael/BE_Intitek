package controllers

import (
	"BE/handlers"
	"BE/model"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var user model.Users

	//check if client send the form data
	if err := c.BodyParser(&user); err != nil {
		ErrorType = "INVALID_BODY"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//check if the form is filled or not
	if user.Password == "" || user.Email == "" {
		ErrorType = "INVALID_BODY"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//create the new user and send error if meet any validation constraint
	result := DB.Create(&user)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "23505") {
			ErrorType = "EMAIL_EXIST"
			return handlers.ErrorHandler(ErrorType, c)
		}
		return handlers.ErrorHandler("DATABASE_ERROR", c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "You've succesfully registered",
	})
}

func Login(c *fiber.Ctx) error {
	var user model.Users
	var foundUser model.Users

	//check if client send the form data
	if err := c.BodyParser(&user); err != nil {
		ErrorType = "INVALID_BODY"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//check if the form is filled or not
	if user.Email == "" || user.Password == "" {
		ErrorType = "EMAIL_AND_PASSWORD_REQ"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//check if email exist or not
	result := DB.Raw("select * from \"Users\" where email = ?", user.Email).Scan(&foundUser)
	if result.RowsAffected == 0 {
		ErrorType = "NOT_FOUND"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//Compare login password and password from DB
	checkPass := handlers.ComparePass(user.Password, foundUser.Password)
	if checkPass == nil {
		token, err := handlers.SignToken(foundUser.Email, c)
		if err != nil {
			return handlers.ErrorHandler("internal server error", c)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"access_token": "Bearer " + token,
		})
	} else {
		ErrorType = "INVALID_PASSWORD"
		return handlers.ErrorHandler(ErrorType, c)
	}
}
