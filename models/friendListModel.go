package model

import "time"

type FriendList struct {
	FriendListID uint      `gorm:"primaryKey" json:"id_friend_list"`
	UserID       int       `json:"user_id"`
	PersonID     int       `json:"person_id"`
	Status       string    `gorm:"size:20" json:"status"`
	Friend       User      `gorm:"foreignKey:UserID;references:PersonID"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
