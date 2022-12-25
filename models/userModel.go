package model

import (
	"time"

	"gorm.io/datatypes"
)

type User struct {
	UserID      uint           `gorm:"primaryKey" json:"id_user"`
	Name        string         `gorm:"size:50" json:"name"`
	Username    string         `gorm:"size:20" json:"username"`
	Password    string         `json:"password"`
	Email       string         `gorm:"size:30" json:"email"`
	Phone       string         `gorm:"size:20" json:"phone"`
	Birthday    datatypes.Date `json:"birthday"`
	Address     string         `gorm:"size:75" json:"address"`
	FriendLists []FriendList   `gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
