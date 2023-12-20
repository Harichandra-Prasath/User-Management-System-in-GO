package UserMiddleware

import (
	"fmt"
	"strings"

	"github.com/Harichandra-Prasath/User-Management-System-in-GO/config"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/database"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Authorize(c *fiber.Ctx) error {
	var tokenstring string
	authorization := c.Get("Authorization")
	if strings.HasPrefix(authorization, "Bearer ") { // if the client doesnt support cookie
		tokenstring = strings.TrimPrefix(authorization, "Bearer ")

	} else if c.Cookies("token") != "" { // if the client can send it in form of cookies
		tokenstring = c.Cookies("token")
	}

	if tokenstring == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized, please login"})
	}

	//checking the sign in method
	token, err := jwt.Parse(tokenstring, func(jwttoken *jwt.Token) (interface{}, error) {
		if _, ok := jwttoken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwttoken.Header["alg"])
		}
		return []byte(config.Config("SECRET_KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"Message": fmt.Sprintf("Invalid token %v", err)})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid token claim"})
	}

	var user model.User
	db := database.DB
	result := db.Find(&user, "id=?", fmt.Sprint(claims["sub"]))
	if result.Error != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status": "error",
			"data":   result.Error})
	}
	return c.Next()
}
