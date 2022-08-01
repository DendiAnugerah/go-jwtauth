package main

// 1:57:01
import (
	"log"
	"os"

	"github.com/DendiAnugerah/go-jwtauth/routers"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routers.AuthRoutes(router)
	routers.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Acces granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Acces granted for api-2"})
	})

	router.Run(":" + port)
}
