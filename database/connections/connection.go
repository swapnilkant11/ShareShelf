package connections

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	// Get the absolute path to the root directory
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting root directory: %v", err)
	}

	// Navigate up two level to the root directory
	rootDir = filepath.Dir(filepath.Dir(rootDir))

	// Construct the path to the .env file in the root directory
	envPath := filepath.Join(rootDir, ".env")

	// Load the .env file
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbOptions := os.Getenv("DB_OPTIONS")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbOptions)
	Instance, dbError := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return Instance, dbError
}
