package main

import (
	"github.com/gin-gonic/gin"

	controller "github.com/madmuzz05/dotask-api.git/controllers"
	"github.com/madmuzz05/dotask-api.git/initialize"
)

func init() {
	initialize.LoadInitializeEnv()
}

func main() {
	r := gin.Default()
	db := initialize.ConnectToDB()

	controller.RegisterRoutes(r, db)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
	r.Run()
}
