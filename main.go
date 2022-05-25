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
	router.GET("/users/:id", controllers.FindUser)
	router.POST("/users", controllers.CreateUser)
	router.PATCH("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.GET("/messages", controllers.FindMessages)
	router.GET("/messages/:id", controllers.FindMessage)
	router.POST("/messages", controllers.CreateMessage)
	router.PATCH("/messages/:id", controllers.UpdateMessage)
	router.DELETE("/messages/:id", controllers.DeleteMessage)

  router.Run()
}