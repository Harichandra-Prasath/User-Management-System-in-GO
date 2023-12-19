package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Harichandra-Prasath/User-Management-System-in-GO/config"
	"github.com/Harichandra-Prasath/User-Management-System-in-GO/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("Cant parse the port number")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password =%s dbname=%s sslmode=disable", config.Config("DB_HOST"),
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect Database")

	}
	DB.AutoMigrate(&model.User{})
	fmt.Println("Connection Established")
}
