package database

import (
	"fmt"

	"github.com/anthanh17/react-go-jwt-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456789",
		"jwt_auth")

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
