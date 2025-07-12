package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/models/eventModel"
)

func getEvents(context *gin.Context) {
	events, err := eventModel.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "Could not parse request data.",
			"status": false,
			"err":    err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":    "All the events",
		"status": true,
		"events": events,
	})
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Could not parse event Id.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	event, err := eventModel.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Could not fetch event.",
			"err": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg":    "All the events",
		"status": true,
		"event": event,
	})

}

func createEvents(context *gin.Context) {
	var event eventModel.Event

	// It reads the JSON body from the incoming request.
	// It then maps (binds) the JSON fields to the fields of the event struct (eventModel.Event).
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Could not parse request data",
			"status": false,
			"err": err.Error(),
		})
		// slog.Info("err: ",err.Error())
		return
	}

	event.Id = 1
	event.UserId = 1

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Could not create event. Try again later!",
			"err": err.Error(),
		})
	}

	context.JSON(http.StatusCreated, gin.H{"msg": "Event created", "event": event})
}

func updateEvent(context *gin.Context){
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Could not parse event Id.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	_, err = eventModel.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "--- Could not fetch the event. --- ",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	var updatedEvent eventModel.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Could not parse request data.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	updatedEvent.Id = eventId
	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not Update the event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
			"msg": "Event Updated successfully.",
			"status": true,
		})

}