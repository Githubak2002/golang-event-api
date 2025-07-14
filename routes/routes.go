package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/middlewares"
)

func RegisterRoutes(server *gin.Engine){

	// ===== Events routes =====
	server.GET("/events", getEvents)   
	server.GET("/event/:id", getEvent) 

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", createEvents)
	authenticated.DELETE("/event/:id", deleteEvent)
	authenticated.PUT("/event/:id", updateEvent)

	// server.POST("/events", middlewares.Authenticate ,createEvents)
	

	// ===== Users routes =====
	server.GET("/users", getAllUsers)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}