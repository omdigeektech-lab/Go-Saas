package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"saas-task-manager/config"
	"saas-task-manager/routes"
)

func main() {
	_ = godotenv.Load()
	config.ConnectDB()

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
