package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	if DB == nil {
		DB = getConnection()
	}

	return DB
}

func getConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv(`DATABASE_NAME`)

	dbHost := os.Getenv(`DATABASE_HOST`)
	dbUser := os.Getenv(`DATABASE_USER`)
	dbPass := os.Getenv(`DATABASE_PASSWORD`)
	dbPort := os.Getenv(`DATABASE_PORT`)
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("? Connected Successfully to the Database")

	return db
}
