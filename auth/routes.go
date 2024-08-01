package auth

import (
    "github.com/gin-gonic/gin"
)

// SetupRoutes sets up the OAuth routes
func SetupRoutes() *gin.Engine {
    router := gin.Default()
    router.GET("/login", GoogleLoginHandler)
    router.GET("/callback", GoogleCallbackHandler)
    return router
}
