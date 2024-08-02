package auth

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// User represents Google user information structure
type User struct {
	Email          string `json:"email"`
	Family_name    string `json:"family_name"`
	Given_name     string `json:"given_name"`
	Id             string `json:"id"`
	Locale         string `json:"locale"`
	Name           string `json:"name"`
	Picture        string `json:"picture"`
	Verified_email bool   `json:"verified_email"`
}

var oauthConf *oauth2.Config

// SetupRoutes sets up the OAuth routes for authentication using Gin
func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// Load environment variables
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading .env file:", err)
	}

	oauthConf = &oauth2.Config{
		ClientID:     viper.GetString("GOOGLE_CLIENT_ID"),
		ClientSecret: viper.GetString("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  viper.GetString("REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	router.GET("/auth/google", func(c *gin.Context) {
		url := oauthConf.AuthCodeURL("state") // Get auth URL
		c.Redirect(http.StatusFound, url)     // Redirect to Google auth URL
	})

	router.GET("/auth/redirect", func(c *gin.Context) {
		code := c.Query("code") // Get code from query params
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
			return
		}

		token, err := oauthConf.Exchange(context.Background(), code) // Exchange code for token
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token: " + err.Error()})
			return
		}

		client := oauthConf.Client(context.Background(), token)                      // Set client for getting user info
		response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo") // Get user info
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info: " + err.Error()})
			return
		}
		defer response.Body.Close()

		var user User
		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body: " + err.Error()})
			return
		}
		if err := json.Unmarshal(bytes, &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling JSON body: " + err.Error()})
			return
		}

		// Do something with user info (e.g., store in database)

		c.JSON(http.StatusOK, user) // Return user info
	})

	return router
}
