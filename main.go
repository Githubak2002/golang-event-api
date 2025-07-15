package main

import (
	// "log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/db"
	"github.com/githubak2002/golang-event-api/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil{
		panic("Error loading .env file")
	}

	db.InitDB()

	// Default returns an Engine instance with the Logger & Recovery middleware already attached
	server := gin.Default()

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":    "Hello from Golang!",
			"status": true,
		})
	})

	routes.RegisterRoutes(server)
	server.Run(":8080")

}
