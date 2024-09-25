package controllers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/idtoken"
	"github.com/aliangelinaaa/witu-newsletter-backend/models"
)

func GoogleOAuth(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	idToken := c.Query("id_token")

	log.Printf("Received ID token: %s", idToken)

	if idToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID token is required"})
		return
	}

	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	if clientID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Client ID not found"})
		return
	}

	payload, err := verifyGoogleIDToken(idToken, clientID)
	if err != nil {
		log.Printf("Error verifying ID token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
		return
	}

	email := payload.Claims["email"].(string)
	name := payload.Claims["name"].(string)

	// Check if user exists in the database
	var user models.User
	if err := models.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// User not found, create a new user
		user = models.User{
			Email:     email,
			Name:      name,
			CreatedAt: time.Now(),
		}
		if err := models.DB.Create(&user).Error; err != nil {
			log.Printf("Error creating user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
			return
		}
	}

	// Return the user info
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	})
}

func verifyGoogleIDToken(idToken string, clientID string) (*idtoken.Payload, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	payload, err := idtoken.Validate(ctx, idToken, clientID)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
