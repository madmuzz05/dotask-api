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
	routes := api.Group("/task")
	routes.GET("/", h.GetTasks)
	routes.GET("/find/:id", h.GetTask)
	routes.PUT("/update/:id", h.UpdateTask)
	routes.DELETE("/delete/:id", h.DeleteTask)
	routes.POST("/store", h.StoreTask)

}
