package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	model "github.com/madmuzz05/dotask-api.git/models"
	"gorm.io/datatypes"
)

type User struct {
	Name     string `json:"name" form:"Name" binding:"required"`
	Username string `json:"username" form:"Username" binding:"required"`
	Password string `json:"password" form:"Password" binding:"required"`
	Email    string `json:"email" form:"Email" binding:"required"`
	Phone    string `json:"phone" form:"Phone" binding:"required"`
	Birthday string `json:"birthday" form:"Birthday" binding:"required"`
	Address  string `json:"address" form:"Address" binding:"required"`
}

func (h handler) GetUsers(c *gin.Context) {
	var users []model.User

	result := h.DB.Preload("FriendLists", "status = ?", "accepted").Preload("FriendLists.Friend").Find(&users)

	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Success",
			"data":    &users,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemukan",
		})
	}
}

func (h handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	result := h.DB.Preload("FriendLists", "status=?", "accepted").Preload("FriendLists.Friend").First(&user, id)
	if result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Success",
			"data":    &user,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemukan",
		})
	}
}

func (h handler) StoreUser(c *gin.Context) {
	req := User{}

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	birthday, _ := time.Parse("2006-01-02", req.Birthday)

	var user model.User

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	user.Address = req.Address
	user.Phone = req.Phone
	user.Birthday = datatypes.Date(birthday)

	if result := h.DB.Create(&user); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Success",
			"data":    "Data berhasil ditambahkan",
			"result":  &user,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal menambahkan data",
		})
	}
}

func (h handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	req := User{}

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var user model.User
	data := h.DB.First(&user, id)

	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemukan",
		})
		return
	}

	birthday, _ := time.Parse("2006-01-02", req.Birthday)

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	user.Address = req.Address
	user.Phone = req.Phone
	user.Birthday = datatypes.Date(birthday)

	if result := h.DB.Save(&user); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Success",
			"data":    "Berhasil mengupdate data",
			"result":  &user,
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal mengupdate data",
		})
	}
}

func (h handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	data := h.DB.First(&user, id)

	if data.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Data tidak ditemukan",
		})
		return
	}

	if result := h.DB.Delete(&user); result.Error == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"success": "Berhasil menghapus data",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"success": "Gagal menghapus data",
		})
	}

}
