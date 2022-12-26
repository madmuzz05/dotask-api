package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h :=
		&handler{
			DB: db,
		}
	api := router.Group("/api")
	task := api.Group("/task")
	task.GET("/", h.GetTasks)
	task.GET("/find/:id", h.GetTask)
	task.PUT("/update/:id", h.UpdateTask)
	task.DELETE("/delete/:id", h.DeleteTask)
	task.POST("/store", h.StoreTask)

	user := api.Group("/user")
	user.GET("/", h.GetUsers)
	user.GET("/find/:id", h.GetUser)
	user.POST("/store", h.StoreUser)
	user.PUT("/update/:id", h.UpdateUser)
	user.DELETE("/delete/:id", h.DeleteUser)

	event := api.Group("/event")
	event.GET("/:id", h.GetEvents)
	event.GET("/find/:id", h.GetEvent)
	event.POST("/store", h.StoreEvent)
	event.PUT("/update/:id", h.UpdateEvent)
	event.DELETE("/delete/:id", h.DeleteEvent)

}
