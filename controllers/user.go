package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "messages/models"
)

type CreateUserInput struct {
  Name  string `json:"name" binding:"required"`
}

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
  var users []models.User
  models.DB.Find(&users)

  c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new users
func CreateUser(c *gin.Context) {
  // Validate input
  var input CreateUserInput

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create user
  user := models.User{Name: input.Name}
  models.DB.Create(&user)

  c.JSON(http.StatusOK, gin.H{"data": user})
}