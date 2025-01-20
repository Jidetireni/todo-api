package routes

import (
	"net/http"
	"strconv"

	"github.com/Jidetireni/todo-api/models"
	"github.com/gin-gonic/gin"
)

func getTodolist(context *gin.Context) {
	todolist, err := models.GetAllTodolist()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todolist. Try again later.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, todolist)
}

func extractValidateId(context *gin.Context) (*models.Todo, bool) {
	todoId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse todo ID.", "error": err.Error()})
		return nil, false
	}

	todo, err := models.GetTodoById(todoId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch todo. Try again later.", "error": err.Error()})
		return nil, false
	}

	userId := context.GetInt64("userId")
	if userId != todo.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update todo."})
		return nil, false
	}

	return todo, true
}

func getTodo(context *gin.Context) {
	todo, ok := extractValidateId(context)
	if !ok {
		return
	}
	context.JSON(http.StatusOK, todo)
}

func createTodo(context *gin.Context) {

	var todo models.Todo
	err := context.ShouldBindJSON(&todo)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	userId := context.GetInt64("userId")
	todo.UserID = userId

	err = todo.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save todo", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Todo item created!", "todo": todo})

}

func updateTodo(context *gin.Context) {

	todo, ok := extractValidateId(context)
	if !ok {
		return
	}

	var updatedTodo models.Todo
	err := context.ShouldBindJSON(&updatedTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	updatedTodo.ID = todo.ID

	err = updatedTodo.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update todo", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo update!"})

}

func markCompleteTodo(context *gin.Context) {
	todo, ok := extractValidateId(context)
	if !ok {
		return
	}

	var completedTodo models.Todo
	err := context.ShouldBindJSON(&completedTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	completedTodo.ID = todo.ID

	err = completedTodo.Complete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to complete todo", "error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo completed!"})
}

func deleteTodo(context *gin.Context) {
	todo, ok := extractValidateId(context)
	if !ok {
		return
	}

	err := todo.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete todo", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted!"})

}
