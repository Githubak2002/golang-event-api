package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/githubak2002/golang-event-api/models/user"
	"github.com/githubak2002/golang-event-api/utils"
)

func getAllUsers(ctx *gin.Context) {
	users, err := userModel.GetUsers()

	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "Could not get the Users!",
			"status": false,
			"error":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "All the Users",
		"status": true,
		"users": users,
	})
}

func login(ctx *gin.Context) {
	var user userModel.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Could not parse the data",
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := user.ValidateCreadentials(); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg":    "Could not Authenticate the user",
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "Could not Authenticate the user",
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "User Logged in Successfully!",
		"status": true,
		"token":  token,
	})

}

func signUp(ctx *gin.Context) {
	var user userModel.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":    "Could not parse the data",
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "Could not SignUp the user",
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	user.Password = "# Password Hashed #"
	// user.Password = ""

	ctx.JSON(http.StatusCreated, gin.H{
		"msg":    "User SignUp successfully!",
		"status": true,
		"user":   user,
	})

}
