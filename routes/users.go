package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/models/user"
)

func signUp(ctx *gin.Context) {
	var user userModel.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg" : "Could not parse the data",
			"Status" : false,
			"error" : err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg" : "Could not SignUp the user",
			"Status" : false,
			"error" : err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated,gin.H{
		"msg" : "User SignUp successfully!",
		"Status" : true,
		"user" : user,
	})

}