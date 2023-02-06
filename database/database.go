package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/victorbischoff/GOFIBER-TMPL/pkg/models"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&models.Song{})
	fmt.Println("Database Migrated")
}
