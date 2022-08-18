package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@(localhost:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		"admin",
		"admin123",
		"testmysql")

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database")
	}

}
