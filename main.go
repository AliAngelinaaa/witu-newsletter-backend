// package main

// import (
// 	"github.com/aliangelinaaa/witu-newsletter-backend/controllers"
// 	"github.com/aliangelinaaa/witu-newsletter-backend/models"
// 	"github.com/gin-gonic/gin"
// )
// func main() {
	
// 	router := gin.Default()

// 	models.ConnectDatabase()

// 	router.POST("/posts", controllers.CreatePost)
// 	router.GET("/posts", controllers.FindPosts)
// 	router.GET("/posts/:id", controllers.FindPost)
// 	router.PATCH("/posts/:id", controllers.UpdatePost)
// 	router.DELETE("/posts/:id", controllers.DeletePost)

// 	router.Run("localhost:8080")
// }


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
        AllowOrigins:     []string{"http://localhost:5173"}, // Allow only your front-end URL
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        AllowCredentials: true,
    }))

    models.ConnectDatabase()

    router.POST("/posts", controllers.CreatePost)
    router.GET("/posts", controllers.FindPosts)
    router.GET("/posts/:id", controllers.FindPost)
    router.PATCH("/posts/:id", controllers.UpdatePost)
    router.DELETE("/posts/:id", controllers.DeletePost)

    router.Run("localhost:8080")
}
