package main

import (
	"log"
	"net/http"

	"app.com/db"
	"app.com/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Automigrate the schema

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	server.GET("/users", func(c *gin.Context) {
		var users []models.User
		result :=db.Find(&users)
		if result.Error != nil {
			log.Fatalf("failed to retrieve users: %v", result.Error)
		}
		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully created user",
			"user":    users,
		})

	})
	server.POST("/users", func(c *gin.Context) {
		var user models.User

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the data"})
			return
		}
		user.Save()
		c.JSON(http.StatusCreated, gin.H{
			"message": "Successfully created user",
			"user":    user,
		})
	})

	server.Run(":8080")
}
