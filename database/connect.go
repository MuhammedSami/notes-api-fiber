package database

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/MuhammedSami/notes-api-fiber/config"
	"github.com/MuhammedSami/notes-api-fiber/internal/model"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT");
	port, err := strconv.ParseUint(p, 10, 32);

	if err != nil {
		log.Println("Could not get the credential from ENV")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to DB")
	}
	
	fmt.Println("Connection Opened to Database!")

	DB.AutoMigrate(&model.Note{})
	DB.AutoMigrate(&model.User{})
}