package main

import (
	"ShareShelf/database/connection"
	"ShareShelf/middlewares"
	"ShareShelf/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	_, dbError := connection.ConnectDB()
	if dbError != nil {
		log.Fatal(dbError)
	}

	PORT := os.Getenv("PORT")

	// Initialize Router
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.Auth())
	routes.SecureRoutes(router)
	router.Run(PORT)
}
