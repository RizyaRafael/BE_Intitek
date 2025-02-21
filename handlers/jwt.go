package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func SignToken(email string, c *fiber.Ctx) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	//Create a payload to encrypt
	claims := jwt.MapClaims{
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encrypt token
	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", ErrorHandler("Internal Server Error", c)
	}
	return signedToken, nil
}

func VerifyToken(access_token string, c *fiber.Ctx) (string, error) {
	//decrypt access token
	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(access_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", ErrorHandler("Internal server error", c)
	}
	claim := token.Claims.(jwt.MapClaims)
	return claim["email"].(string), nil
}

