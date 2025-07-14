package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/models/event"
	// eventModel "github.com/githubak2002/golang-event-api/models/event"
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
			"status": false,
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

	userIdValue, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"msg":    "User ID not found in context",
			"status": false,
		})
		return
	}

	userId, ok := userIdValue.(int64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "User ID has invalid type",
			"status": false,
		})
		return
	}

	event.UserId = userId

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg": "Could not create event. Try again later!",
			"status": false,
			"err": err.Error(),
		})
	}

	context.JSON(http.StatusCreated, gin.H{
		"msg": "Event created", 
		"status": true,
		"event": event,
	})
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
			"msg": "Could not fetch the event.",
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
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not UPDATE the event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg": "Event UPDATED successfully.",
		"status": true,
	})

}

func deleteEvent(context *gin.Context) {
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
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not fetch the event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not DELETE the event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"msg": "Event DELETED successfully.",
		"status": true,
	})


}
