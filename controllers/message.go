package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "messages/models"
)

type CreateMessageInput struct {
  UserToID     uint   `json:"userToID" binding:"required"`
	UserFromID     uint   `json:"userFromID" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

// GET /messages
// Get all messages
func FindMessages(c *gin.Context) {
  var messages []models.Message
  models.DB.Find(&messages)

  c.JSON(http.StatusOK, gin.H{"data": messages})
}

// POST /message
// Create new message
func CreateMessage(c *gin.Context) {
  // Validate input
  var input CreateMessageInput

  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Create message
  message := models.Message{UserToID: input.UserToID, UserFromID: input.UserFromID, Message: input.Message}
  models.DB.Create(&message)

  c.JSON(http.StatusOK, gin.H{"data": message})
}