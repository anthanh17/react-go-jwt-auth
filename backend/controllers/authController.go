package controllers

import (
	"github.com/anthanh17/react-go-jwt-auth/database"
	"github.com/anthanh17/react-go-jwt-auth/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	passWd, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: passWd,
	}
	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	data := make(map[string]string)

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// data in database
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// compare data database and data web request
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	// passwd successfully
	return c.JSON(user)
}
