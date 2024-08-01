package auth

import (
    "context"
    "net/http"
    "github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
)

func GoogleLoginHandler(c *gin.Context) {
	
    url := oauth2Config.AuthCodeURL("state", oauth2.AccessTypeOffline)
    c.Redirect(http.StatusFound, url)
}

func GoogleCallbackHandler(c *gin.Context) {
    code := c.DefaultQuery("code", "")
    if code == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
        return
    }

    token, err := oauth2Config.Exchange(context.Background(), code)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
        return
    }

    idToken, ok := token.Extra("id_token").(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No ID token in response"})
        return
    }

    payload, err := VerifyToken(idToken)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": payload})
}
