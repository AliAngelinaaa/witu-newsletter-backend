package main

import (
	"github.com/aliangelinaaa/witu-newsletter-backend/controllers"
	"github.com/aliangelinaaa/witu-newsletter-backend/models"
	"github.com/gin-gonic/gin"
)
func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/posts", controllers.FindPosts)
	router.GET("/posts/:id", controllers.FindPost)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run("localhost:8080")
}