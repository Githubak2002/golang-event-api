package main

import (
	// "log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/db"
	"github.com/githubak2002/golang-event-api/models"
)

func main() {

	db.InitDB()

	// Default returns an Engine instance with the Logger & Recovery middleware already attached
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.Run(":8080")

}

func getEvents(context *gin.Context){
	events,err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not parse request data.",
			"err":err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context){
	var event models.Event

	// It reads the JSON body from the incoming request.
	// It then maps (binds) the JSON fields to the fields of the event struct (models.Event).
	err := context.ShouldBindJSON(&event)
	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":"Could not parse request data", 
			"err":err.Error(),
		})
		// slog.Info("err: ",err.Error())
		return
	}

	event.Id = 1
	event.UesrId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":"Could not create event. Try again later!", 
			"err":err.Error(),
		})
	}


	context.JSON(http.StatusCreated, gin.H{"msg": "Event created", "event":event})
}

