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
	db := initialize.ConnectToDB()
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/getTask", controller.GetTasks)
	// r.POST("/storeTask", controller.StoreTask)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
