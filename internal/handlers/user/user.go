package userHandlers

import (
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/config"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/database"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/model"
	"github.com/golang-jwt/jwt"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	db := database.DB
	payload := new(model.Register)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error",
			"message": "Invalid Input.. Review your request body",
			"data":    err})
	}
	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error",
			"message": "Invalid Input.. Review your request body",
			"data":    errors})
	}
	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	}
	hashbytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"staus": "error", "message": "Internal Server error"})
	}
	payload.Password = string(hashbytes)
	newuser := model.User{
		ID:        uuid.New(),
		UserName:  payload.Username,
		Email:     payload.Email,
		Password:  string(hashbytes),
		FirstName: payload.FirstName,
	}

	err = db.Create(&newuser).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error",
			"message": "Cant register the user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User successfully registered",
		"data":    newuser.ID})
}

func Dashboard(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"Status":  "error",
			"message": "No user entries found",
			"data":    nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "users found", "data": users})
}

func Signin(c *fiber.Ctx) error {
	db := database.DB
	payload := new(model.Login)

	err := c.BodyParser(payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error",
			"message": "Invalid Input.. Review your request body",
			"data":    err})
	}
	errors := model.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error",
			"message": "Invalid Input.. Review your request body",
			"data":    errors})
	}
	var user model.User
	result := db.Where("email=?", payload.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid email or password"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Invalid password or email"})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	time_now := time.Now().UTC()
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["exp"] = time_now.Add(time.Minute * 60).Unix()
	claims["iat"] = time_now.Unix()
	claims["nbf"] = time_now.Unix()

	tokenstring, err := token.SignedString([]byte(config.Config("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Failure in Genrating the jwt token"})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenstring,
		Path:     "/",
		MaxAge:   60 * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  tokenstring})

}
