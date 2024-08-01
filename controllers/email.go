package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/aliangelinaaa/witu-newsletter-backend/models"
)

// Input struct for creating an email
type CreateEmailInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"` // Added email validation
}

// Handler function to create a new email
func CreateEmail(c *gin.Context) {
	var input CreateEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new email record
	email := models.Email{Name: input.Name, Email: input.Email}
	if result := models.DB.Create(&email); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": email})
}

func FindEmails(c *gin.Context) {
	var emails []models.Email
	models.DB.Find(&emails)

	c.JSON(http.StatusOK, gin.H{"data": emails})
}

func FindEmail(c *gin.Context) {
	var email models.Email

	if err := models.DB.Where("id = ?", c.Param("id")).First(&email).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": email})
}

type UpdateEmailInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func UpdateEmail(c *gin.Context) {
	var email models.Email
	if err := models.DB.Where("id = ?", c.Param("id")).First(&email).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdateEmailInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEmail := models.Email{Name: input.Name, Email: input.Email}

	models.DB.Model(&email).Updates(&updatedEmail)
	c.JSON(http.StatusOK, gin.H{"data": email})
}

func DeleteEmail(c *gin.Context) {
	var email models.Email
	if err := models.DB.Where("id = ?", c.Param("id")).First(&email).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&email)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}