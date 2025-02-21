package main

import (
	"github.com/gin-gonic/gin"
	"backend/db"
	"backend/handlers"
	"backend/websockets"
)

func main() {
	db.ConnectDatabase()
	r := gin.Default()
	r.POST("/login", handlers.LoginHandler)
	r.POST("/signup", handlers.SignupHandler)
	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
	r.GET("/ws", websockets.HandleConnections)
	r.Run(":8080")
}
