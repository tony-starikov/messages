package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"messages/models"
	"messages/controllers"
)

func main() {
  router := gin.Default()

	models.ConnectDataBase()

	router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "Main page"})
  })

	router.GET("/users", controllers.FindUsers)
	router.POST("/users", controllers.CreateUser)

	router.GET("/messages", controllers.FindMessages)
	router.POST("/messages", controllers.CreateMessage)

  router.Run()
}