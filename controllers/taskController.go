package controller

import (
	"net/http"
	"strconv"
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
	Status    string `json:"status" form:"Status" binding:"required"`
	UserID    int    `json:"user_id" form:"UserID" binding:"required"`
}

func (h handler) GetTasks(c *gin.Context) {

	var tasks []model.Task

	allCount := h.DB.Find(&tasks).RowsAffected
	countSuccess := h.DB.Where("status = ?", "1").Find(&tasks).RowsAffected
	countError := h.DB.Where("status = ?", "0").Find(&tasks).RowsAffected

	data := map[string]interface{}{
		"all_count":     allCount,
		"count_success": countSuccess,
		"count_error":   countError,
		"tasks":         &tasks,
	}

	if result := h.DB.Preload("User").Preload("User.FriendLists", "status = ? ", "accepted").Preload("User.FriendLists.Friend").Find(&tasks); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    data,
			"success": "Success",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ada",
		})
	}
}

func (h handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	if result := h.DB.Preload("User").Preload("User.FriendLists", "status = ?", "accepted").Preload("User.FriendLists.Friend").First(&task, id); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    &task,
			"success": "success",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ada",
		})
	}

}

func (h handler) StoreTask(c *gin.Context) {
	req := Task{}

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	DateTask, _ := time.Parse("2006-01-02", req.DateTask)

	var task model.Task
	status, _ := strconv.Atoi(req.Status)
	task.TitleTask = req.TitleTask
	task.DateTask = datatypes.Date(DateTask)
	task.StartTask = req.StartTask
	task.EndTask = req.EndTask
	task.Category = req.Category
	task.Status = status
	task.UserID = req.UserID

	if result := h.DB.Create(&task); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Success",
			"data":    "Data berhasil ditambahkan",
			"result":  &task,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal menambahkan data",
		})
	}
}

func (h handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	req := Task{}

	var task model.Task
	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"sucess": "Data tidak ada",
		})
		return

	}

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	DateTask, _ := time.Parse("2006-01-02", req.DateTask)
	status, _ := strconv.Atoi(req.Status)
	task.TitleTask = req.TitleTask
	task.DateTask = datatypes.Date(DateTask)
	task.StartTask = req.StartTask
	task.EndTask = req.EndTask
	task.Category = req.Category
	task.Status = status
	task.UserID = req.UserID
	if result := h.DB.Save(&task); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    "Berhasil mengupdate data",
			"success": "Success",
			"result":  &task,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal mengupdate data",
		})
	}
}

func (h handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task

	if result := h.DB.First(&task, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ada",
		})
		return
	}

	if result := h.DB.Delete(&task); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    "Data berhasil dihapus",
			"success": "Success",
			"result":  &task,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal menghapus data",
		})
	}
}
