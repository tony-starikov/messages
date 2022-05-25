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

type UpdateMessageInput struct {
	Message  string `json:"message"`
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


// GET /messages/:id
// Find a messages
func FindMessage(c *gin.Context) {  // Get model if exist
  var message models.Message

  if err := models.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": message})
}

// PATCH /message/:id
// Update a message
func UpdateMessage(c *gin.Context) {
  // Get model if exist
  var message models.Message

  if err := models.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  // Validate input
  var input UpdateMessageInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // Update attributes with `map`
  // models.DB.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
  models.DB.Model(&message).Update("Message", input.Message)

  c.JSON(http.StatusOK, gin.H{"data": message})
}

// DELETE /message/:id
// Delete a message
func DeleteMessage(c *gin.Context) {
  // Get model if exist
  var message models.Message
  if err := models.DB.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
    return
  }

  models.DB.Delete(&message)

  c.JSON(http.StatusOK, gin.H{"data": true})
}