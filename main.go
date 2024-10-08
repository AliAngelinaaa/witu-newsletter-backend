package main

import (
    "github.com/aliangelinaaa/witu-newsletter-backend/controllers"
    "github.com/aliangelinaaa/witu-newsletter-backend/models"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Setup CORS
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    }))

    models.ConnectDatabase()

    router.POST("/posts", controllers.CreatePost)
    router.GET("/posts", controllers.FindPosts)
    router.GET("/posts/:id", controllers.FindPost)
    router.PATCH("/posts/:id", controllers.UpdatePost)
    router.DELETE("/posts/:id", controllers.DeletePost)

	router.POST("/emails", controllers.CreateEmail)
	router.GET("/emails", controllers.FindEmails)
	router.GET("/emails/:id", controllers.FindEmail)
	router.PATCH("/emails/:id", controllers.UpdateEmail)
	router.DELETE("/emails/:id", controllers.DeleteEmail)

	router.POST("/sendingnewsletter", controllers.SendNewsletterEmail)

    router.GET("/auth/google", controllers.GoogleOAuth)

    router.Run("localhost:8080")
}
