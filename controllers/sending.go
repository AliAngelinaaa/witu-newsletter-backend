package controllers

import (
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load environment variables from the .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// sendEmail sends an email using SMTP
func sendEmail(to string, subject string, body string) error {
	// Retrieve SMTP settings from environment variables
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	smtpUser := os.Getenv("EMAIL") // SMTP username
	smtpPass := os.Getenv("PSWD")  // SMTP password

	// Set up authentication information.
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpServer)

	// Create the email message.
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Send the email.
	addr := smtpServer + ":" + smtpPort
	err := smtp.SendMail(addr, auth, smtpUser, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}

// SendNewsletterEmail handles the request to send a newsletter email
func SendNewsletterEmail(c *gin.Context) {
	var input struct {
		To      string `json:"to" binding:"required,email"`
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body := input.Title + "\n\n" + input.Content
	err := sendEmail(input.To, input.Title, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
