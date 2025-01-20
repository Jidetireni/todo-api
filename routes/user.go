package routes

import (
	"net/http"

	"github.com/Jidetireni/todo-api/models"
	"github.com/Jidetireni/todo-api/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save todo", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user successfully created!"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = user.Validate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "could not auth user!", "error": err.Error()})
		return
	}

	token, err := utils.GenereateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "user login successful!", "token": token})
}
