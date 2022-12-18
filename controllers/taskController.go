package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/madmuzz05/dotask-api.git/models"
	"gorm.io/datatypes"
)

type Task struct {
	TitleTask string `json:"title_task" form:"TitleTask" binding:"required"`
	DateTask  string `json:"date_task" form:"DateTask" binding:"required"`
	StartTask string `json:"start_task" form:"StartTask" binding:"required"`
	EndTask   string `json:"end_task" form:"EndTask" binding:"required"`
	Category  string `json:"category" form:"Category" binding:"required"`
}

func (h handler) GetTasks(c *gin.Context) {

	var tasks []model.Task

	if result := h.DB.Find(&tasks); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   &tasks,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "Data tidak ada",
		})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
}

func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	if result := h.DB.First(&task, id); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   &task,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "Data tidak ada",
		})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

}

func (h handler) StoreTask(c *gin.Context) {
	req := Task{}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	DateTask, _ := time.Parse("2006-01-02", req.DateTask)

	var task model.Task
	task.TitleTask = req.TitleTask
	task.DateTask = datatypes.Date(DateTask)
	task.StartTask = req.StartTask
	task.EndTask = req.EndTask
	task.Category = req.Category

	if result := h.DB.Create(&task); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "Data berhasil ditambahkan",
			"result": &task,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "Gagal menambahkan data",
		})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
}

func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	req := Task{}

	var task model.Task
	if result := h.DB.First(&task, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "Data tidak ada",
		})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	DateTask, _ := time.Parse("2006-01-02", req.DateTask)

	task.TitleTask = req.TitleTask
	task.DateTask = datatypes.Date(DateTask)
	task.StartTask = req.StartTask
	task.EndTask = req.EndTask
	task.Category = req.Category
	if result := h.DB.Save(&task); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "Berhasil mengupdate data",
			"result": &task,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "Gagal mengupdate data",
		})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
}

func (h handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "Data tidak ada",
		})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Delete(&task); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "Data berhasil dihapus",
			"result": &task,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"data":   "Gagal menghapus data",
		})
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
}
