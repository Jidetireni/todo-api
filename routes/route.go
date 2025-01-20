package routes

import (
	"github.com/Jidetireni/todo-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/todolist", getTodolist)

	authed := router.Group("/")
	authed.Use(middlewares.Auth)

	authed.GET("/todolist/:id", getTodo)
	authed.POST("/todolist", createTodo)
	authed.PUT("/todolist/:id", updateTodo)
	authed.PUT("/todolist/:id/complete", markCompleteTodo)
	authed.DELETE("/todolist/:id", deleteTodo)

	router.POST("/signup", signUp)
	router.POST("/login", login)
}
