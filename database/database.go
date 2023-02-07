package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/victorbischoff/GOFIBER-TMPL/pkg/config"
	"github.com/victorbischoff/GOFIBER-TMPL/pkg/models"
)

var DBConn *gorm.DB


func InitDatabase(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	fmt.Println("Connection Opened to Database")
	DBConn.AutoMigrate(&models.Song{})
	fmt.Println("Database Migrated")
}
