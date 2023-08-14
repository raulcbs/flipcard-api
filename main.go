package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/raulcbs/flipcard-api/controller"
	"github.com/raulcbs/flipcard-api/database"
	"github.com/raulcbs/flipcard-api/middleware"
	"github.com/raulcbs/flipcard-api/migration"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	migration.Migration()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/deck", controller.AddDeck)
	protectedRoutes.GET("/deck/getAll", controller.GetAllDecks)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
