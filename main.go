package main

import (
	"fmt"
	"net/http"
	"task3/Auth"
	"task3/models"
	"task3/store"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	mongoStore := &store.MongoStore{}
	mongoStore.OpenConnectionWithMongoDB()

	// Server (using gin framework)
	router := gin.Default()

	// Configure CORS middleware
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	// Setup CORS middleware
	// This applies the default CORS policies
	router.POST("/user/signup", func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := mongoStore.StoreUserData(newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store user data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
	})
	router.POST("/user/login", func(c *gin.Context) {
		var loginCredentials models.User
		err := c.ShouldBindJSON(&loginCredentials)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isAuthenticated := mongoStore.UserLogin(loginCredentials.Username, loginCredentials.Password)
		if isAuthenticated {
			c.JSON(http.StatusOK, gin.H{"message": "successful login"})
			signedToken, signedRefreshToken, err := Auth.GenerateAllTokens(loginCredentials.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": signedToken, "refreshToken": signedRefreshToken})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		}
	})
	// Listen and serve on the specified port
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	if err := router.Run(port); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
