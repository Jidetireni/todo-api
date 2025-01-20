package main

import (
	"github.com/Jidetireni/todo-api/db"
	"github.com/Jidetireni/todo-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080") // localhost:8080
}
