package main

import (
	"ShareShelf/database/connection"
	"ShareShelf/database/models"
	"log"
)

func main() {
	// Connect to the database
	Instance, dbError := connection.ConnectDB()
	if dbError != nil {
		log.Fatal(dbError)
	}

	// Perform database migration
	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	log.Println("Database Migration Completed!")
}
