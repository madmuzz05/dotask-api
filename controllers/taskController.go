package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/madmuzz05/dotask-api.git/models"
	"gorm.io/datatypes"
)

func (h handler) GetTasks(c *gin.Context) {

	var tasks []model.Task
	result := h.DB.Find(&tasks)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &tasks,
	})
}
func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	result := h.DB.First(&task, id)
	if result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &task,
	})
}

func (h handler) StoreTask(c *gin.Context) {
	var task model.Task
	DateTask, _ := time.Parse("2006-01-02", c.PostForm("date_task"))
	StartTask, _ := time.Parse("2006-01-02 15:04", c.PostForm("start_task"))
	EndTask, _ := time.Parse("2006-01-02 15:04", c.PostForm("end_task"))

	task.TitleTask = c.PostForm("title_task")
	task.DateTask = datatypes.Date(DateTask)
	task.StartTask = StartTask
	task.EndTask = EndTask
	task.Category = c.PostForm("category")

	if result := h.DB.Create(&task); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusCreated, &task)
}

func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	DateTask, _ := time.Parse("2006-01-02", c.PostForm("date_task"))
	StartTask, _ := time.Parse("2006-01-02 15:04", c.PostForm("start_task"))
	EndTask, _ := time.Parse("2006-01-02 15:04", c.PostForm("end_task"))

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	task.TitleTask = c.PostForm("title_task")
	task.DateTask = datatypes.Date(DateTask)
	task.StartTask = StartTask
	task.EndTask = EndTask
	task.Category = c.PostForm("category")
	h.DB.Save(&task)

	c.JSON(http.StatusCreated, &task)
}

func (h handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	h.DB.Delete(&task)

	c.JSON(http.StatusCreated, &task)
}
