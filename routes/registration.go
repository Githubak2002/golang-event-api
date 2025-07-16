package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	eventModel "github.com/githubak2002/golang-event-api/models/event"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Could not parse event Id.",
			"status": false,
			"err":    err.Error(),
		})
		return
	}

	event, err := eventModel.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not fetch event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	err = event.Register(userId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not Register user for event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"msg": "Registered user for event.",
		"status": true,
	})

}

func cancelRegistrationForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Could not parse event Id.",
			"status": false,
			"err":    err.Error(),
		})
		return
	}

	event, err := eventModel.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not fetch event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Could not Cancel Registration of user for the event.",
			"status": false,
			"err": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"msg": "You have successfully Cancelled your Registration for the event.",
		"status": true,
	})
}
