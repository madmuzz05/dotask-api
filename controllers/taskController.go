package controller

import (
	"github.com/gin-gonic/gin"
	model "github.com/madmuzz05/dotask-api.git/models"
	"gorm.io/gorm"
)

func GetTasks(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var tasks []model.Task
	result := db.Find(&tasks)
	if result.Error != nil {
		c.Status(404)
		return
	}
	c.JSON(200, gin.H{
		"data": result,
	})
}

// func StoreTask(c *gin.Context) {

// 	var body struct {
// 		TitleTask string
// 		DateTask  datatypes.Date
// 		StartTask datatypes.Time
// 		EndTask   datatypes.Time
// 		Category  string
// 	}
// 	c.Bind(&body)
// 	post := model.Task{TitleTask: body.TitleTask, DateTask: body.DateTask, StartTask: body.StartTask, EndTask: body.EndTask, Category: body.Category}
// 	result := initialize.DB.Create(&post)

// 	if result.Error != nil {
// 		c.JSON(404, gin.H{
// 			"error": result.Error,
// 		})
// 		return
// 	}
// 	c.JSON(200, gin.H{
// 		"data":   post,
// 		"result": result,
// 	})

// }
