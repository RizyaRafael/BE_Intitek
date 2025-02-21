package middleware

import (
	"BE/handlers"
	"BE/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Authorization(c *fiber.Ctx) error {
	var user model.Users
	access_token := c.Get("Authorization")

	//check if access token exist or not
	if access_token == "" {
		return handlers.ErrorHandler("UNAUTHORIZED", c)
	}

	//check access token format
	checkBearer := strings.Split(access_token, " ")
	if checkBearer[0] != "Bearer" {
		return handlers.ErrorHandler("UNAUTHORIZED", c)
	}

	//verify token and verify if email exist in DB
	verifyToken, err := handlers.VerifyToken(checkBearer[1], c)
	if err != nil {
		return handlers.ErrorHandler("internal server error", c)
	}
	checkEmail := DB.Raw("select * from \"Users\" where \"email\" = ?", verifyToken).Scan(&user)
	if checkEmail.RowsAffected == 0 {
		return handlers.ErrorHandler("UNAUTHORIZED", c)
	}

	c.Locals("userId", user.ID)
	return c.Next()
}
