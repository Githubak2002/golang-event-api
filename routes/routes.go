package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)   // GET, POST, PUT, PATCH, DELETE
	server.GET("/event/:id", getEvent) // /events/1  /events/2

	server.POST("/events", createEvents)

	server.PUT("/event/:id", updateEvent)
}