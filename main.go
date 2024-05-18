package main

import (
	"ShareShelf/controllers"
	"ShareShelf/database/connection"
	"ShareShelf/middlewares"
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
	router := initRouter()
	router.Run(PORT)
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
