package controllers

import (
	"BE/handlers"
	"BE/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ErrorType string

func GetAllProducts(c *fiber.Ctx) error {
	var response []model.Products
	DB.Raw("select * from \"Products\"").Scan(&response)
	return c.Status(fiber.StatusOK).JSON(response)

}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct model.Products

	//check if client send the form data
	if err := c.BodyParser(&newProduct); err != nil {
		ErrorType = "INVALID_BODY"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//check if the form is filled or not
	if newProduct.Location == "" || newProduct.SKU == "" || newProduct.Status == "" {
		ErrorType = "INVALID_BODY"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//insert data into DB and check if its succesfull or not
	result := DB.Create(&newProduct)
	if result.Error != nil {
		return handlers.ErrorHandler("DATABASE_ERROR", c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "You've succesfully registered",
	})

}
