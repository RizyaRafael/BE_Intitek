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
	sort := c.Query("sort")
	status := c.Query("status")

	query := "select * from \"Products\""

	//update querry with the requested status
	if status != "" {
		query += " where status = " + status
	}

	//update querry to sort from stock
	if sort == "asc" {
		query += " order by quantity asc"
	}

	DB.Raw(query).Scan(&response)

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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "You've succesfully created the product",
	})

}

func UpdateProduct(c *fiber.Ctx) error {
	var updatedProduct model.Products

	//Check if body exist
	if err := c.BodyParser(&updatedProduct); err != nil {
		ErrorType = "INVALID_BODY"
		return handlers.ErrorHandler(ErrorType, c)
	}
	//Check if data format the same or not
	if updatedProduct.ID == 0 || updatedProduct.Location == "" || updatedProduct.SKU == "" || updatedProduct.Status == "" {
		ErrorType = "INVALID_BODY"
		return handlers.ErrorHandler(ErrorType, c)
	}

	//Check if data id exist or not
	if response := DB.Raw("Select * from \"Products\" where id = ?", updatedProduct.ID); response.Error != nil {
		ErrorType = "NOT_FOUND"
		return handlers.ErrorHandler(ErrorType, c)
	}

	// update the data
	response := DB.Exec("update \"Products\" set \"sku\" = ?, \"quantity\" = ?, \"location\" = ?, \"status\" = ? where id = ?", updatedProduct.SKU, updatedProduct.Quantity, updatedProduct.Location, updatedProduct.Status, updatedProduct.ID)
	if response.Error != nil || response.RowsAffected == 0 {
		return handlers.ErrorHandler("Internal server error", c)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product succesfully updated",
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("id")
	if err := DB.Exec("delete from \"Products\" where id = ?", productId); err.Error != nil {
		return handlers.ErrorHandler("internal server error", c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data succesfully deleted",
	})
}
