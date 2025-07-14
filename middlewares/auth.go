package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/utils"
)


// NOTE: AbortWithStatusJSON - Stops Gin handler chain
// return - Stops your function
// context.Next() - To fo the the next middleware -or- handler 

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":    "Not Authorized! - Token missing",
			"status": false,
		})
		return
	}

	userId, err := utils.ValidToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"msg":    "Not Authorized! - Invalid Token",
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	context.Set("userId",userId)			// this sets a key-value pair in the request's context — like storing metadata
	context.Next()							// To fo the the next middleware -or- handler

}