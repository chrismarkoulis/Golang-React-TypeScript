package database

import (
	"go-auth-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("goauth:goauth@/goauthapi"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB = conn

	conn.AutoMigrate(&models.User{})
}
