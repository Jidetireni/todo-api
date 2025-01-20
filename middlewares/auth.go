package middlewares

import (
	"net/http"

	"github.com/Jidetireni/todo-api/utils"
	"github.com/gin-gonic/gin"
)

func Auth(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}
	context.Set("userId", userId)
	context.Next()
}
