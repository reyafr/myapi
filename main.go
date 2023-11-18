package main

import (
	"github.com/gin-gonic/gin"
	"D:\Rey\Belajar\myapi\controllers" // Update import path to reference local package
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/api/users", controllers.GetUsers)
	r.GET("/api/users/:id", controllers.GetUserByID)
	r.POST("/api/users", controllers.CreateUser)
	r.PUT("/api/users/:id", controllers.UpdateUserByID)
	r.DELETE("/api/users/:id", controllers.DeleteUserByID)

	r.Run(":8080")
}
