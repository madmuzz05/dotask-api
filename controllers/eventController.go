package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/madmuzz05/dotask-api.git/models"
	"gorm.io/datatypes"
)

type Event struct {
	TitleEvent string `json:"title_event" form:"TitleEvent" binding:"required"`
	DateEvent  string `json:"date_event" form:"DateEvent" binding:"required"`
	StartEvent string `json:"start_event" form:"StartEvent" binding:"required"`
	EndEvent   string `json:"end_event" form:"EndEvent" binding:"required"`
	Location   string `json:"location" form:"Location" binding:"required"`
	Status     string `json:"status" form:"Status" binding:"required"`
	UserID     int    `json:"user_id" form:"UserID" binding:"required"`
}

// It's a function to get all data from database.
func (h handler) GetEvents(c *gin.Context) {
	var events []model.Event

	allCount := h.DB.Find(&events).RowsAffected
	countSuccess := h.DB.Where("status=?", "1").Find(&events).RowsAffected
	countError := h.DB.Where("status=?", "0").Find(&events).RowsAffected

	data := map[string]interface{}{
		"all_count":     allCount,
		"count_success": countSuccess,
		"count_error":   countError,
		"events":        &events,
	}

	if result := h.DB.Preload("User").Preload("User.FriendLists", "status = ?", "accepted").Preload("User.FriendLists.Friend").Find(&events); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    data,
			"success": "Success",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemnukan",
		})
	}
}

func (h handler) GetEvent(c *gin.Context) {
	id := c.Param("id")
	var event model.Event

	if result := h.DB.Preload("User").Preload("User.FriendLists", "status = ? ", "accepted").Preload("User.FriendLists.Friend").Where("user_id = ?", id).Find(&event); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    &event,
			"success": "Success",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemukan",
		})
	}
}

func (h handler) StoreEvent(c *gin.Context) {
	var req = Event{}

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	DateEvent, _ := time.Parse("2006-01-02", req.DateEvent)
	var event model.Event
	status, _ := strconv.Atoi(req.Status)
	event.TitleEvent = req.TitleEvent
	event.DateEvent = datatypes.Date(DateEvent)
	event.StartEvent = req.StartEvent
	event.EndEvent = req.EndEvent
	event.Location = req.Location
	event.Status = status
	event.UserID = req.UserID

	if result := h.DB.Create(&event); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Success",
			"data":    "Data berhasil ditambahkan",
			"result":  &event,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal menambahkan data",
		})
	}
}

func (h handler) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	var req = Event{}

	var event model.Event
	if data := h.DB.First(&event, id); data.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemukan",
		})
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	DateEvent, _ := time.Parse("2006-01-02", req.DateEvent)
	status, _ := strconv.Atoi(req.Status)
	event.TitleEvent = req.TitleEvent
	event.DateEvent = datatypes.Date(DateEvent)
	event.StartEvent = req.StartEvent
	event.EndEvent = req.EndEvent
	event.Location = req.Location
	event.Status = status
	event.UserID = req.UserID

	if result := h.DB.Save(&event); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Success",
			"data":    "Berhasil mengupdate data",
			"result":  &event,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal menambahkan data",
		})
	}
}

func (h handler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")

	var event model.Event
	if data := h.DB.First(&event, id); data.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemukan",
		})
		return
	}

	if result := h.DB.Delete(&event); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    "Data berhasil dihapus",
			"success": "Success",
			"result":  &event,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal menghapus data",
		})
	}
}
